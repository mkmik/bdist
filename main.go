package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/kr/binarydist"
)

func diff(oldName, newName string) error {
	old, err := os.Open(oldName)
	if err != nil {
		return err
	}
	defer old.Close()
	oldHash := sha256.New()

	new, err := os.Open(newName)
	if err != nil {
		return err
	}
	defer new.Close()
	newHash := sha256.New()

	patch, err := ioutil.TempFile(".", ".patch")
	if err != nil {
		return err
	}
	defer patch.Close()
	defer os.Remove(patch.Name())

	if err := binarydist.Diff(io.TeeReader(old, oldHash), io.TeeReader(new, newHash), patch); err != nil {
		return err
	}

	patchName := fmt.Sprintf("%s-to-%s.bpatch",
		hex.EncodeToString(oldHash.Sum(nil)),
		hex.EncodeToString(newHash.Sum(nil)),
	)

	return os.Rename(patch.Name(), patchName)
}

func run() error {
	if len(flag.Args()) < 2 {
		flag.Usage()
	}
	return diff(flag.Arg(0), flag.Arg(1))
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
