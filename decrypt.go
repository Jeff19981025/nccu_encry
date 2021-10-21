package main

import (
    "crypto/aes"
    "crypto/cipher"
    "fmt"
    "io/ioutil"
)

func Decrypt() {

    // key := []byte("passphrasewhichneedstobe32bytes!")
    fmt.Print("type in the key: ")
    var key []byte
    var keystr string
    fmt.Scanln(&keystr)
    key = decodeBase64(keystr)
    ciphertext, err := ioutil.ReadFile("encrypted.data")
    // if our program was unable to read the file
    // print out the reason why it can't
    if err != nil {
        fmt.Println(err)
    }

    c, err := aes.NewCipher(key)
    if err != nil {
        fmt.Println(err)
    }

    gcm, err := cipher.NewGCM(c)
    if err != nil {
        fmt.Println(err)
    }

    nonceSize := gcm.NonceSize()
    fmt.Printf("nonceSize: %d\n",nonceSize)
    if len(ciphertext) < nonceSize {
        fmt.Println(err)
    }

    nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
    plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        fmt.Println(err)
    }
    // fmt.Println(string(plaintext))
    _ = ioutil.WriteFile("decrypted.jpg",plaintext,0777)
}
