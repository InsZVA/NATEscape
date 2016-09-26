package main

import (
	"log"
	"nat/util"
	"net"
)

func main() {
	nameMap := make(map[string]string)
	laddr, err := net.ResolveUDPAddr("udp4", "0.0.0.0:3478")
	if err != nil {
		log.Fatalln(err)
	}
	conn, err := net.ListenUDP("udp4", laddr)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		buff := make([]byte, 16)
		_, raddr, _ := conn.ReadFromUDP(buff)
		log.Println(raddr)
		msg := util.Bytes2Msg(buff)
		log.Println(msg)
		if msg.Type == util.REGISTER_NAME {
			name := util.Array2bytes(msg.Name)
			nameMap[string(name)] = raddr.String()
			conn.WriteToUDP([]byte(raddr.String()), raddr)
		} else {
			name := util.Array2bytes(msg.Name)
			if addr, ok := nameMap[string(name)]; ok {
				log.Println("found name")
				conn.WriteToUDP([]byte(addr), raddr)
			} else {
				log.Println("not found")
				conn.WriteToUDP([]byte("unkown"), raddr)
			}
		}
	}

}
