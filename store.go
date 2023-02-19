package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const dafaultRootFolderName = "ggnetwork" //Nome del root path che puo essere cambiato

/*
CAS
In base ad una chiave inserita la trasforma in sha1
Divide la lunghezza in set da 5 byte
Creando una path
Restituisce il path crato e sha1
*/
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
		Filename: hashstr,
	}

}

type PathTransformFunc func(string) PathKey

type PathKey struct {
	PathName string
	Filename string
}

func (p PathKey) FristPathName() string {
	paths := strings.Split(p.PathName, "/")
	if len(paths) == 0 {
		return ""
	}
	return paths[0]
}

func (p PathKey) FullPath() string {
	return fmt.Sprintf("%s/%s", p.PathName, p.Filename)
}

type StoreOpts struct {
	Root              string //Folder name del root che contiene tutti i file e cartelle
	PathTransformFunc PathTransformFunc
}

var DefaultPathTransformFunc = func(key string) PathKey {
	return PathKey{
		PathName: key,
		Filename: key,
	}
}

type Store struct {
	StoreOpts
}

// Crea un nuovo store
func NewStore(opts StoreOpts) *Store {
	if opts.PathTransformFunc == nil {
		opts.PathTransformFunc = DefaultPathTransformFunc
	}
	if len(opts.Root) == 0 {
		opts.Root = dafaultRootFolderName
	}
	return &Store{
		StoreOpts: opts,
	}
}

// Esistenza
func (s *Store) Has(key string) bool {

	pathKey := s.PathTransformFunc(key)
	fullPathWithRoot := fmt.Sprintf("%s/%s", s.Root, pathKey.FullPath())

	_, err := os.Stat(fullPathWithRoot)

	return !errors.Is(err, os.ErrNotExist)

}

// Elimina tutto nella root compresa la root
func (s *Store) Clear() error {
	return os.RemoveAll(s.Root)
}

// Elimina il file
func (s *Store) Delate(key string) error {
	pathKey := s.PathTransformFunc(key)

	defer func() {
		log.Printf("deleted [%s] from disk", pathKey.Filename)
	}()
	firstPathNamewithRoot := fmt.Sprintf("%s/%s", s.Root, pathKey.FristPathName()) //Root + First Path
	return os.RemoveAll(firstPathNamewithRoot)
}

func (s *Store) Write(key string, r io.Reader) error {
	return s.writeStream(key, r)
}

// Legge il file
func (s *Store) Read(key string) (io.Reader, error) {
	f, err := s.readStream(key)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, f)

	return buf, nil
}

/*
Apre il file
*/
func (s *Store) readStream(key string) (io.ReadCloser, error) {
	pathkey := s.PathTransformFunc(key)
	pathKeyWithRoot := fmt.Sprintf("%s/%s", s.Root, pathkey.FullPath())
	return os.Open(pathKeyWithRoot)
}

/*
Crea il file
*/
func (s *Store) writeStream(key string, r io.Reader) error {
	pathKey := s.PathTransformFunc(key)
	pathNameWithRoot := fmt.Sprintf("%s/%s", s.Root, pathKey.PathName)

	if err := os.MkdirAll(pathNameWithRoot, os.ModePerm); err != nil {
		return err
	}

	fullPathWithRoot := fmt.Sprintf("%s/%s", s.Root, pathKey.FullPath())

	f, err := os.Create(fullPathWithRoot)
	if err != nil {
		return err
	}

	n, err := io.Copy(f, r)
	if err != nil {
		return err
	}

	log.Printf("written (%d) bytes to disk: %s", n, fullPathWithRoot)

	return nil

}
