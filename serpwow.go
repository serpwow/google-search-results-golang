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
	"strconv"
	"fmt"
	"bytes"
)

type SerpWowResponse map[string]interface{}

func httpGet(parameters map[string]interface {}, output string, apiKey string, endpoint string) *http.Response {

	query := url.Values{}
	for k, v := range parameters {
		query.Add(fmt.Sprintf("%v", k), fmt.Sprintf("%v", v))
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

func httpDelete(apiKey string, endpoint string) *http.Response {

	query := url.Values{}
	query.Add("api_key", apiKey)
	query.Add("source", "go")
	urlFinal := "https://api.serpwow.com/live" + endpoint + "?" + query.Encode()
	var client = &http.Client{
		Timeout: time.Second * 90,
	}
	req, err := http.NewRequest("DELETE", urlFinal, nil)
	if err != nil {
		panic(err.Error())
	}
	
	rsp, err := client.Do(req)
	if err != nil {
		panic(err.Error())
	}
		
	return rsp
}

func httpPost(apiKey string, endpoint string, data map[string]interface {}) *http.Response {

	query := url.Values{}
	query.Add("api_key", apiKey)
	query.Add("source", "go")
	urlFinal := "https://api.serpwow.com/live" + endpoint + "?" + query.Encode()
	var client = &http.Client{
		Timeout: time.Second * 90,
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(data)
	req, err := http.NewRequest("POST", urlFinal, buf)
	if err != nil {
		panic(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
		
	rsp, err := client.Do(req)
	if err != nil {
		panic(err.Error())
	}

	if err != nil {
		panic(err.Error())
	}
		
	return rsp
}

func httpPut(apiKey string, endpoint string, data map[string]interface {}) *http.Response {

	query := url.Values{}
	query.Add("api_key", apiKey)
	query.Add("source", "go")
	urlFinal := "https://api.serpwow.com/live" + endpoint + "?" + query.Encode()
	var client = &http.Client{
		Timeout: time.Second * 90,
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(data)
	req, err := http.NewRequest("PUT", urlFinal, buf)
	if err != nil {
		panic(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
		
	rsp, err := client.Do(req)
	if err != nil {
		panic(err.Error())
	}

	if err != nil {
		panic(err.Error())
	}
		
	return rsp
}

func GetJSON(parameters map[string]interface{}, apiKey string) (SerpWowResponse, error) {
	rsp := httpGet(parameters, "json", apiKey, "/search")
	return decodeJSON(rsp.Body)
}

func GetHTML(parameters map[string]interface{}, apiKey string) (string, error) {
	rsp := httpGet(parameters, "json", apiKey, "/search")
	return decodeText(rsp.Body)
}

func GetCSV(parameters map[string]interface{}, apiKey string) (string, error) {
	rsp := httpGet(parameters, "csv", apiKey, "/search")
	return decodeText(rsp.Body)
}

func GetLocations(parameters map[string]interface{}, apiKey string) (SerpWowResponse, error) {
	rsp := httpGet(parameters, "json", apiKey, "/locations")
	return decodeJSON(rsp.Body)
}

func GetAccount(apiKey string) (SerpWowResponse, error) {
	rsp := httpGet(map[string]interface{}{}, "json", apiKey, "/account")
	return decodeJSON(rsp.Body)
}


func CreateBatch(parameters map[string]interface {}, apiKey string) (SerpWowResponse, error)  {
	rsp := httpPost(apiKey, "/batches", parameters)
	return decodeJSON(rsp.Body)
}

func UpdateBatch(batchId string, parameters map[string]interface {}, apiKey string) (SerpWowResponse, error)  {
	rsp := httpPut(apiKey, "/batches/" + batchId, parameters)
	return decodeJSON(rsp.Body)
}

func StartBatch(batchId string, apiKey string) (SerpWowResponse, error) {
	rsp := httpGet(map[string]interface{}{}, "json", apiKey, "/batches/" + batchId + "/start")
	return decodeJSON(rsp.Body)
}

func StopBatch(batchId string, apiKey string) (SerpWowResponse, error) {
	rsp := httpGet(map[string]interface{}{}, "json", apiKey, "/batches/" + batchId + "/stop")
	return decodeJSON(rsp.Body)
}

func GetBatch(batchId string, apiKey string) (SerpWowResponse, error) {
	rsp := httpGet(map[string]interface{}{}, "json", apiKey, "/batches/" + batchId)
	return decodeJSON(rsp.Body)
}

func ListBatches(apiKey string) (SerpWowResponse, error) {
	rsp := httpGet(map[string]interface{}{}, "json", apiKey, "/batches")
	return decodeJSON(rsp.Body)
}

func ListBatchSearches(batchId string, page int, apiKey string) (SerpWowResponse, error) {
	rsp := httpGet(map[string]interface{}{}, "json", apiKey, "/batches/" + batchId + "/searches/" + strconv.Itoa(page))
	return decodeJSON(rsp.Body)
}

func ListAllBatchSearchesAsJSON(batchId string, apiKey string) (SerpWowResponse, error) {
	rsp := httpGet(map[string]interface{}{}, "json", apiKey, "/batches/" + batchId + "/searches/json")
	return decodeJSON(rsp.Body)
}

func ListAllBatchSearchesAsCSV(batchId string, apiKey string) (SerpWowResponse, error) {
	rsp := httpGet(map[string]interface{}{}, "json", apiKey, "/batches/" + batchId + "/searches/csv")
	return decodeJSON(rsp.Body)
}

func ListBatchResultSets(batchId string, apiKey string) (SerpWowResponse, error) {
	rsp := httpGet(map[string]interface{}{}, "json", apiKey, "/batches/" + batchId + "/results")
	return decodeJSON(rsp.Body)
}

func GetBatchResultSet(batchId string, resultSetId int, apiKey string) (SerpWowResponse, error) {
	rsp := httpGet(map[string]interface{}{}, "json", apiKey, "/batches/" + batchId + "/results/" + strconv.Itoa(resultSetId))
	return decodeJSON(rsp.Body)
}

func GetBatchResultSetAsCSV(batchId string, resultSetId int, apiKey string) (SerpWowResponse, error) {
	rsp := httpGet(map[string]interface{}{}, "json", apiKey, "/batches/" + batchId + "/results/" + strconv.Itoa(resultSetId) + "/csv")
	return decodeJSON(rsp.Body)
}

func DeleteBatch(batchId string, apiKey string) (SerpWowResponse, error) {
	rsp := httpDelete(apiKey, "/batches/" + batchId)
	return decodeJSON(rsp.Body)
}

func DeleteBatchSearch(batchId string, searchId string, apiKey string) (SerpWowResponse, error) {
	rsp := httpDelete(apiKey, "/batches/" + batchId + "/" + searchId)
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