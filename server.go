package main

import "github.com/Jxancestral17/DeceStore/p2p"

type FileServerOpts struct {
	StorageRoot       string
	PathTransformFunc PathTransformFunc
	Transport         p2p.Transport
	TCPTransport      p2p.TCPTransportOpts
}

type FileServer struct {
	FileServerOpts
	store *Store
}

func NewFileServer(opts FileServerOpts) *FileServer {
	storeOpts := StoreOpts{
		Root:              opts.StorageRoot,
		PathTransformFunc: opts.PathTransformFunc,
	}
	return &FileServer{
		FileServerOpts: opts,
		store:          NewStore(storeOpts),
	}
}

func (s *FileServerOpts) Start() error {
	if err := s.Transport.ListenAndAccept(); err != nil {
		return err
	}
	return nil
}
