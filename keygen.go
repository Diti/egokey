package main

import (
	"log"
	"os"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/packet"
)

func createNewKey(name, comment, email string, keysize int) *openpgp.Entity {
	cfg := &packet.Config{RSABits: keysize}

	newKey, err := openpgp.NewEntity(name, comment, email, cfg)
	if err != nil {
		log.Fatalf("Cannot create a new key: %s", err)
	}
	return newKey
}

func saveKeyToFile(entity *openpgp.Entity, filename string) {
	cfg := &packet.Config{}

	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		log.Fatalf("Cannot save to file “%s”: %s", filename, err)
	}

	entity.SerializePrivate(f, cfg)
}
