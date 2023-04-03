package httplogs

type Record struct {
	IP            string
	ResolvedHosts []string
	Method        string
	Path          string
	UserAgent     string
	StatusCode    int
}
