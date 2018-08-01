package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base32"
	"errors"
	"fmt"
	"io"
	"net/url"
)

func main() {
	// pad := 4
	// keySize := crypto.SHA256.Size()
	// key := make([]byte, keySize+pad)
	// _, err := rand.Read(key)
	// if err != nil {
	// 	fmt.Println("Read Error: ", err)

	// }
	// secretKey := base32.StdEncoding.EncodeToString(key)[0:56]

	// fmt.Println(secretKey)

	secretKey := "QBEJWOQD56ZUN5FV2R3RROZO23EM427BDVJ4VALYUYWWIQ6CEQBQ===="
	accessToken := "rIbCU43C8KJEMnxXyiG4WOObNnKlFvRhqYWYRL2orv"

	key, _ := base32.StdEncoding.DecodeString(secretKey)

	ciphertext, err := AESEncrypt([]byte(accessToken), key)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("ciphertext: ", ciphertext)
	fmt.Println("ciphertext string: ", string(ciphertext))

	// Gửi về user Encodestring ciphertext
	ciphertextEncode := base32.StdEncoding.EncodeToString(ciphertext)
	fmt.Println("User: ", ciphertextEncode)

	ciphertextDecode, _ := base32.StdEncoding.DecodeString(ciphertextEncode)

	plaintext, err := AESDecrypt(ciphertextDecode, key)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("plaintext: ", plaintext)
	fmt.Println("plaintext string: ", string(plaintext))

	// URL := GetQRURL("RockShip", "congluc19297@gmail.com", secretKey)
	// fmt.Println(URL)

}

func GetQRURL(company, email, secretKey string) string {

	label := fmt.Sprintf("%s:%s", url.QueryEscape(company), email)
	fmt.Println(label)

	u := url.URL{}
	v := url.Values{}
	u.Scheme = "otpauth"
	u.Host = "totp"
	u.Path = label
	v.Add("secret", secretKey)

	u.RawQuery = v.Encode()
	return u.String()
}

// AESEncrypt encrypts data using 256-bit AES-GCM.
func AESEncrypt(plaintext []byte, key []byte) (ciphertext []byte, err error) {
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

// AESDecrypt decrypts data using 256-bit AES-GCM.
func AESDecrypt(ciphertext []byte, key []byte) (plaintext []byte, err error) {
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < gcm.NonceSize() {
		return nil, errors.New("malformed ciphertext")
	}

	return gcm.Open(nil,
		ciphertext[:gcm.NonceSize()],
		ciphertext[gcm.NonceSize():],
		nil,
	)
}
