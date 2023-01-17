package authorization

import (
	"bufio"
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"os"
	"strings"
)

var secret = []byte("secret_key")

func HasPermision(rawJWT string, hndlr string) bool {

	jwtClaims, err := JwtReader(rawJWT)
	println("autorizer 1, rawJWT: " + rawJWT)
	role := jwtClaims["role"].(string)
	//confirmed := jwtClaims["CCode"].(string)
	//println("autorizer role and ccode: " + role + confirmed)
	println("autorizer role and ccode: " + role)
	f, err := os.Open("permissions/permissions.txt")
	if err != nil {
		log.Fatal(err)
		println("baca eror za otvaranje fajla ")
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ",")
		//if split[0] == role && split[1] == confirmed && split[2] == hndlr {
		if split[0] == role && split[2] == hndlr {
			println("ovo kaze da valja sve")
			return true
		}
	}
	//println("ovo kaze da se ne poklapa sa zahtevima u fajlu: " + role + confirmed + hndlr)
	println("ovo kaze da se ne poklapa sa zahtevima u fajlu: " + role + hndlr)
	return false
}

func JwtReader(rawAuth string) (jwt.MapClaims, error) {

	splitToken := strings.Split(rawAuth, "Bearer ")
	rawJWT := splitToken[1]

	token, err := jwt.Parse(rawJWT, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
