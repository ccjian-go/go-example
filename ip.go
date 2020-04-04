package main

import "net"

func ClientIP() (ip string) {
	//获取所有网卡
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		ip = ""
	}
	for _, value := range addrs{
		if ipnet, ok := value.(*net.IPNet); ok && !ipnet.IP.IsLoopback(){
			if ipnet.IP.To4() != nil{
				ip = ipnet.IP.String()
			}
		}
	}
	return ip
}
