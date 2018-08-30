package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

const defaultRedirect = "https://google.com"
const usage = `Usage of %s:

This simple utility creates a HTTP server that always redirect to the
URL given as parameter. If no parameter is given, it redirects to https://google.com
by default.

Example:

	%s http://google.com

Parameters:

`

var url string = "https://google.com"

func init() {
	flag.Usage = func() {
		c := os.Args[0]
		fmt.Fprintf(os.Stderr, usage, c, c)
		flag.PrintDefaults()
	}
}

type Flags struct {
	code int
	port int
}

func main() {
	var flags Flags
	flag.IntVar(&flags.code, "code", 302, "Status code of the response")
	flag.IntVar(&flags.port, "port", 8080, "Web server listening port")

	flag.Parse()

	switch flag.NArg() {
	case 0:
		url = defaultRedirect
	case 1:
		url = flag.Arg(0)
	default:
		fmt.Printf("Wrong number of argument.  Use -h for help.")
		os.Exit(1)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, url, flags.code)
	})

	fmt.Println("Server listening on port", flags.port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", flags.port), nil))
}
