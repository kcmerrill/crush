package core

import (
	"net"

	log "github.com/sirupsen/logrus"
)

// UDP starts a udp sever for incoming requests
func (s *Sherlock) UDP(port string) {
	server, serverErr := net.ResolveUDPAddr("udp", ":"+port)
	if serverErr != nil {
		log.WithFields(log.Fields{"type": "udp", "port": port}).Error("Unable to start server")
	}
	conn, err := net.ListenUDP("udp", server)
	if err != nil {
		log.Fatal(err.Error())
		return
	} else {
		defer conn.Close()
	}
	defer conn.Close()

	buf := make([]byte, 1024)

	for {
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.WithFields(log.Fields{"type": "udp", "port": port}).Error(err.Error())
		} else {
			go s.Process(string(buf[0:n]), "|")
		}
	}
}
