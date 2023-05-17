# CDN-Sniper
A program to scan all the IPs of an ASN and return ones matching a commonName

# Purpose of CDN-Sniper
I wanted a way to get all CDN edge nodes of a specific ASN provider. For example. I wanted the GGC nodes of AS203214 so I can generate a list of addresses and use it in my router & firewall. This works with any ASN and commonName and can be used to identify virtually anything.  

# Build
```go build -o CDN-Sniper *.go```

Nothing much to it. This project requires no external libraries. Everything is contained within Go's standard libraries
# Usage
```
  -asn string
    	The ASN to pull the addresses from (default "203214")
  -commonName string
    	The Common Name to search for in the TLS certificate (default "*.googlevideo.com")
  -rate int
    	How many addresses to scan per second (default 100)
  -timeout int
    	How many seconds to wait for each host to respond (default 10)
```
