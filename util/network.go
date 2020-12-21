package util

import (
	"log"
	"net"
)

func GetNid() (string, error) {
	if ifaces, err := net.Interfaces(); err != nil {
		log.Printf("Error : %s", err.Error())
		return "", err
	} else {
		for _, i := range ifaces {
			if addrs, err := i.Addrs(); err != nil {
				log.Printf("Error : %s", err.Error())
				return "", err
			} else {
				for _, addr := range addrs {
					var ip net.IP
					switch v := addr.(type) {
					case *net.IPNet:
							ip = v.IP
					case *net.IPAddr:
							ip = v.IP
					}
					return ip.String(), nil
				}
				return "", err
			}
		}
		return "", err
	}
}