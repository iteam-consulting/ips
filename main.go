package main

import (
	"flag"
	"fmt"
	"net"
)

var cidr string

func init() {
	flag.StringVar(&cidr, "cidr", "0.0.0.0/24", "Provide an IP/4 CIDR address")
}

func main() {
	flag.Parse()
	ip, nw, err := net.ParseCIDR(cidr)

	if err != nil {
		fmt.Print(err)
		return
	}

	var ips []string
	for ip := ip.Mask(nw.Mask); nw.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
