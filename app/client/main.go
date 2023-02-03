package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type addParam struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type addResult struct {
	Code int `json:"code"`
	Data int `json:"data"`
}

func main() {
	url := "http://127.0.0.1:9090/add"
	param := addParam{X: 10, Y: 20}
	jsonStr, _ := json.Marshal(param)
	res, _ := http.Post(url, "application/json", bytes.NewReader(jsonStr))
	defer res.Body.Close()

	respBytes, _ := io.ReadAll(res.Body)
	var respData addResult
	json.Unmarshal(respBytes, &respData)
	fmt.Println(respData.Data) // 30
}