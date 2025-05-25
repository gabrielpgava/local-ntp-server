package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

const (
	ntpEpochOffset = 2208988800
	port           = ":123"
	logFileName    = "ntp_access.log"
)

var logger *log.Logger

func initLogger() {
	f, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}
	logger = log.New(f, "", log.LstdFlags)
}

func main() {
	initLogger()

	addr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		log.Fatalf("failed to resolve UDP addr: %v", err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer conn.Close()

	log.Printf("NTP server listening on %s", port)

	for {
		buf := make([]byte, 48)
		n, clientAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Printf("error reading: %v", err)
			continue
		}
		if n < 48 {
			log.Printf("received short packet from %v", clientAddr)
			continue
		}
		go handleNTPRequest(conn, clientAddr, buf)
	}
}

func handleNTPRequest(conn *net.UDPConn, addr *net.UDPAddr, req []byte) {
	resp := make([]byte, 48)
	resp[0] = 0x1C

	now := time.Now().UTC()
	secs := uint32(now.Unix() + ntpEpochOffset)
	frac := uint32(uint64(now.Nanosecond()) * (1 << 32) / 1e9)

	// Reference Timestamp
	binary.BigEndian.PutUint32(resp[16:], secs)
	binary.BigEndian.PutUint32(resp[20:], frac)
	// Originate Timestamp: from request (T1)
	copy(resp[24:], req[40:48])
	// Receive Timestamp (T2)
	binary.BigEndian.PutUint32(resp[32:], secs)
	binary.BigEndian.PutUint32(resp[36:], frac)
	// Transmit Timestamp (T3)
	binary.BigEndian.PutUint32(resp[40:], secs)
	binary.BigEndian.PutUint32(resp[44:], frac)

	_, err := conn.WriteToUDP(resp, addr)
	if err != nil {
		log.Printf("error sending response to %v: %v", addr, err)
	}

	go logRequest(addr)
}

func logRequest(addr *net.UDPAddr) {
	ip := addr.IP.String()
	names, err := net.LookupAddr(ip)
	hostname := "-"
	if err == nil && len(names) > 0 {
		hostname = names[0]
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	logger.Printf("Hora consultada em %s por %s (hostname: %s)\n", now, ip, hostname)
	go fmt.Printf("Hora consultada em %s por %s (hostname: %s)\n", now, ip, hostname)
}