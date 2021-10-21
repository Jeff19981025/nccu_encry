package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"os/user"
	"path/filepath"
	"strings"
)

func encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func decodeBase64(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}
func Encrypt() {
	key := make([]byte, 32)
	rand.Read(key)
	fmt.Print("rand key: ")
	keystring := encodeBase64(key[:])
	fmt.Println(keystring)

	fmt.Print("type in the file name to encrypt: ")
	var path string
	fmt.Scan(&path)
	usr, _ := user.Current()
	dir := usr.HomeDir
	if path == "~" {
		// In case of "~", which won't be caught by the "else if"
		path = dir
	} else if strings.HasPrefix(path, "~/") {
		// Use strings.HasPrefix so we don't match paths like
		// "/something/~/something/"
		path = filepath.Join(dir, path[2:])
	}
	// filepath:="spaceship-launch.jpg"
	text, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	// generate a new aes cipher using our 32 byte long key
	c, err := aes.NewCipher(key)
	// if there are any errors, handle them
	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile("encrypted.data", gcm.Seal(nonce, nonce, text, nil), 0777)
	// handle this error
	if err != nil {
		// print it out
		fmt.Println(err)
	}
}
