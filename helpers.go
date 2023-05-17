package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/netip"
	"strings"
)

func Hosts(cidr string) ([]string, error) {

	prefix, err := netip.ParsePrefix(cidr)
	if err != nil {
		panic(err)
	}

	var ips []string
	for addr := prefix.Addr(); prefix.Contains(addr); addr = addr.Next() {
		ips = append(ips, addr.String())
	}

	if len(ips) < 2 {
		return ips, nil
	}

	return ips, nil
}

func getASNRanges(asn string) []string {
	resp, err := http.Get("https://stat.ripe.net/data/announced-prefixes/data.json?resource=" + asn)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)

	var result map[string]any
	json.Unmarshal([]byte(sb), &result)

	prefixes := result["data"].(map[string]any)["prefixes"].([]interface{})

	hostMap := make(map[string]bool)

	for _, prefix := range prefixes {
		prefix := prefix.(map[string]any)["prefix"].(string)

		if !strings.Contains(prefix, ":") {
			hosts, err := Hosts(prefix)
			if err != nil {
				panic(err)
			}

			for _, host := range hosts {
				hostMap[host] = false
			}
		}
	}

	var totalHosts []string

	for host := range hostMap {
		totalHosts = append(totalHosts, host)
	}

	return totalHosts
}
