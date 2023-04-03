package httplogs

import (
	"io/ioutil"
	"log"
	"net"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/spudtrooper/goutil/io"
	"github.com/spudtrooper/goutil/slice"
)

// Two-level cache. First try in `cache`, then read from disk. In case we want to force
// resolving IPs for which we found not hosts, but only do it once.
type ipCache struct {
	dir                 string
	sep                 string
	cache               map[string][]string
	cacheMu             sync.Mutex
	forceResolveEmpties bool
}

func newDiskCache(forceResolveEmpties bool) (*ipCache, error) {
	dir, err := io.MkdirAll(".ipCache")
	if err != nil {
		return nil, err
	}
	res := &ipCache{
		dir:                 dir,
		sep:                 ",",
		cache:               map[string][]string{},
		forceResolveEmpties: forceResolveEmpties,
	}
	return res, nil
}

func (d *ipCache) Get(ip string) ([]string, bool, error) {
	d.cacheMu.Lock()
	defer d.cacheMu.Unlock()
	if hosts, ok := d.cache[ip]; ok {
		return hosts, true, nil
	}
	f := path.Join(d.dir, ip)
	if !io.FileExists(f) {
		return []string{}, false, nil
	}
	c, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, false, err
	}
	res := slice.Strings(string(c), d.sep)
	// If we're forcing to resolve empties, don't add empty lists of hosts from disk
	// until we write them in this session
	if d.forceResolveEmpties && len(res) == 0 {
		return []string{}, false, nil
	}
	d.cache[ip] = res
	return res, true, nil
}

func (d *ipCache) Put(ip string, hosts []string) error {
	f := path.Join(d.dir, ip)
	c := strings.Join(hosts, d.sep)
	if err := ioutil.WriteFile(f, []byte(c), 0755); err != nil {
		return err
	}
	d.cacheMu.Lock()
	defer d.cacheMu.Unlock()
	d.cache[ip] = hosts
	return nil
}

//go:generate genopts --function ResolveIPs verbose threads:int:10 useCache forceResolveEmpties
func ResolveIPs(recs []Record, optss ...ResolveIPsOption) ([]Record, error) {
	opts := MakeResolveIPsOptions(optss...)
	threads := opts.Threads()
	verbose := opts.Verbose()
	useCache := opts.UseCache()
	forceResolveEmpties := opts.ForceResolveEmpties()

	start := time.Now()
	if verbose {
		log.Printf("start resolving ips")
	}

	ipSet := map[string]bool{}
	for _, rec := range recs {
		ipSet[rec.IP] = true
	}

	ips := make(chan string)
	go func() {
		for ip := range ipSet {
			ips <- ip
		}
		close(ips)
	}()

	type resolvedIP struct {
		ip    string
		hosts []string
	}
	resolvedIPs := make(chan resolvedIP)
	var dc *ipCache
	if useCache {
		d, err := newDiskCache(forceResolveEmpties)
		if err != nil {
			return nil, err
		}
		dc = d
	}
	go func() {
		var wg sync.WaitGroup
		wg.Add(threads)
		for i := 0; i < threads; i++ {
			go func() {
				defer wg.Done()
				for ip := range ips {
					if dc != nil {
						hosts, found, err := dc.Get(ip)
						if err != nil {
							log.Printf("error getting ip from cache: %v", err)
						} else if found {
							resolvedIPs <- resolvedIP{ip, hosts}
							continue
						}
					}
					hosts, err := net.LookupAddr(ip)
					if verbose {
						log.Printf("resolved IP %s -> %+v, err (%v)", ip, hosts, err)
					}
					if dc != nil {
						if err := dc.Put(ip, hosts); err != nil {
							log.Printf("error persisting ip: %s and hosts: %v", ip, hosts)
						}
					}
					if err != nil {
						log.Printf("error resolving IP %s: %v", ip, err)
					} else {
						resolvedIPs <- resolvedIP{ip, hosts}
					}
				}
			}()
		}
		wg.Wait()
		close(resolvedIPs)
	}()

	resolvedIPsMap := map[string][]string{}
	resolvedIPsMapDone := make(chan bool, 1)
	go func() {
		for resolvedIP := range resolvedIPs {
			resolvedIPsMap[resolvedIP.ip] = resolvedIP.hosts
		}
		resolvedIPsMapDone <- true
	}()
	<-resolvedIPsMapDone

	var res []Record
	for _, rec := range recs {
		if hosts, ok := resolvedIPsMap[rec.IP]; ok {
			rec.ResolvedHosts = hosts
		}
		res = append(res, rec)
	}

	if verbose {
		log.Printf("resolved ips in %s", time.Since(start))
	}

	return res, nil
}
