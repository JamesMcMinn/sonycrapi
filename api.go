package sonycrapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const cameraURL = "http://192.168.122.1:8080/sony/"

// Endpoint defined a specfic postifx for the ActionListURL which the API makes a particualr request.
// At present sony defines only 3, and instances of these can be found in the endpoints struct
type endpoint string

var endpoints = struct {
	Camera    endpoint
	System    endpoint
	AVContent endpoint
}{"camera", "system", "avContent"}

type request struct {
	Camera   *Camera       `json:"-"`
	Endpoint endpoint      `json:"-"`
	Method   string        `json:"method"`
	Params   []interface{} `json:"params"`
	Version  string        `json:"version"`
	ID       int           `json:"id"`
}

type response struct {
	Result []json.RawMessage
	Error  []interface{}
	ID     int
}

var versionMap = map[string]map[string]struct{}{
	"1.1": map[string]struct{}{
		"deleteContent": struct{}{},
	},
	"1.2": map[string]struct{}{
		"getContentCount": struct{}{},
	},
	"1.3": map[string]struct{}{
		"getContentList": struct{}{},
		"getEvent":       struct{}{},
	},
}

func (c *Camera) newRequest(ep endpoint, method string, params ...interface{}) *request {
	c.requestID++

	if len(params) == 0 {
		params = []interface{}{}
	}

	version := "1.0"
	for v, methods := range versionMap {
		if _, found := methods[method]; found {
			version = v
		}
	}

	return &request{
		Camera:   c,
		Endpoint: ep,
		Method:   method,
		Params:   params,
		Version:  version,
		ID:       c.requestID,
	}
}

// Do makes the actual REST API call and abtains the response and / or any errors
func (r *request) Do() (cameraResp *response, err error) {
	j, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	dest := r.Camera.ActionListURL + "/" + string(r.Endpoint)

	resp, err := http.Post(dest, "", bytes.NewBuffer(j))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("got %d HTTP status code, Expected 200", resp.StatusCode)
		return
	}

	jresp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	cameraResp = &response{}
	err = json.Unmarshal(jresp, cameraResp)
	if err != nil {
		log.Println("JSON Error:", err)
		log.Println(string(jresp))
		return
	}

	if len(cameraResp.Error) > 1 {
		fmt.Println(r.Params)
		code := cameraResp.Error[0].(float64)
		msg := cameraResp.Error[1].(string)
		return cameraResp, fmt.Errorf("%f: %s", code, msg)
	}

	return
}
