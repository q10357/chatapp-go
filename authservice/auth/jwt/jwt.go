package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Generate tokens
func GenerateToken(claimsMap map[string]string) (string, error) {
	//create new hash of type SHA256, pass secret to it
	secret := os.Getenv("TOKEN_SECRET")

	header := map[string]string{
		"alg": "HS256",
		"typ": "JWT",
	}

	headerStr, err := json.Marshal(header)

	if err != nil {
		//Mashal function returned an error
		fmt.Println("Error generating Token")
		//return payloadStr as string and error
		return string(headerStr), err
	}

	h := hmac.New(sha256.New, []byte(secret))

	//Encode header as base64 string
	header64 := base64.StdEncoding.EncodeToString([]byte(headerStr))

	//Marshal (ibject to json string) claims, then base encode the string
	payloadStr, err := json.Marshal(claimsMap)

	if err != nil {
		//Mashal function returned an error
		fmt.Println("Error generating Token")
		//return payloadStr as string and error
		return string(payloadStr), err
	}
	//Continue
	payload64 := base64.StdEncoding.EncodeToString([]byte(payloadStr))

	//Concat strings, seperate by "."
	message := header64 + "." + payload64

	//Unsigned msg ready
	unsignedStr := string(headerStr) + string(payloadStr)

	//write to SHA256 to hash. Use to generate signature
	h.Write([]byte(unsignedStr))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	//This is the token
	tokenStr := message + "." + signature
	return tokenStr, nil
}

func ValidateToken(token string) (bool, error) {
	//Three parts seperated by "."
	splitToken := strings.Split(token, ".")
	fmt.Println(splitToken)

	if len(splitToken) != 3 {
		return false, nil
	}

	//decode header & claimsMap into strings
	//header is the first element of the array
	header, err := base64.StdEncoding.DecodeString(splitToken[0])
	if err != nil {
		return false, err
	}
	//followed by our claimsMap
	claimsMap, err := base64.StdEncoding.DecodeString(splitToken[1])
	if err != nil {
		return false, err
	}

	//create signature
	unsignedStr := string(header) + string(claimsMap)
	h := hmac.New(sha256.New, []byte(os.Getenv("TOKEN_SECRET")))
	h.Write([]byte(unsignedStr))

	//base encode unsignedStr to base64 string
	signature := base64.StdEncoding.EncodeToString([]byte(unsignedStr))
	fmt.Println(signature)

	// if not equal => reject token
	if signature != splitToken[2] {
		return false, nil
	}

	// VALID TOKEN
	return true, nil
}
