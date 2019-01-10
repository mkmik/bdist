package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"hash"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/kr/binarydist"
)

var (
	dir = flag.String("dir", ".", "output directory")
)

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s OLDBIN NEWBIN\n", os.Args[0])
	flag.PrintDefaults()
}

func init() {
	flag.Usage = usage
}

func diff(oldName, newName, patchName string) error {
	old, err := os.Open(oldName)
	if err != nil {
		return err
	}
	defer old.Close()

	new, err := os.Open(newName)
	if err != nil {
		return err
	}
	defer new.Close()

	patch, err := os.Create(patchName)
	if err != nil {
		return err
	}
	defer patch.Close()

	return binarydist.Diff(old, new, patch)
}

func hashFile(name string) (hash.Hash, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	h := sha256.New()
	io.Copy(h, f)
	return h, nil
}

func run() error {
	if len(flag.Args()) < 2 {
		flag.Usage()
		return nil
	}
	oldName, newName := flag.Arg(0), flag.Arg(1)

	oldHash, err := hashFile(oldName)
	if err != nil {
		return err
	}
	newHash, err := hashFile(newName)
	if err != nil {
		return err
	}

	patchName := filepath.Join(*dir, fmt.Sprintf("%s-to-%s.bpatch",
		hex.EncodeToString(oldHash.Sum(nil)),
		hex.EncodeToString(newHash.Sum(nil)),
	))

	_, err = os.Stat(patchName)
	if os.IsNotExist(err) {
		return diff(oldName, newName, patchName)
	}
	return err
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
