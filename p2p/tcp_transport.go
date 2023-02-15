package p2p

import (
	"fmt"
	"net"
	"sync"
)

// Rappresenta quando la connessione TCP è stabilita con il nodo
type TCPPeer struct {
	//Sottostante la connessione del peer
	conn net.Conn
	//se dial(componi) una connesione => outbound == true
	//se accetti una connesione => outbound == false
	outbound bool
}

type TCPTransport struct {
	listenAddress   string
	listener        net.Listener
	shankeHands     HandShakeFunc
	decoder 		Decoder

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

// Crea un nuvo TCP peer
func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

// Crea un nupvp transport TCP
func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		shankeHands: NOPHandshakeFunc,
		listenAddress: listenAddr,
	}
}

// Qui rimane in ascolto
func (t *TCPTransport) ListenAndAccept() error {

	var err error

	t.listener, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()

	return nil
}

// Looppa l'accettazione della connesione
func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}
		fmt.Printf("new incoming connection %+v\n", conn)
		go t.handleConn(conn)
	}

}

type Temp struct {}


// Gestisce la nuove connessioni
func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)

	if err := t.shankeHands(conn); err != nil {
		conn.
	}


	/*
		lenDecodeError := 0

		lenDecodeError ++ 

		if lenDecodeError == 5 {
			conn.Close()
		}
		Filtro antispam

	*/
	//Loop di lettura 
	msg := &Temp{}
	for {
		if err := t.Decoder.Decode(conn, msg ); err != nil {
			fmt.Printf("TPC error: %s\n", err)
			continue 
		}
	}
	
}
