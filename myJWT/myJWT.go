package myJWT

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

type str string

type JwtHeader struct {
	Alg string
	Typ string
}
type JWTPayLoad struct {
	Iss string
	Iat string
	Jti uint
}
type JSONMarshaler interface {
	ToJSON() ([]byte, error)
}

func (p JWTPayLoad) ToJSON() ([]byte, error) {
	output, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (h JwtHeader) ToJSON() ([]byte, error) {
	output, err := json.Marshal(h)
	if err != nil {
		return nil, err
	}
	return output, nil
}
func NewMarshaler(n JSONMarshaler) ([]byte, error) {
	return n.ToJSON()
}

func (s str) Base64Url() string {
	stri := strings.Replace(string(s), "=", "", -1)
	stri = strings.Replace(stri, "+", "-", -1)
	stri = strings.Replace(stri, "/", "_", -1)
	return stri
}

func Base64Encode(p, h []byte, s string) string {
	var Header, payload str
	Header = str(base64.StdEncoding.EncodeToString(h))
	payload = str(base64.StdEncoding.EncodeToString(p))
	return fmt.Sprintf("%s.%s.%s", Header.Base64Url(), payload.Base64Url(), s)
}
