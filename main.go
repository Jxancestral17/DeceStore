package main

import (
	"log"

	"github.com/Jxancestral17/DeceStore/p2p"
)

func main() {
	tcptTransporOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandShakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}

	tcpTransport := p2p.NewTCPTransport(tcptTransporOpts)

	fileServerOpts := FileServerOpts{
		StorageRoot:       "3000_network",
		PathTransformFunc: CASPathTransformFunc,
		Transport:         tcpTransport,
	}
	s := NewFileServer(fileServerOpts)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
	select {}
}
