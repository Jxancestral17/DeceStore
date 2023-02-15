package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func CASPathTransformFunc(key string) PathKey {
	hash := sha1.Sum([]byte(key)) //[20]byte
	hashstr := hex.EncodeToString(hash[:])

	blocksize := 5
	sliceLen := len(hashstr) / blocksize

	paths := make([]string, sliceLen)

	for i := 0; i < sliceLen; i++ {
		from, to := i*blocksize, (i*blocksize)+blocksize
		paths[i] = hashstr[from:to]
	}

	return PathKey{
		PathName: strings.Join(paths, "/"),
		Orginal:  hashstr,
	}

}

type PathTransformFunc func(string) PathKey

type PathKey struct {
	PathName string
	Orginal  string
}

func (p PathKey) Filename() string {
	return fmt.Sprintf("%s/%s", p.PathName, p.Orginal)
}

type StoreOpts struct {
	PathTransformFunc PathTransformFunc
}

var DefaultPathTransfomrFunc = func(key string) string {
	return key
}

type Store struct {
	StoreOpts
}

func NewStore(opts StoreOpts) *Store {
	return &Store{
		StoreOpts: opts,
	}
}

func (s *Store) writeStream(key string, r io.Reader) error {
	pathKey := s.PathTransformFunc(key)

	if err := os.MkdirAll(pathKey.PathName, os.ModePerm); err != nil {
		return err
	}

	pathAndFilename := pathKey.Filename()

	f, err := os.Create(pathAndFilename)
	if err != nil {
		return err
	}

	n, err := io.Copy(f, r)
	if err != nil {
		return err
	}

	log.Printf("written (%d) bytes to disk: %s", n, pathAndFilename)

	return nil

}
