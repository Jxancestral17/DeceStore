package p2p

import (
	"errors"
	"fmt"
	"log"
	"net"
)

// Rappresenta quando la connessione TCP è stabilita con il nodo
type TCPPeer struct {
	//Sottostante la connessione del peer
	conn net.Conn
	//se dial(componi) una connesione => outbound == true
	//se accetti una connesione => outbound == false
	outbound bool
}

// Crea un nuvo TCP peer
func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

// Chiusura dell'interfaccia del Peer
func (p *TCPPeer) Close() error {
	return p.conn.Close()
}

type TCPTransportOpts struct {
	ListenAddr    string
	HandShakeFunc HandShakeFunc
	Decoder       Decoder
	OnPeer        func(Peer) error
}

type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener
	rpcch    chan RPC
}

// Crea un nupvp transport TCP
func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
		rpcch:            make(chan RPC),
	}
}

/*
Cosume impremente l'interfaccia trapsort e
ritorno la sola luttare del channel
per leggere il messagio in arrivo da un altro peer della rete
*/
func (t *TCPTransport) Consume() <-chan RPC {
	return t.rpcch
}

// Chiude il transport interface
func (t *TCPTransport) Close() error {
	return t.listener.Close()
}

// Qui rimane in ascolto
func (t *TCPTransport) ListenAndAccept() error {

	var err error

	t.listener, err = net.Listen("tcp", t.ListenAddr)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()

	log.Printf("TCP transport listening on port: %s\n", t.ListenAddr)

	return nil
}

// Looppa l'accettazione della connesione
func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()

		if errors.Is(err, net.ErrClosed) {
			return
		}

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
	var err error
	defer func() {
		fmt.Printf("Dropping perr connection: %s", err)
		conn.Close()
	}()
	peer := NewTCPPeer(conn, true)

	if err = t.HandShakeFunc(peer); err != nil {
		//conn.Close()
		//fmt.Printf("TPC handshake error: %s\n", err)
		return
	}

	if t.OnPeer != nil {
		if err = t.OnPeer(peer); err != nil {
			return
		}
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
	rpc := RPC{}
	for {
		err = t.Decoder.Decode(conn, &rpc)
		if err != nil {
			return
		}
		rpc.From = conn.RemoteAddr()
		t.rpcch <- rpc

	}

}
