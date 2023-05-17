package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

func getCN(addr string, CN string, timeout int, ch chan string, stCh chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}
	dialer := net.Dialer{Timeout: time.Duration(timeout) * time.Second}
	conn, err := tls.DialWithDialer(&dialer, "tcp", addr+":443", conf)
	stCh <- addr
	if err != nil {
		return
	}
	defer conn.Close()
	certs := conn.ConnectionState().PeerCertificates

	if strings.Contains(certs[0].Subject.CommonName, CN) {
		ch <- addr
	}
}

func hostPrint(ch chan string) {
	for data := range ch {
		fmt.Println(data)
	}
}

func statusPrint(ch chan string, total int) {
	currentScanned := 1
	for data := range ch {
		currentScanned++
		print("Addresses scanned so far: ", currentScanned, "/", total, " (", data, ")", "\r")
	}
}

func main() {
	ASPtr := flag.String("asn", "203214", "The ASN to pull the addresses from")
	CNPtr := flag.String("commonName", "*.googlevideo.com", "The Common Name to search for in the TLS certificate")
	ratePtr := flag.Int64("rate", 100, "How many addresses to scan per second")
	timeoutPtr := flag.Int("timeout", 10, "How many seconds to wait for each host to respond")
	flag.Parse()

	addresses := getASNRanges(*ASPtr)
	ipCh := make(chan string)
	stCh := make(chan string)

	go hostPrint(ipCh)
	go statusPrint(stCh, len(addresses))

	var wg sync.WaitGroup
	for _, addresses := range addresses {
		time.Sleep(time.Second / time.Duration(*ratePtr))
		wg.Add(1)
		go getCN(addresses, *CNPtr, *timeoutPtr, ipCh, stCh, &wg)
	}
	wg.Wait()
}
