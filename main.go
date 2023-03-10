package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/Jxancestral17/DeceStore/p2p"
)

func makeServer(listenAddr string, nodes ...string) *FileServer {
	tcptransporOpts := p2p.TCPTransportOpts{
		ListenAddr:    listenAddr,
		HandShakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}

	tcpTransport := p2p.NewTCPTransport(tcptransporOpts)

	fileServerOpts := FileServerOpts{
		EncKey:            newEncryptionKey(),
		StorageRoot:       listenAddr + "_network",
		PathTransformFunc: CASPathTransformFunc,
		Transport:         tcpTransport,
		BootstrapNodes:    nodes,
	}

	s := NewFileServer(fileServerOpts)

	tcpTransport.OnPeer = s.OnPeer

	return s

}

func main() {
	s1 := makeServer(":3000", "")
	s2 := makeServer(":4000", ":3000")
	s3 := makeServer(":5000", ":3000", ":4000")

	go func() {
		log.Fatal(s1.Start())

	}()
	time.Sleep(2 * time.Second)

	go s2.Start()
	time.Sleep(500 * time.Millisecond)

	go s3.Start()

	time.Sleep(2 * time.Second)

	key := "file.jpg"
	data := bytes.NewReader([]byte("test"))

	s2.Store(key, data)

	// if err := s2.store.Delete(key); err != nil {
	// 	log.Fatal(err)
	// }

	r, err := s2.Get(key)
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

	//select {}
}
