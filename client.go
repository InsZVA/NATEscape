package main

import (
	"bytes"
	"flag"
	"log"
	"nat/util"
	"net"
	"time"
)

var hostport = flag.String("hostport", "80", "hostpost")
var remoteaddr = flag.String("remoteaddr", "ramnode.inszva.com:3478", "remoteaddr")
var content = flag.String("content", "", "content")
var name = flag.String("name", "88888888", "name(must size of 8)")
var target = flag.String("target", "77777777", "name(must size of 8)")

func main() {
	flag.Parse()
	raddr, err := net.ResolveUDPAddr("udp4", *remoteaddr)
	if err != nil {
		log.Fatalln(err)
	}
	laddr, err := net.ResolveUDPAddr("udp4", "0.0.0.0:"+*hostport)
	if err != nil {
		log.Fatalln(err)
	}
send:
	log.Println("I'm tring to tell the server my address")
	conn, err := net.DialUDP("udp4", laddr, raddr)
	defer conn.Close()
	if err != nil {
		log.Fatalln(err)
	}
	msg := util.Msg{
		Type: util.REGISTER_NAME,
		Name: util.Bytes2array([]byte(*name)),
	}
	conn.SetDeadline(time.Now().Add(time.Second))
	_, err = conn.Write(util.Msg2Bytes(&msg))
	if err != nil {
		log.Printf("send timeout, I'm tring to retry")
		goto send
	}
	conn.SetDeadline(time.Now().Add(time.Second))
	buff := make([]byte, 32)
	_, err = conn.Read(buff)
	if err != nil {
		log.Printf("read timeout, I'm tring to retry")
		goto send
	}
	log.Println("server returns:", string(buff))
	log.Println("I am waiting 5 seconds for his starting...")
	time.Sleep(5 * time.Second)

	//find

	log.Println("I'm tring to question the server his address")
	msg.Type = util.SEARCH_NAME
	msg.Name = util.Bytes2array([]byte(*target))
	conn.SetDeadline(time.Now().Add(time.Second))
	_, err = conn.Write(util.Msg2Bytes(&msg))
	if err != nil {
		log.Printf("send timeout, I'm tring to retry")
		goto send
	}
	conn.SetDeadline(time.Now().Add(time.Second))
	_, err = conn.Read(buff)
	if err != nil {
		log.Printf("read timeout, I'm tring to retry")
		goto send
	}
	log.Println("sever told me his address is:", string(buff))
	buff = bytes.Split(buff, []byte{0})[0]
	//Start

	realAddr, err := net.ResolveUDPAddr("udp4", string(buff))
	if err != nil {
		log.Fatalln(err)
	}
	conn.Close()
	conn2, err := net.DialUDP("udp4", laddr, realAddr)
	defer conn2.Close()
	if err != nil {
		log.Fatalln(err)
	}
	conn2.Write([]byte(*content))
	log.Println("I send a message to him, but it is not sure he can see")
	log.Println("Now I am reading...")
	conn2.Read(buff)
	log.Println("I succeeded reading:", string(buff))
}
