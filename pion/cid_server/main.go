package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/pion/dtls/v2"
)

const (
	dtlsServerPortEnv      = "DTLS_INTEROP_SERVER_PORT"
	dtlsServerCIDLengthEnv = "DTLS_INTEROP_SERVER_CID_LENGTH"
	dtlsClienPSKEnv        = "DTLS_INTEROP_CLIENT_PSK"
)

func main() {
	port := 5684
	if p, ok := os.LookupEnv(dtlsServerPortEnv); ok {
		i, err := strconv.ParseInt(p, 10, 64)
		if err != nil {
			panic(err)
		}
		port = int(i)
	}
	cidLen := 8
	if l, ok := os.LookupEnv(dtlsServerCIDLengthEnv); ok {
		i, err := strconv.ParseInt(l, 10, 64)
		if err != nil {
			panic(err)
		}
		cidLen = int(i)
	}
	psk := "secretPSK"
	if p, ok := os.LookupEnv(dtlsClienPSKEnv); ok {
		psk = p
	}
	addr := &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: port}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	config := &dtls.Config{
		PSK: func(hint []byte) ([]byte, error) {
			return []byte(psk), nil
		},
		CipherSuites: []dtls.CipherSuiteID{dtls.TLS_PSK_WITH_AES_128_CCM_8},
		ConnectContextMaker: func() (context.Context, func()) {
			return context.WithTimeout(ctx, 30*time.Second)
		},
		ConnectionIDGenerator: dtls.RandomCIDGenerator(cidLen),
	}

	listener, err := dtls.Listen("udp", addr, config)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := listener.Close(); err != nil {
			panic(err)
		}
	}()

	fmt.Println("Listening...")

	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}

	b := make([]byte, 8192)
	for {
		n, err := conn.Read(b)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Echoing received message: %s\n", string(b[:n]))
		if _, err := conn.Write(b[:n]); err != nil {
			panic(err)
		}
	}
}
