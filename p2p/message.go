package p2p

const (
	IncomingStream  = 0x2
	IncomingMessage = 0x1
)

/*
Message
rappresente qualsiasi dato arbitariamento dal tipo che vine trasportato
tra due nodi in una rete
*/
type RPC struct {
	From    string
	Payload []byte
	Stream  bool
}
