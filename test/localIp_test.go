package main

import (
	"log"
	"net"
	"testing"
)

func TestLocalIP(t *testing.T) {
	addrs, err := net.InterfaceAddrs()
	log.Println(addrs)
	if err != nil {
		log.Println(err)
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				log.Println(ipnet.IP.String())
			}
		}
	}
}
