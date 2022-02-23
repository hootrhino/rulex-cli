package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ngaut/log"
)

//HTTP Post
func post(data map[string]interface{}, host string, api string) (int, string) {
	p, errs1 := json.Marshal(data)
	if errs1 != nil {
		log.Fatal(errs1)
	}
	r, errs2 := http.Post("http://"+host+":2580/api/v1/"+api,
		"application/json",
		bytes.NewBuffer(p))
	if errs2 != nil {
		log.Fatal(errs2)
	}
	defer r.Body.Close()

	body, errs5 := ioutil.ReadAll(r.Body)
	if errs5 != nil {
		log.Fatal(errs5)
	}
	return r.StatusCode, string(body)
}

// HTTP Get
func get(host string, api string) string {
	// Get list
	r, errs := http.Get(("http://" + host + ":2580/api/v1/" + api))
	if errs != nil {
		log.Error(errs)
		return ""
	}
	defer r.Body.Close()
	body, errs2 := ioutil.ReadAll(r.Body)
	if errs2 != nil {
		log.Error(errs2)
		return ""
	}
	return string(body)
}

// HTTP delete
func delete(host string, api string, param string) string {
	req, err0 := http.NewRequest(
		"DELETE",
		"http://"+host+":2580/api/v1/"+api+"?"+param,
		nil,
	)
	if err0 != nil {
		log.Error(err0)
		return ""
	}
	response, err1 := http.DefaultClient.Do(req)
	if err1 != nil {
		log.Error(err1)
		return ""
	}
	defer response.Body.Close()
	body, errs5 := ioutil.ReadAll(response.Body)
	if errs5 != nil {
		return ""
	}
	return string(body)
}

// HTTP put
func put(host string, api string, data map[string]interface{}) string {
	p, errs1 := json.Marshal(data)
	if errs1 != nil {
		log.Fatal(errs1)
	}
	req, err0 := http.NewRequest(
		"PUT",
		"http://"+host+":2580/api/v1/"+api,
		bytes.NewBuffer(p),
	)
	if err0 != nil {
		log.Error(err0)
		return ""
	}
	response, err1 := http.DefaultClient.Do(req)
	if err1 != nil {
		log.Error(err1)
		return ""
	}
	defer response.Body.Close()
	body, errs5 := ioutil.ReadAll(response.Body)
	if errs5 != nil {
		log.Error(errs5)
		return ""
	}

	return string(body)
}
