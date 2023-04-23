package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	mathRand "math/rand"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

const startOFWeek = time.Monday // Sunday = 0; Monday = 1

func GetStartOfFirstWeek(weeksAgo uint64) uint64 {
	roundDay := time.Now().Round(24 * time.Hour)
	if roundDay.Weekday() != time.Now().Weekday() {
		roundDay = roundDay.Add(-24 * time.Hour)
	}
	var daysSinceMonday time.Duration
	switch roundDay.Weekday() {
	case time.Sunday:
		daysSinceMonday = time.Duration(6 * 24 * time.Hour)
	default:
		daysSinceMonday = time.Duration(roundDay.Weekday() - startOFWeek)
	}

	return uint64(roundDay.Add(-daysSinceMonday*24*time.Hour - time.Duration(weeksAgo*7*24*uint64(time.Hour))).Unix())
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
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	return base64.StdEncoding.EncodeToString(ciphertext)
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
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	return strings.Trim(fmt.Sprintf("%s", ciphertext), "\b")
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
