package main

import (
	"log"
	"strconv"

	"golang.org/x/crypto/openpgp/packet"
)

func newUserId(args docoptArgs) userId {
	name, _ := args["<name>"].(string)
	email, _ := args["<email>"].(string)
	comment, _ := args["<comment>"].(string)

	return packet.NewUserId(name, comment, email)
}

func parseKeysize(keysizeStr string) int {
	keysize, err := strconv.Atoi(keysizeStr)
	if err != nil {
		log.Fatalf("Invalid keysize number – %s\n", err)
	}
	return keysize
}
