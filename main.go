package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
    proxies := getProxies()

    fmt.Printf("Refreshed proxies at %s, number of proxies: %d\n", time.Now().Format(time.RFC3339), len(proxies))

    writeToFile(proxies)

    ticker := time.NewTicker(3 * time.Hour)

    for {
        select {
        case <-ticker.C:
            proxies = getProxies()

            fmt.Printf("Refreshed proxies at %s, number of proxies: %d\n", time.Now().Format(time.RFC3339), len(proxies))

            writeToFile(proxies)
        }
    }
}

func getProxies() []string {
	var proxies []string

	proxies = append(proxies, getProxiesFromWebsite("https://api.proxyscrape.com/?request=getproxies&proxytype=http&timeout=10000&country=all&ssl=all&anonymity=all")...)

	return proxies
}

func getProxiesFromWebsite(url string) []string {
	// create a slice of strings to store the proxies
	var proxies []string

	// make a get request to the url
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	// read the body of the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// convert the body to a string
	bodyString := string(body)

	// split the body string by a new line
	bodySplit := strings.Split(bodyString, "\n")

	// add the proxies to the slice
	for _, proxy := range bodySplit {
		proxies = append(proxies, proxy)
	}

	return proxies
}

func writeToFile(proxies []string) {
	file, err := os.Create("proxies.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, proxy := range proxies {
		file.WriteString(proxy + "\n")
	}

	file.Close()
}
