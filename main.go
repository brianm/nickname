package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/brianm/mdns"
)

var helpText = `Usage: mdnsp NAME IP

Publishes NAME.local pointing to IP.

So to create a record for something.local. pointing to 127.3.4.5,
which would look like "something.local. 5 IN A 127.3.4.5", you
would run the command as:

  $ mdnsp something 127.3.4.5
`
var help = false

func init() {
	flag.BoolVar(&help, "h", false, "Show help")
	flag.Parse()
}

func main() {
	if help || len(os.Args) != 3 {
		fmt.Println(helpText)
		os.Exit(1)
	}

	name := os.Args[1]
	ip := os.Args[2]

	record := fmt.Sprintf("%s.local. 5 IN A %s", name, ip)
	err := mdns.Publish(record)
	if err != nil {
		panic(err)
	}

	select {}
}
