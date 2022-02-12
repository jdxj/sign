package util

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"time"

	"golang.org/x/term"
)

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

func Encrypt(key, text []byte) (d []byte) {
	sum := sha256.Sum256(key)
	iv := make([]byte, aes.BlockSize)
	_, _ = rand.Read(iv)
	ciphertext := encrypt(sum[:], iv, text)
	d = append(d, iv...)
	d = append(d, ciphertext...)
	return
}

func Decrypt(key, ciphertext []byte) []byte {
	sum := sha256.Sum256(key)
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	return encrypt(sum[:], iv, ciphertext)
}

func encrypt(key, iv, text []byte) []byte {
	block, _ := aes.NewCipher(key)
	ciphertext := make([]byte, len(text))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext, text)
	return ciphertext
}

// Deprecated
func PostJson(url string, req, rsp interface{}) error {
	return SendJson(url, req, rsp, WithMethod(http.MethodPost))
}

// Deprecated
func PutJson(url string, req, rsp interface{}) error {
	return SendJson(url, req, rsp, WithMethod(http.MethodPut))
}

// Deprecated
func DeleteJson(url string, req, rsp interface{}) error {
	return SendJson(url, req, rsp, WithMethod(http.MethodDelete))
}

func newSendJsonOption() *sendJsonOption {
	return &sendJsonOption{
		method: http.MethodPost,
		header: make(map[string]string),
		value:  make(map[string]string),
	}
}

type (
	sendJsonOption struct {
		method string
		header map[string]string
		path   []string
		value  map[string]string
	}
	SendJsonOption func(sjo *sendJsonOption)
)

func WithMethod(m string) SendJsonOption {
	return func(o *sendJsonOption) {
		o.method = m
	}
}

func WithHeader(h map[string]string) SendJsonOption {
	return func(o *sendJsonOption) {
		for k, v := range h {
			o.header[k] = v
		}
	}
}

func WithBearer(t string) SendJsonOption {
	return WithHeader(map[string]string{
		"Authorization": fmt.Sprintf("bearer %s", t),
	})
}

// WithJoin 注意拼接顺序
func WithJoin(p string) SendJsonOption {
	return func(o *sendJsonOption) {
		o.path = append(o.path, p)
	}
}

func WithValue(v map[string]string) SendJsonOption {
	return func(o *sendJsonOption) {
		for k, v := range v {
			o.value[k] = v
		}
	}
}

func SendJson(u string, req, rsp interface{}, options ...SendJsonOption) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 初始化选项
	o := newSendJsonOption()
	for _, opt := range options {
		opt(o)
	}

	// 编码 url
	uu, err := url.Parse(u)
	if err != nil {
		return err
	}
	for _, p := range o.path {
		uu.Path = path.Join(uu.Path, p)
	}
	rawQuery := uu.Query()
	for k, v := range o.value {
		rawQuery.Add(k, v)
	}
	uu.RawQuery = rawQuery.Encode()

	// 编码 body
	body, err := json.Marshal(req)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(body)

	// 创建 request
	httpReq, err := http.NewRequestWithContext(ctx, o.method, uu.String(), reader)
	if err != nil {
		return err
	}

	// 编码 header
	for k, v := range o.header {
		httpReq.Header.Add(k, v)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	// 发送请求
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

const (
	EnterPassword  = "Enter Password"
	RepeatPassword = "Repeat Password"
)

var (
	ErrPasswordInconsistent = errors.New("passwords are inconsistent")
)

func ReadPassword(prompt string) (string, error) {
	fmt.Printf("%s: ", prompt)
	data, err := term.ReadPassword(0)
	fmt.Println()
	return string(data), err
}

func GetPassword() (string, error) {
	return ReadPassword(EnterPassword)
}

func ConfirmPassword() (string, error) {
	pass1, err := ReadPassword(EnterPassword)
	if err != nil {
		return "", err
	}
	pass2, err := ReadPassword(RepeatPassword)
	if err != nil {
		return "", err
	}
	if pass1 != pass2 {
		return "", ErrPasswordInconsistent
	}
	return pass1, nil
}

func GetJson(url string, header map[string]string, rsp interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	c := &http.Client{}
	httpRsp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		_ = httpRsp.Body.Close()
	}()

	decoder := json.NewDecoder(httpRsp.Body)
	return decoder.Decode(rsp)
}

func NewTLSConfig(ca, cert, key string) *tls.Config {
	kp, err := tls.LoadX509KeyPair(cert, key)
	if err != nil {
		log.Printf("LoadX509KeyPair: %s\n", err)
		return nil
	}
	d, err := os.ReadFile(ca)
	if err != nil {
		log.Printf("ReadFile: %s\n", err)
		return nil
	}
	cp := x509.NewCertPool()
	cp.AppendCertsFromPEM(d)
	tc := &tls.Config{
		Certificates: []tls.Certificate{kp},
		RootCAs:      cp,
	}
	return tc
}
