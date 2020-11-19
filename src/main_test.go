package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddHandlerFunc(t *testing.T) {
	tt := []struct {
		name string
		firstVal,secondVal string
		sum int
		status int
		err string
	}{
		{name:"sum of one and two",firstVal: "1",secondVal: "2", status: http.StatusOK,sum :3},
		{name:"sum of one and string",firstVal: "1",secondVal: "aa", status: http.StatusBadRequest,sum :0},
		{name:"sum of string and string",firstVal: "bb",secondVal: "aa", status: http.StatusBadRequest,sum :0},
		{name:"sum of -1 and 2",firstVal: "-1",secondVal: "2", status: http.StatusOK,sum :1},
		{name:"sum of missing val and 2",firstVal: "",secondVal: "2", status: http.StatusBadRequest,sum :0},
	}
	for _,tc := range tt {
		t.Run(tc.name,func(t *testing.T){
			req, err := http.NewRequest("GET", "http://localhost:8000/add/a/1/b/2", nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}
			m := make(map[string]string)
			if tc.firstVal != ""{
				m["first"] = tc.firstVal
			}
			if tc.secondVal != ""{
				m["second"] = tc.secondVal
			}
			req = mux.SetURLVars(req, m)
			w := httptest.NewRecorder()
			addHandlerFunc(w, req)
			resp := w.Result()
			defer resp.Body.Close()
			if resp.StatusCode != tc.status {
				t.Fatalf("Expected %v; Got %v",tc.status, resp.Status)
			}
			respBody := sumResponse{}
			decoder := json.NewDecoder(resp.Body)
			err = decoder.Decode(&respBody)
			if err != nil {
				t.Errorf("could not read expected json response: %v", err)
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					t.Fatalf("could not read response body using ioutil: %v", err)
				} else {
					if strings.EqualFold(string(body), `{"error":"Incorrect input"}`) {
						t.Errorf("error response is %s", string(body))
					}
				}
			}
			if respBody.Sum != tc.sum {
				t.Fatalf("expected sum to be %d; Got %d",tc.sum, respBody.Sum)
			}
		})
	}
}

func TestHeartbeatHandlerFunc(t *testing.T) {
	req,err := http.NewRequest("GET", "http://localhost:8000/", nil)
	if err != nil{
		t.Fatalf("could not create request: %v",err)
	}
	w := httptest.NewRecorder()
	heartbeatHandlerFunc(w, req)
	resp := w.Result()
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK{
		t.Errorf("Expected Status OK; Got %v",resp.Status)
	}
	respBody := heartbeatResponse{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&respBody)
	if err != nil{
		t.Errorf("could not read expected json response: %v", err)
		body,err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("could not read response body using ioutil: %v", err)
		}else{
			t.Errorf("error response is %s",string(body))
		}
	}
	if respBody.Code != http.StatusOK{
		t.Fatalf("expected sum to be 4; Got %v",respBody.Code)
	}
}

//testing http handler routing
func TestRouting(t *testing.T){
	srv := httptest.NewServer(handler())
	defer srv.Close()

	tt := []struct {
		name string
		url string
		status int
	}{
		{name:"sum of one and two",url:"%s/add/a/1/b/2", status: http.StatusOK},
		{name:"health check",url:"%s/", status: http.StatusOK},
		{name:"incorrect sum ",url:"%s/add/a/1", status: http.StatusNotFound},
	}
	for _,tc:=range tt {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := http.Get(fmt.Sprintf(tc.url, srv.URL))
			if err != nil {
				t.Fatalf("could not send GET request: %v", err)
			}
			if resp.StatusCode != tc.status {
				t.Errorf("Expected %v; Got %v",tc.status, resp.Status)
			}
			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("could not read response :%v", err)
			}
			if strings.EqualFold(string(data), `{"error":"Incorrect input"}`) {
				t.Fatalf("error response is %s", string(data))
			}
		})
	}
}