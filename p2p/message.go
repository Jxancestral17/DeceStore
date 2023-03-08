package p2p

/*
Message
rappresente qualsiasi dato arbitariamento dal tipo che vine trasportato
tra due nodi in una rete
*/
type RPC struct {
	From    string
	Payload []byte
}
