package main

import (
	"log"
	"os"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/packet"
)

func generateKeypair(uid userId, keysize int) keyPair {
	if verbose {
		if uid.Id != "" {
			log.Printf("Generating %d-bit RSA keys with UID: [%s]\n", keysize, uid.Id)
		} else {
			log.Printf("Generating %d-bit RSA keys with empty UID\n", keysize)
		}
	}
	cfg := &packet.Config{RSABits: keysize}
	newPair, err := openpgp.NewEntity(uid.Name, uid.Comment, uid.Email, cfg)
	if err != nil {
		log.Fatalf("Cannot create a new key: — %s\n", err)
	}
	return newPair
}

func saveKeyToFile(entity *openpgp.Entity, filename string) {
	cfg := &packet.Config{}
	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		log.Fatalf("Cannot save to file “%s” — %s\n", filename, err)
	}
	entity.SerializePrivate(f, cfg)
}
