package auth

import (
	"net/http"
	"log"
	"strings"
	"io/ioutil"
	// "example.org/utils"
	"github.com/skratchdot/open-golang/open"
	. "fmt"
	"encoding/json"
)

func GetAuth() {
	respBody := strings.NewReader(`token=a3cc4f01-c7e6-42e2-ab56-b3b792ab7c48&grant_type=authorization_code&token_type_hint=access_token`)
	resp, err := http.NewRequest("POST", "http://sh35.boxes.network:8081/v1/oauth/tokens", respBody)
	if err != nil {
		log.Fatalln(err)
	}
	resp.SetBasicAuth("test_client_1", "test_secret")
	resp.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	req, err := http.DefaultClient.Do(resp)
	if err != nil {
		// handle err
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		Println(err)
	}

	var payload map[string]interface{}
	err = json.Unmarshal([]byte(body), &payload)

	if _, ok := payload["error"]; ok {
		Println("Uh oh! Error with auth", ok)
	}
	Println(payload["error"])
}

func OpenLogin() {
	open.Run("https://google.com/")
}