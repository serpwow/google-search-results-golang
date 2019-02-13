package serpwow

/*
	This package provides access to Google Search Results via the SerpWow API
 	https://serpwow.com/
*/

import (
	"net/http"
	"net/url"
	"time"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
)

type SerpWowResponse map[string]interface{}

func execute(parameters map[string]string, output string, apiKey string, endpoint string) *http.Response {

	query := url.Values{}
	for k, v := range parameters {
		query.Add(k, v)
	}
	query.Add("api_key", apiKey)
	query.Add("source", "go")
	query.Add("output", output)
	urlFinal := "https://api.serpwow.com/live" + endpoint + "?" + query.Encode()
	var client = &http.Client{
		Timeout: time.Second * 90,
	}
	rsp, err := client.Get(urlFinal)

	if err != nil {
		panic(err.Error())
	}
	return rsp
}

func GetJSON(parameters map[string]string, apiKey string) (SerpWowResponse, error) {
	rsp := execute(parameters, "json", apiKey, "/search")
	return decodeJSON(rsp.Body)
}

func GetHTML(parameters map[string]string, apiKey string) (string, error) {
	rsp := execute(parameters, "json", apiKey, "/search")
	return decodeText(rsp.Body)
}

func GetCSV(parameters map[string]string, apiKey string) (string, error) {
	rsp := execute(parameters, "csv", apiKey, "/search")
	return decodeText(rsp.Body)
}

func GetLocations(parameters map[string]string, apiKey string) (SerpWowResponse, error) {
	rsp := execute(parameters, "json", apiKey, "/locations")
	return decodeJSON(rsp.Body)
}

func GetAccount(apiKey string) (SerpWowResponse, error) {
	rsp := execute(map[string]string{}, "json", apiKey, "/account")
	return decodeJSON(rsp.Body)
}

func decodeJSON(body io.ReadCloser) (SerpWowResponse, error) {

	decoder := json.NewDecoder(body)
	var serpwowResponse SerpWowResponse
	err := decoder.Decode(&serpwowResponse)
	if err != nil {
		return nil, errors.New("fail to decode")
	}

	success := serpwowResponse["request_info"].(map[string]interface{})["success"]

	if success == true {
		return serpwowResponse, nil
	} else {
		return nil, errors.New(serpwowResponse["request_info"].(map[string]interface{})["message"].(string))
	}

}

func decodeText(body io.ReadCloser) (string, error) {
	buffer, err := ioutil.ReadAll(body)
	if err != nil {
		panic(err)
	}
	responseString := string(buffer)
	return responseString, nil
}