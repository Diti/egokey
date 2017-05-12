package main

import (
	"fmt"
	"log"
	"strconv"

	"golang.org/x/crypto/openpgp"

	"github.com/docopt/docopt-go"
)

const (
	USAGE = `EgoKey, an OpenPGP vanity key generator.

Usage:
  egokey
  egokey (run | dry-run) <name> [<email> <comment>] [options]
  egokey -h | --help | --version

Options:
  -h --help         Print this help message.
  -o --output=FILE  Output file for the keys [default: egokeyring.pgp]
  -q --quiet        Print less on standard output.
  -s --keysize=<N>  Set the RSA key size to N [default: 2048].
  -v --verbose		Print more on standard output.
  --version         Print this program’s version.`
	VERSION = `EgoKey v0.1`
)

func generateKeypairWithArgs(args map[string]interface{}) *openpgp.Entity {
	var ok bool
	var name, email, comment string

	// Parse arguments
	keysize, err := strconv.Atoi(args["--keysize"].(string))
	if err != nil {
		log.Fatalf("Invalid keysize number -- %s", err)
	}
	if name, ok = args["<name>"].(string); !ok {
		name = ""
	}
	if email, ok = args["<email>"].(string); !ok {
		email = ""
	}
	if comment, ok = args["<comment>"].(string); !ok {
		comment = ""
	}

	return createNewKey(name, comment, email, keysize)
}

func main() {
	args, err := docopt.Parse(USAGE, nil, true, VERSION, false)
	if err != nil {
		log.Fatal(err)
	}

	switch {
	case args["run"] == args["dry-run"]:
		// “No argument” case, since these args are mutually exclusive
		fmt.Println(USAGE)
	case args["dry-run"] == true:
		keypair := generateKeypairWithArgs(args)
		fmt.Printf("\n%+v", keypair) // FIXME
	case args["run"] == true:
		keypair := generateKeypairWithArgs(args)
		saveKeyToFile(keypair, args["--output"].(string))
		fmt.Printf("%X\n", keypair.PrimaryKey.Fingerprint)
	}
}
