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

type TCPTransportOpts struct {
	ListenAddr    string
	HandShakeFunc HandShakeFunc
	Decoder       Decoder
}

type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener

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
func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
	}
}

// Qui rimane in ascolto
func (t *TCPTransport) ListenAndAccept() error {

	var err error

	t.listener, err = net.Listen("tcp", t.ListenAddr)
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

type Temp struct{}

// Gestisce la nuove connessioni
func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)

	if err := t.HandShakeFunc(peer); err != nil {
		conn.Close()
		fmt.Printf("TPC handshake error: %s\n", err)
		return
	}

	/*
		lenDecodeError := 0

		lenDecodeError ++

		if lenDecodeError == 5 {
			conn.Close()
		}
		Filtro antispam



		buf := make([]bytes, 2000)->Fuori dal for

		nel for
		n, err := conn.Read(buf)

		if err != nil {fmt.Printf("TPC error: %s\n", err)}

	*/
	//Loop di lettura
	msg := &Message{}
	for {

		if err := t.Decoder.Decode(conn, msg); err != nil {
			fmt.Printf("TPC error: %s\n", err)
			continue
		}
		msg.From = conn.RemoteAddr()
		fmt.Printf("Message: %+v\n", msg)
	}

}
