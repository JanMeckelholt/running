package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
	mathRand "math/rand"
	"time"

	log "github.com/sirupsen/logrus"
)

const startOFWeek = time.Monday // Sunday = 0; Monday = 1

func GetStartOfFirstWeek(weeksAgo uint64) uint64 {
	roundDay := time.Now().Round(24 * time.Hour)
	if roundDay.Weekday() != time.Now().Weekday() {
		roundDay = roundDay.Add(-24 * time.Hour)
	}
	var daysSinceMonday int64
	switch roundDay.Weekday() {
	case time.Sunday:
		daysSinceMonday = 6
	default:
		daysSinceMonday = int64(roundDay.Weekday() - startOFWeek)
	}

	return uint64(roundDay.Add(time.Duration(-daysSinceMonday)*24*time.Hour - time.Duration(weeksAgo*7*24*uint64(time.Hour))).Unix())
}

func Encrypt(keyString string, stringToEncrypt string) (encryptedString string) {
	key, err := hex.DecodeString(keyString)
	if err != nil {
		log.Infof("ENCRYPT: could not create hex.Decode key %s", err.Error())
	}
	plaintext := []byte(stringToEncrypt)
	log.Infof("key: %s", string(key))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	encrypter := cipher.NewCBCEncrypter(block, iv)
	ciphertext = PKCS5Padding(ciphertext, aes.BlockSize)

	crypted := make([]byte, len(ciphertext))
	encrypter.CryptBlocks(crypted, ciphertext)

	return base64.StdEncoding.EncodeToString(crypted)
}

// decrypt from base64 to decrypted string
func Decrypt(keyString string, stringToDecrypt string) string {
	log.Infof("DECRYPT: keytring %s", keyString)
	log.Infof("DECRYPT: stringToDecrypt %s", stringToDecrypt)
	key, err := hex.DecodeString(keyString)
	if err != nil {
		log.Infof("DECRYPT: could not create hex.Decode key %s", err.Error())
	}
	ciphertext, err := base64.StdEncoding.DecodeString(stringToDecrypt)
	if err != nil {
		log.Infof("DECRYPT: could not create ciphertext %s", err.Error())
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	log.Infof("iv: %s", iv[:])
	log.Infof("len cipher without iv: %d", len(ciphertext[aes.BlockSize:]))

	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(ciphertext[aes.BlockSize:]))
	decrypter.CryptBlocks(decrypted, ciphertext[aes.BlockSize:])
	return string(PKCS5Trimming(decrypted)[:])
}

// n is the length of random string we want to generate
func RandStr(n int) string {
	var charset = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]byte, n)
	for i := range b {
		// randomly select 1 character from given charset
		b[i] = charset[mathRand.Intn(len(charset))]
	}

	return string(b)
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5Trimming(encrypt []byte) []byte {
	log.Infof("Trimming: %s", encrypt)
	log.Infof("Trimming string: %s", string(encrypt[:]))
	padding := encrypt[len(encrypt)-1]
	log.Infof("padding: %s - %d", padding, int(padding))
	return encrypt[:len(encrypt)-int(padding)]
}
