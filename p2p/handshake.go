package p2p

/*
HandShakeFunc ...?
*/
type HandShakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error { return nil }
