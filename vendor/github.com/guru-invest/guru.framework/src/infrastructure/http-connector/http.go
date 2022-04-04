package http_connector

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type HttpClient struct {
}

func (c *HttpClient) Get(uri string, headers map[string]string) ([]byte, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", uri, nil)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	res, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.Wrap(err, "Error on executing get request")
	} else {
		if res.StatusCode == http.StatusOK {
			reqBody, _ := ioutil.ReadAll(res.Body)
			return reqBody, nil
		}
	}
	return []byte{}, errors.Wrap(errors.New(res.Status), "Error on executing get request")
}

func (c *HttpClient) Post(uri string, v interface{}, headers map[string]string) ([]byte, error) {

	requestBody, err := json.Marshal(v)
	if err != nil {
		return []byte{}, errors.Wrap(err, "Error on parsing request body")
	}
	client := &http.Client{}
	req, _ := http.NewRequest("POST", uri, bytes.NewBuffer(requestBody))
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	res, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.Wrap(err, "Error on executing request")
	}

	reqBody, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return reqBody, nil
	}

	return reqBody, errors.Wrap(errors.New(strconv.Itoa(res.StatusCode)), res.Status)
}

func formatRequest(r *http.Request) string {
	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}
	// Return the request as a string
	return strings.Join(request, "\n")
}

func (c *HttpClient) Delete(uri string, v interface{}, headers map[string]string) ([]byte, error) {

	requestBody, err := json.Marshal(v)
	if err != nil {
		return []byte{}, errors.Wrap(err, "Error on parsing request body")
	}
	client := &http.Client{}
	req, _ := http.NewRequest("DELETE", uri, bytes.NewBuffer(requestBody))
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	res, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.Wrap(err, "Error on executing request")
	}

	reqBody, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return reqBody, nil
	}

	return reqBody, errors.Wrap(errors.New(strconv.Itoa(res.StatusCode)), res.Status)
}

func (c *HttpClient) Put(uri string, v interface{}, headers map[string]string) ([]byte, error) {

	requestBody, err := json.Marshal(v)
	if err != nil {
		return []byte{}, errors.Wrap(err, "Error on parsing request body")
	}
	client := &http.Client{}
	req, _ := http.NewRequest("PUT", uri, bytes.NewBuffer(requestBody))
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	res, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.Wrap(err, "Error on executing request")
	}

	reqBody, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return reqBody, nil
	}

	return reqBody, errors.Wrap(errors.New(strconv.Itoa(res.StatusCode)), res.Status)
}

func (c *HttpClient) Patch(uri string, v interface{}, headers map[string]string) ([]byte, error) {

	requestBodyBuffered := &bytes.Buffer{}
	if v != nil {
		requestBody, err := json.Marshal(v)
		if err != nil {
			return []byte{}, errors.Wrap(err, "Error on parsing request body")
		}

		requestBodyBuffered = bytes.NewBuffer(requestBody)
	}
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPatch, uri, requestBodyBuffered)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	if err != nil {
		return []byte{}, errors.Wrap(err, "Error on executing new request")
	}

	res, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.Wrap(err, "Error on executing request")
	}

	defer res.Body.Close()

	reqBody, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return reqBody, nil
	}

	return reqBody, errors.Wrap(errors.New(strconv.Itoa(res.StatusCode)), res.Status)
}
