package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/jdxj/sign/internal/pkg/logger"
)

func Hold() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	s := <-quit
	logger.Infof("receive signal: %d", s)
}

func Salt() string {
	b := make([]byte, 32)
	_, _ = rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func WithSalt(pass, salt string) string {
	pass = fmt.Sprintf("%s:%s", pass, salt)
	sum := sha256.Sum256([]byte(pass))
	return base64.StdEncoding.EncodeToString(sum[:])
}

func Encrypt(key, text string) string {
	sum := sha256.Sum256([]byte(key))
	iv := make([]byte, aes.BlockSize)
	_, _ = rand.Read(iv)
	ciphertext := encrypt(sum[:], iv, []byte(text))
	ciphertext = append(iv, ciphertext...)
	return base64.StdEncoding.EncodeToString(ciphertext)
}

func Decrypt(key, text string) string {
	sum := sha256.Sum256([]byte(key))
	ciphertext, _ := base64.StdEncoding.DecodeString(text)
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	res := encrypt(sum[:], iv, ciphertext)
	return string(res)
}

func encrypt(key, iv, text []byte) []byte {
	block, _ := aes.NewCipher(key)
	ciphertext := make([]byte, len(text))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext, text)
	return ciphertext
}

func PostJson(url string, req, rsp interface{}) error {
	return sendJson(http.MethodPost, url, req, rsp)
}

func PutJson(url string, req, rsp interface{}) error {
	return sendJson(http.MethodPut, url, req, rsp)
}

func DeleteJson(url string, req, rsp interface{}) error {
	return sendJson(http.MethodDelete, url, req, rsp)
}

func sendJson(method, url string, req, rsp interface{}) error {
	body, err := json.Marshal(req)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(body)
	httpReq, err := http.NewRequest(method, url, reader)
	if err != nil {
		return err
	}
	client := &http.Client{}
	httpRsp, err := client.Do(httpReq)
	if err != nil {
		return err
	}
	defer func() {
		_ = httpRsp.Body.Close()
	}()
	if rsp == nil {
		return nil
	}
	decoder := json.NewDecoder(httpRsp.Body)
	return decoder.Decode(rsp)
}
