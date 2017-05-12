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
  -o, --output=FILE  Output file for the keys [default: egokeyring.pgp]
  -q, --quiet        Print less on standard output.
  -s, --keysize=<N>  Set the RSA key size to N [default: 2048].
  -v, --verbose      Print more on standard output.
  --version          Print this program’s version.`
	VERSION = `EgoKey v0.1`
)

var verbose bool

type keyPair *openpgp.Entity
type userId *packet.UserId
type docoptArgs map[string]interface{}

func run(uid userId, keysize int, saveToFile bool) {
	keypairs := make(chan keyPair, 4)

	go generateKeypair(keypairs, uid, keysize)

	for {
		select {
		case kp := <-keypairs:
			fpr := kp.PrimaryKey.Fingerprint
			if isPrettyKey(fpr) {
				if saveToFile == true {
					filename := fmt.Sprintf("%X.pgp", fpr)
					saveKeyToFile(kp, filename)
				}
				if verbose {
					fmt.Print("\nFound a pretty key ID: ")
				}
				fmt.Printf("%X\n", fpr)
			}
		default:
			time.Sleep(time.Second)
			if verbose {
				fmt.Print(".")
			}
		}
	}
}

func main() {
	args, err := docopt.Parse(USAGE, nil, true, VERSION, false)
	if err != nil {
		log.Fatalf("Argument parse error – %s\n", err)
	}

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
