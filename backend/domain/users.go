package domain

import "crypto/sha256"

type Users struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password Password `json:"-"`
}
type Password struct {
	Hashed []byte
	String string
}

func (p *Password) Hash() {
	h := sha256.Sum256([]byte(p.String))
	p.Hashed = h[:]
}

func (p *Password) Match(value string) bool {
	h := sha256.Sum256([]byte(value))
	return string(h[:]) == string(p.Hashed)
}
