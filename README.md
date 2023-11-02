# ProxyAPI

Just a simple script to retrieve proxies from multiple websites and dump them to a file then refresh them every 3 hours.

# Usage 

You can use something like `pm2` to keep the script running at all times
To run simply do: `go run main.go` or to build and run do: `go build .`, `chmod 777 *` and `./ProxyAPI`

# Misc

To add more websites to dump from simply copy and paste the following line:
`proxies = append(proxies, getProxiesFromWebsite("https://api.proxyscrape.com/?request=getproxies&proxytype=http&timeout=10000&country=all&ssl=all&anonymity=all")...)`
and just change the website URL
