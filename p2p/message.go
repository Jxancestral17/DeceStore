package p2p

import "net"

/*
Message
rappresente qualsiasi dato arbitariamento dal tipo che vine trasportato
tra due nodi in una rete
*/
type RPC struct {
	From    net.Addr
	Payload []byte
}
