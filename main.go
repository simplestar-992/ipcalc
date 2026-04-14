package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {
	ip := flag.String("i", "192.168.1.0", "IP address")
	mask := flag.Int("m", 24, "CIDR mask")
	flag.Parse()
	
	ipAddr := net.ParseIP(*ip)
	ipNet := net.IPNet{IP: ipAddr, Mask: net.CIDRMask(*mask, 32)}
	
	fmt.Printf("IP:      %s\n", ipAddr)
	fmt.Printf("Mask:    %s\n", net.IP(ipNet.Mask))
	fmt.Printf("Network: %s\n", ipNet.String())
}
