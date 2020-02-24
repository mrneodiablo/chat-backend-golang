package utilities

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"github.com/astaxie/beego"
	"io"
	"strconv"
	"time"
)

func Decrypt(encrypted string) ([]byte, error) {
	key := []byte(beego.AppConfig.String("secretkey"))
	ciphertext, err := base64.RawURLEncoding.DecodeString(encrypted)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(ciphertext, ciphertext)
	return ciphertext, nil
}

func Encrypt(data []byte) ([]byte, error) {
	key := []byte(beego.AppConfig.String("secretkey"))
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte(""), err
	}
	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return []byte(""), err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], data)
	return []byte(base64.RawURLEncoding.EncodeToString(ciphertext)), nil
}

func GenerateSession(userId int64) string {
	// md5sum(secretkey + userId + yymmddhhmm)
	key := (beego.AppConfig.String("secretkey"))
	utc := time.Now().UTC().Format("200601021504")
	data := string(key) + strconv.FormatInt(userId, 10) + string(utc)
	hasher := md5.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}
