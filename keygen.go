package main

import (
	"log"
	"os"
	"time"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/packet"
)

func generateKeypair(out chan<- keyPair, uid userId, keysize int) {
	if verbose {
		if uid.Id != "" {
			log.Printf("Generating %d-bit RSA keys with UID: [%s]\n", keysize, uid.Id)
		} else {
			log.Printf("Generating %d-bit RSA keys with empty UID\n", keysize)
		}
	}
	cfg := &packet.Config{RSABits: keysize}

	// Run as many keygens as the remaining elements in channel
	for cap(out)-len(out) >= 0 {
		time.Sleep(time.Millisecond * 500)
		newPair, err := openpgp.NewEntity(uid.Name, uid.Comment, uid.Email, cfg)
		if err != nil {
			log.Fatalf("Cannot create a new key: — %s\n", err)
		}
		out <- newPair
	}
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
