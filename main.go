package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/packet"

	"github.com/docopt/docopt-go"
)

const (
	USAGE = `EgoKey, an OpenPGP vanity key generator.

Usage:
  egokey
  egokey (run | dry-run) [<name>] [<email> <comment>] [options]
  egokey -h | --help | --version

Options:
  -h, --help         Print this help message.
  -q, --quiet        Print less on standard output.
  -s, --keysize=<N>  Set the RSA key size to N [default: 2048].
  -v, --verbose      Print more on standard output.
  --version          Print this program’s version.`
	VERSION = `EgoKey v0.1`
)

var (
	quiet   bool
	verbose bool
)

type keyPair *openpgp.Entity
type userId *packet.UserId
type docoptArgs map[string]interface{}

func run(uid userId, keysize int, saveToFile bool) {
	keypairs := make(chan keyPair, 4)

	go generateKeypair(keypairs, uid, keysize)

	ticker := time.NewTicker(time.Millisecond * 500)
	for range ticker.C {
		select {
		case kp := <-keypairs:
			if verbose {
				fmt.Print(".")
			}
			fpr := fmt.Sprintf("%X", kp.PrimaryKey.Fingerprint)
			if isPrettyKey(fpr) {
				if saveToFile == true {
					saveKeyToFile(kp, fpr+".pgp")
				}
				if verbose {
					fmt.Print("\nFound a pretty key ID: ")
				}
				if !quiet {
					fmt.Println(fpr)
				}
			}
		}
	}
}

func main() {
	args, err := docopt.Parse(USAGE, nil, true, VERSION, false)
	if err != nil {
		log.Fatalf("Argument parse error – %s\n", err)
	}

	quiet = args["--quiet"].(bool)
	verbose = args["--verbose"].(bool)

	switch {
	case args["run"] == args["dry-run"]:
		// “No argument” case, since these args are mutually exclusive
		fmt.Println(USAGE)
	case args["run"] == true, args["dry-run"] == true:
		uid := newUserId(args)
		ks := parseKeysize(args["--keysize"].(string))
		save := args["run"] == true

		run(uid, ks, save)
	}
}
