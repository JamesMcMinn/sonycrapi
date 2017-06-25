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

type Request struct {
	Camera   *Camera       `json:"-"`
	Endpoint endpoint      `json:"-"`
	Method   string        `json:"method"`
	Params   []interface{} `json:"params"`
	Version  string        `json:"version"`
	ID       int           `json:"id"`
}

type Response struct {
	Result []json.RawMessage
	Error  []interface{}
	ID     int
}

func (c *Camera) newRequest(ep endpoint, method string, params ...interface{}) *Request {
	c.requestID++

	if len(params) == 0 {
		params = []interface{}{}
	}

	return &Request{
		Camera:   c,
		Endpoint: ep,
		Method:   method,
		Params:   params,
		Version:  "1.0",
		ID:       c.requestID,
	}
}

// Do makes the actual REST API call and abtains the response and / or any errors
func (r *Request) Do() (response *Response, err error) {
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

	jresp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	response = &Response{}
	err = json.Unmarshal(jresp, response)
	if err != nil {
		log.Println("JSON Error:", err)
		log.Println(string(jresp))
		return
	}

	if len(response.Error) > 1 {
		fmt.Println(r.Params)
		code := response.Error[0].(float64)
		msg := response.Error[1].(string)
		return response, fmt.Errorf("%f: %s", code, msg)
	}

	return
}
