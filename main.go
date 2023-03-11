package main

import (
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

	go func() {
		log.Fatal(s1.Start())
	}()
	time.Sleep(2 * time.Second)

	go s2.Start()

	time.Sleep(2 * time.Second)

	// data := bytes.NewReader([]byte("test"))

	// s2.Store(fmt.Sprintf("privatekey_%d", i), data)
	// time.Sleep(time.Millisecond * 5)

	r, err := s2.Get("foo")
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
