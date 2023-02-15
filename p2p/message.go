package p2p

import "net"

/*
Message
rappresente qualsiasi dato arbitariamento dal tipo che vine trasportato
tra due nodi in una rete
*/
type Message struct {
	From    net.Addr
	PayLoad []byte
}
