package jwt

import (
	"fmt"
	"testing"
)

func TestNewJWT(t *testing.T) {
	JWT := NewJWT([]byte("test"))
	fmt.Println(JWT)
}

func TestCreateToken(t *testing.T) {
	JWT := NewJWT([]byte("test"))
	fmt.Println(JWT)
	token, err := JWT.CreateToken(CustomClaims{
		Id: int64(10010),
	})
	fmt.Println(token, err)
}

func TestParseToken(t *testing.T) {
	JWT := NewJWT([]byte("test"))
	fmt.Println(JWT)
	token, err := JWT.CreateToken(CustomClaims{
		Id: int64(10010),
	})
	fmt.Println(token, err)
	claim, err := JWT.ParseToken(token)
	fmt.Println(claim, err)
}
