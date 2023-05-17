# CDN-Sniper
A program to scan all the IPs of an ASN and return ones matching a commonName

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
