package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

func TestHandleGetUser(t *testing.T) {
	t0 := time.Now()
	s := NewServer()
	ts := httptest.NewServer(http.HandlerFunc(s.handleGetUser))
	count := 1000
	wg := &sync.WaitGroup{}

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(i int) {
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
			wg.Done()
		}(i)
		time.Sleep(time.Millisecond * 1)
	}
	wg.Wait()
	fmt.Println("times we hit the database", s.dbhit)
	fmt.Println("Time spent", time.Since(t0))
}
