package main

import (
	"encoding/base64"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func ImageToText(token string) error {
	image := ReadImg()

	apiURL := fmt.Sprintf("https://aip.baidubce.com/rest/2.0/ocr/v1/accurate_basic?access_token=%s", token)
	param := "image=" + url.QueryEscape(base64.StdEncoding.EncodeToString(image))
	resp, err := http.Post(apiURL, "application/x-www-form-urlencoded", strings.NewReader(param))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read Response Body Error")
		return err
	}
	//var result Result

	json, err := simplejson.NewJson([]byte(body))
	if err != nil {
		fmt.Println("Unmarshal Result Json Error")
		return err
	}
	rows,err := json.Get("words_result").Array()
	if err != nil{
		fmt.Println("Json Get Error: ",err)
	}
	//fmt.Println(rows)
	for k,_ := range rows{
		words := govalidator.ToString(rows[k])
		words = string([]rune(words)[10:])
		words = strings.Trim(words, "]")
		fmt.Println(words)
	}

	return nil
}

func ReadImg() []byte {
	img, err := ioutil.ReadFile("img/111.JPG")
	if err != nil {
		panic(err)
	}
	return img
}