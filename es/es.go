package es

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println(GetData("1"))
	fmt.Println(GetData("2"))
	res, err := CreateData("3", "{\"name\": \"波波\"}")
	fmt.Println(err)
	fmt.Println(res)
	fmt.Println(GetData("3"))
}

func GetData(id string) string {
	resp, err := http.Get("http://127.0.0.1:9200/customer/_doc/" + id)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}

func CreateData(id string, data string) (string, error) {
	targetUrl := "http://127.0.0.1:9200/customer/_doc/" + id

	payload := strings.NewReader(data)

	req, _ := http.NewRequest("PUT", targetUrl, payload)

	req.Header.Add("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	return string(body), nil
}
