package p2p

import "net"

/*
Peer
è l' interfaccia che rappresente i nodi remoti
*/
type Peer interface {
	//Impelemente tutte le funzioni possibili in net librari
	net.Conn
	Send([]byte) error
	CloseStream()
	//Conn() net.Conn
	//RemoteAddr() net.Addr
	//Close() error
}

/*
Transport

Gestice tutta la comunicazione tra i nodi e la rete (TCP, udp, webscocket)
*/
type Transport interface {
	Addr() string
	Dial(string) error
	ListenAndAccept() error //A prescindere della tipologia di connessione vogliamo sapere solo se ci sono errori
	Consume() <-chan RPC
	Close() error
}
