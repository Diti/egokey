package main

import "fmt"

func main() {
	entMaster := createNewKey("John Smith", "Test key, do not use", "john.smith@example.org", 1024)
	saveKeyToFile(entMaster, "test_key")

	fmt.Printf("%X\n", entMaster.PrimaryKey.Fingerprint)
}
