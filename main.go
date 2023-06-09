package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/spudtrooper/goutil/check"
	"github.com/spudtrooper/goutil/slice"
	"github.com/spudtrooper/httplogs/httplogs"
)

var (
	resolveips           = flag.Bool("resolveips", false, "resolve ip")
	useipcache           = flag.Bool("useipcache", false, "use ip cache")
	forceresolveemptyips = flag.Bool("forceresolveemptyips", false, "don't use disk cache for ips that don't resolve to hosts")
	nosummary            = flag.Bool("nosummary", false, "don't print summary")
	byip                 = flag.Bool("byip", false, "group by ip")
	bypath               = flag.Bool("bypath", false, "group by path")
	byuseragent          = flag.Bool("byuseragent", false, "group by user agent")
	bystatuscode         = flag.Bool("bystatuscode", false, "group by status code")
	filterstatuscodes    = flag.String("filterstatuscodes", "", "status codes to filter")
	filterpath           = flag.String("filterpath", "", "path regexp to filter")
	filteroutpath        = flag.String("filteroutpath", "", "path regexp to filter out")
	filteruseragent      = flag.String("filteruseragent", "", "user agent regexp to filter")
	filteroutuseragent   = flag.String("filteroutuseragent", "", "user agent regexp to filter out")
	verboseresolve       = flag.Bool("verboseresolve", false, "verbose resolve")
	verboseparse         = flag.Bool("verboseparse", false, "verbose parse")
)

func realMain() {
	files := flag.Args()

	var recs []httplogs.Record
	var recsCh <-chan httplogs.Record
	var errs <-chan error
	if len(files) == 0 || (len(files) == 1 && files[0] == "-") {
		r, es, err := httplogs.ParseFromStdin(httplogs.ParseVerboseFlag(verboseparse))
		check.Err(err)
		recsCh = r
		errs = es
	} else {
		r, es, err := httplogs.ParseFiles(files, httplogs.ParseVerboseFlag(verboseparse))
		check.Err(err)
		recsCh = r
		errs = es
	}
	go func() {
		for err := range errs {
			log.Printf("error parsing: %s", err)
		}
	}()
	recsDone := make(chan bool, 1)
	go func() {
		for rec := range recsCh {
			recs = append(recs, rec)
		}
		recsDone <- true
	}()
	<-recsDone

	if len(recs) == 0 {
		log.Fatalf("no records found after parse")
		return
	}

	statusCodes, err := slice.Ints(*filterstatuscodes)
	check.Err(err)

	recs = httplogs.Filter(recs,
		httplogs.FilterStatusCodes(statusCodes),
		httplogs.FilterPathFilterFlag(filterpath),
		httplogs.FilterNegPathFilterFlag(filteroutpath),
		httplogs.FilterUserAgentFilterFlag(filteruseragent),
		httplogs.FilterNegUserAgentFilterFlag(filteroutuseragent),
	)

	if len(recs) == 0 {
		log.Fatalf("no records found after filtering")
		return
	}

	if *resolveips {
		resolvedRecs, err := httplogs.ResolveIPs(recs,
			httplogs.ResolveIPsVerboseFlag(verboseresolve),
			httplogs.ResolveIPsUseCacheFlag(useipcache),
			httplogs.ResolveIPsForceResolveEmptiesFlag(forceresolveemptyips),
		)
		check.Err(err)
		recs = resolvedRecs
	}

	grouped := httplogs.Group(recs)

	if *nosummary {
		return
	}

	byPath := *bypath
	if !*byip && !*bypath && !*byuseragent && !*bystatuscode {
		byPath = true
	}

	if byPath {
		fmt.Println()
		fmt.Println("Paths")
		for _, it := range grouped.ByPath {
			fmt.Printf("%7d: %s\n", len(it.Recs), it.Key)
		}
	}

	if *byip {
		fmt.Println()
		fmt.Println("IPs")
		for _, it := range grouped.ByIP {
			ip := it.Key
			if *resolveips {
				var hostnames []string
				if len(it.Recs) > 0 {
					hostnames = it.Recs[0].ResolvedHosts
				}
				fmt.Printf("%7d: %-50s (%s)\n", len(it.Recs), strings.Join(hostnames, ","), ip)
			} else {
				fmt.Printf("%7d: %s\n", len(it.Recs), ip)
			}
		}
	}

	if *byuseragent {
		fmt.Println()
		fmt.Println("UserAgents")
		for _, it := range grouped.ByUserAgent {
			fmt.Printf("%7d: %s\n", len(it.Recs), it.Key)
		}
	}

	if *bystatuscode {
		fmt.Println()
		fmt.Println("StatusCodes")
		for _, it := range grouped.ByStatusCode {
			fmt.Printf("%7d: %d\n", len(it.Recs), it.Key)
		}
	}
}

func main() {
	flag.Parse()
	realMain()
}
