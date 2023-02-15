package p2p

/*
Peer
è l' interfaccia che rappresente i nodi remoti
*/
type Peer interface {
	Close() error
}

/*
Transport

Gestice tutta la comunicazione tra i nodi e la rete (TCP, udp, webscocket)
*/
type Transport interface {
	ListenAndAccept() error //A prescindere della tipologia di connessione vogliamo sapere solo se ci sono errori
	Consume() <-chan RPC
}
