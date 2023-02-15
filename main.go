package main

import (
	"log"

	"github.com/Jxancestral17/DeceStore/p2p"
)

func main() {
	listenAddr := ":4000"
	tr := p2p.NewTCPTransport(listenAddr)

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
