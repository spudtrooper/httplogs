package httplogs

import (
	"log"
	"regexp"
	"strconv"
	"time"

	"github.com/spudtrooper/goutil/io"
)

var (
	// 216.131.114.61 - - [31/Mar/2023:05:18:27 -0600] "GET /userscripts.html HTTP/1.1" 404 315 "http://www.jeffpalm.com" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 Safari/537.36" www.jeffpalm.com 162.241.225.132
	ipRE = regexp.MustCompile(`^(\d+\.\d+\.\d+\.\d+)`)

	// "GET /userscripts.html HTTP/1.1" 200
	methodRE = regexp.MustCompile(`"(GET|POST|HEAD|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH) (\S*) HTTP/\d+\.\d+" (\d+)`)

	// "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) HeadlessChrome/112.0.5614.0 Safari/537.36" jeffpalm.com 162.241.225.132
	userAgentRE1 = regexp.MustCompile(`"https?://[^"]+.com[^"]*" "([^"]+)"'`)
	userAgentRE2 = regexp.MustCompile(`"\-" "([^"]+)"'`)
	userAgentRE3 = regexp.MustCompile(`"([^"]+)" \S+ \d+\.\d+\.\d+\.\d+$`)
	// 301 229 "-" "\\\"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.0 Safari/605.1.15\\\""
	userAgentRE4 = regexp.MustCompile(`\d+ \d+ "?\-?"?\s*(".*") \S+ \d+\.\d+\.\d+\.\d+$`)
)

//go:generate genopts --function Parse verbose
func Parse(filepaths []string, optss ...ParseOption) (chan Record, chan error, error) {
	opts := MakeParseOptions(optss...)
	verbose := opts.Verbose()

	recs := make(chan Record)
	errs := make(chan error)

	ss, stringsFromFilesErrors, err := io.StringsFromFiles(filepaths)
	if err != nil {
		return nil, nil, err
	}
	go func() {
		for err := range stringsFromFilesErrors {
			errs <- err
		}
	}()

	start := time.Now()
	if verbose {
		log.Printf("start parsing")
	}

	go func() {
		for s := range ss {
			if s == "" {
				continue
			}
			var rec Record
			{
				m := ipRE.FindStringSubmatch(s)
				if m == nil {
					log.Printf("no ip: %s", s)
				} else {
					rec.IP = m[1]
				}
			}
			{
				m := methodRE.FindStringSubmatch(s)
				if m == nil {
					log.Printf("no method: %s", s)
				} else {
					rec.Method = m[1]
					rec.Path = m[2]
					httpCode, err := strconv.Atoi(m[3])
					if err != nil {
						errs <- err
					} else {
						rec.StatusCode = httpCode
					}
				}
			}
			{
				m := userAgentRE1.FindStringSubmatch(s)
				if m == nil || len(m) < 2 {
					m = userAgentRE2.FindStringSubmatch(s)
				}
				if m == nil || len(m) < 2 {
					m = userAgentRE3.FindStringSubmatch(s)
				}
				if m == nil || len(m) < 2 {
					m = userAgentRE4.FindStringSubmatch(s)
				}
				if m == nil || len(m) < 2 {
					log.Printf("no user agent: %s", s)
				} else {
					rec.UserAgent = m[1]
				}
			}
			recs <- rec
		}
		close(recs)
		close(errs)
	}()

	if verbose {
		log.Printf("done parsing in %v", time.Since(start))
	}

	return recs, errs, nil

}
