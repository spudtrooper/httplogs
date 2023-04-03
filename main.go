package main

import (
	"flag"
	"log"

	"github.com/spudtrooper/goutil/check"
	"github.com/spudtrooper/httplogs/httplogs"
)

var (
	resolve_ip  = flag.Bool("resolve_ip", false, "resolve ip")
	nosummary   = flag.Bool("nosummary", false, "don't print summary")
	byip        = flag.Bool("byip", false, "group by ip")
	bypath      = flag.Bool("bypath", false, "group by path")
	byuseragent = flag.Bool("byuseragent", false, "group by user agent")
)

func realMain() error {
	recs, err := httplogs.Parse(flag.Args(), httplogs.ParseResolveIPsFlag(resolve_ip))
	if err != nil {
		return err
	}
	grouped := httplogs.Group(recs)

	if *nosummary {
		return nil
	}

	if *bypath {
		log.Println()
		log.Println("Paths")
		for _, it := range grouped.ByPath {
			log.Printf("%7d: %s", len(it.Recs), it.Key)
		}
	}

	if *byip {
		log.Println()
		log.Println("IPs")
		for _, it := range grouped.ByIP {
			ip := it.Key
			if *resolve_ip {
				var hostnames []string
				if len(it.Recs) > 0 {
					hostnames = it.Recs[0].ResolvedHosts
				}
				log.Printf("%7d: %s (%s)", len(it.Recs), ip, hostnames)
			} else {
				log.Printf("%7d: %s", len(it.Recs), ip)
			}
		}
	}

	if *byuseragent {
		log.Println()
		log.Println("UserAgents")
		for _, it := range grouped.ByUserAgent {
			log.Printf("%7d: %s", len(it.Recs), it.Key)
		}
	}

	return nil
}

func main() {
	flag.Parse()
	check.Err(realMain())
}