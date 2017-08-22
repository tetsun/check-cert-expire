package main

import (
	"crypto/tls"
	"time"
)

func Dial(srvName string, hostName string, port string, isInsecure bool) (*time.Time, error) {

	tlsCfg := &tls.Config{
		ServerName:         srvName,
		InsecureSkipVerify: isInsecure,
	}

	conn, err := tls.Dial("tcp", hostName+":"+port, tlsCfg)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	state := conn.ConnectionState()
	certs := state.PeerCertificates
	expire := &certs[0].NotAfter

	return expire, nil
}
