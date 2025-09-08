package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"
)

var ErrMarshalJson = errors.New("error while turning object into json byte")
var ErrWriteToHmac = errors.New("error while writing to hmac")
var ErrInvalidJWT = errors.New("invalid signature encoding")
var ErrInvalidSignature = errors.New("invalid signature")
var ErrInvalidEncoding = errors.New("invalid payload encoding")

type JWT struct {
	Header  JWTHeader
	Payload JWTPayload
	Hash    []byte
}

type JWTHeader struct {
	Alg     string `json:"alg"`
	Typ     string `json:"typ"`
	encoded string
}

type JWTPayload struct {
	Sub     string    `json:"sub"`
	Iss     string    `json:"iss"`
	Iat     time.Time `json:"iat"`
	Exp     time.Time `json:"exp"`
	Aud     string    `json:"aud"`
	encoded string
}

func NewJWT(userID string, secret string) (*JWT, error) {
	h := JWTHeader{
		Alg: "HS256",
		Typ: "JWT",
	}
	payload := JWTPayload{
		Sub: userID,
		Iss: userID,
		Iat: time.Now(),
		Exp: time.Now().Add(time.Hour * 3 * 24),
		Aud: "user",
	}

	hByt, err := json.Marshal(h)
	if err != nil {
		return nil, ErrMarshalJson
	}

	pByt, err := json.Marshal(payload)
	if err != nil {
		return nil, ErrMarshalJson
	}

	h.encoded = base64UrlEncode(hByt)
	payload.encoded = base64UrlEncode(pByt)

	headersAndPayload := h.encoded + "." + payload.encoded

	hmacHash := hmac.New(sha256.New, []byte(secret))
	n, err := hmacHash.Write([]byte(headersAndPayload))

	if n != len(headersAndPayload) || err != nil {
		return nil, ErrWriteToHmac
	}

	hashed := hmacHash.Sum(nil)

	return &JWT{
		Header:  h,
		Payload: payload,
		Hash:    hashed,
	}, nil
}

func (j *JWT) ToString() (string, error) {
	encodedSignature := base64UrlEncode(j.Hash)

	finalJWT := j.Header.encoded + "." + j.Payload.encoded + "." + encodedSignature
	return finalJWT, nil
}

func (j *JWT) Verify(jwtStr string, secret string) (*JWTPayload, error) {
	parts := strings.Split(jwtStr, ",")
	if len(parts) != 3 {
		return nil, ErrInvalidJWT
	}

	encodedHeader := parts[0]
	encodedPayload := parts[1]
	encodedSignature := parts[2]

	expectedSignature, err := base64UrlDecode(encodedSignature)
	if err != nil {
		return nil, ErrInvalidEncoding
	}

	hmacHash := hmac.New(sha256.New, []byte(secret))
	hmacHash.Write([]byte(encodedHeader + "." + encodedPayload))
	computedSignature := hmacHash.Sum(nil)

	if !hmac.Equal(expectedSignature, computedSignature) {
		return nil, ErrInvalidSignature
	}

	payloadByt, err := base64UrlDecode(encodedPayload)
	if err != nil {
		return nil, ErrInvalidJWT
	}

	var payload JWTPayload
	if err := json.Unmarshal(payloadByt, &payload); err != nil {
		return nil, ErrMarshalJson
	}

	if time.Now().After(payload.Exp) {
		return nil, errors.New("token expired")
	}

	return &payload, nil
}

func base64UrlEncode(value []byte) string {
	return base64.RawURLEncoding.EncodeToString(value)
}
func base64UrlDecode(value string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(value)
}
