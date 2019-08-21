package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	clientID     = "cmnL3OrbP6nMnCYGS1VgyaEf"           		// API Key
	clientSecret = "zPkV7fHkE8ePPf13GKw7ab4MV0ZH77Q2"  			// Secret Key
)

type Session struct {
	AccessToken   string `json:"access_token"`
	RefreshToken  string `json:"refresh_token"`
	SessionKey    string `json:"session_key"`
	SessionSecret string `json:"session_secret"`
	Scope         string `json:"scope"`
	ExpiresIn     int    `json:"expires_in"`
}

func main() {
	token, err := accessToken(clientID, clientSecret)
	if err != nil {
		fmt.Println("accessToken Error: ", err)
		return
	}
	fmt.Println("Token:", *token)

	//OCR
	err = ImageToText(*token)
	if err != nil {
		fmt.Println("ImageToText Error: ", err)
		return
	}
}

func accessToken(id string, secret string) (token *string, err error) {
	apiURL := fmt.Sprintf("https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id=%s&client_secret=%s", id, secret)
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("HTTP Get Error")
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read Response Body Error")
		return nil, err
	}
	var session Session
	err = json.Unmarshal(body, &session)
	if err != nil {
		fmt.Println("Unmarshal Session Json Error")
		return nil, err
	}
	token = &session.AccessToken
	return token, nil
}








