package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGetUser(t *testing.T) {
	s := NewServer()
	ts := httptest.NewServer(http.HandlerFunc(s.handleGetUser))
	count := 1000
	for i := 0; i < count; i++ {
		id := i%100 + 1
		url := fmt.Sprintf("%s/?id=%d", ts.URL, id)

		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		user := &User{}
		json.NewDecoder(response.Body).Decode(user)
		if err != nil {
			t.Error(err)
		}

		fmt.Printf("%+v\n", user)
	}
	fmt.Println("times we hit the database", s.dbhit)
}
