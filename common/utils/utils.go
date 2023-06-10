package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"path/filepath"

	"io"
	mathRand "math/rand"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
)

const Location = "Europe/Berlin"
const startOFWeek = time.Monday // Sunday = 0; Monday = 1

func GetStartOfWeek(week int64) uint64 {
	if week > 0 {
		log.Infof("got invalid week-value %d, defaulting to 0", week)
		week = 0
	}
	loc, _ := time.LoadLocation(Location)
	_, offset := time.Now().In(loc).Zone()
	roundDay := time.Now().UTC().Round(24 * time.Hour).Add(-time.Duration(offset) * time.Second)
	if roundDay.In(loc).Weekday() != time.Now().In(loc).Weekday() {
		roundDay = roundDay.Add(-24 * time.Hour)
	}
	var daysSinceMonday int64
	switch roundDay.In(loc).Weekday() {
	case time.Sunday:
		daysSinceMonday = 6
	default:
		daysSinceMonday = int64(roundDay.In(loc).Weekday() - startOFWeek)
	}

	return uint64(roundDay.Add(time.Duration(-daysSinceMonday)*24*time.Hour + time.Duration(week*7*24*int64(time.Hour))).Unix())
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
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

func DecryptPGP(cipherFile, plaintextFile, skBase64 string) (err error) {
	cipherText, err := os.ReadFile(cipherFile)
	if err != nil {
		return err
	}
	sk, err := base64.StdEncoding.DecodeString(skBase64)
	if err != nil {
		return err
	}
	skBuf := bytes.NewBuffer(cipherText)
	armorBlock, _ := armor.Decode(skBuf)
	privateKeyObj, err := openpgp.ReadKeyRing(bytes.NewReader(sk))
	if err != nil {
		return err
	}
	prompt := func(keys []openpgp.Key, symmetric bool) ([]byte, error) {
		err := keys[0].PrivateKey.Decrypt([]byte(""))
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	md, err := openpgp.ReadMessage(armorBlock.Body, privateKeyObj, prompt, nil)
	if err != nil {
		return err
	}
	os.MkdirAll(filepath.Dir(plaintextFile), 0640)
	outFile, err := os.Create(plaintextFile)
	defer outFile.Close()

	_, err = io.Copy(outFile, md.UnverifiedBody)

	return err
}
