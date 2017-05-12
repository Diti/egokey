package main

import (
	"fmt"
	"log"

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

func dryRun(uid userId, keysize int) keyPair {
	keypair := generateKeypair(uid, keysize)
	if isPrettyKeyId(keypair.PrimaryKey.Fingerprint) {
		fmt.Println("Yes, pretty!")
	} else {
		fmt.Println("No, not pretty...")
	}
	return keypair
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
		kp := dryRun(uid, ks)
		if args["run"] == true {
			saveKeyToFile(kp, args["--output"].(string))
			fmt.Printf("%X\n", kp.PrimaryKey.Fingerprint)
		}
	}
}
