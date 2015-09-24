package main

import (
	"C"
	"encoding/json"
	"io/ioutil"
	"net/http"
)
import "fmt"

type Command struct {
	URIs []string `json:"uris"`
}

type Result map[string]*Response
type Response struct {
	Body   string `json:"body"`
	Status int    `json:"status"`
	Err    string `json:"error"`
	Uri    string `json:"uri"`
	Title  string `json:"title"`
}

type Title struct {
	TheTitle string `json:"title"`
}

//export add
// func add(a, b int) int {
// 	return a + b
// }

//export fibonacci
// func fibonacci(n int) int {
// 	a := 1
// 	b := 2

// 	if n == 0 {
// 		return 0
// 	}

// 	var i int
// 	for i = 3; i < n; i++ {
// 		c := a + b
// 		a = b
// 		b = c
// 		fmt.Println(a)
// 		fmt.Println(b)
// 	}
// 	return b
// }

func makeRequest(uri string) *Response {
	resp, err := http.Get(uri)
	if err != nil {
		return &Response{Status: resp.StatusCode, Err: err.Error()}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &Response{Status: resp.StatusCode, Err: err.Error()}
	}
	fmt.Println(string(body))
	t := new(Title)
	err = json.Unmarshal(body, &t)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Title", t.TheTitle)

	return &Response{Uri: uri, Body: string(body), Status: resp.StatusCode}
}

//export scatter_request
func scatter_request(data *C.char) string {
	cmd := Command{}
	err := json.Unmarshal([]byte(C.GoString(data)), &cmd)
	if err != nil {
		return err.Error()
	}

	c := make(chan *Response)
	for _, uri := range cmd.URIs {
		go func(_uri string) {
			c <- makeRequest(_uri)
		}(uri)
	}

	result := Result{}
	for i := 0; i < len(cmd.URIs); i++ {
		resp := <-c
		result[resp.Uri] = resp
	}

	close(c)

	b, err := json.Marshal(result)
	if err != nil {
		return "{}"
	}

	return string(b)
}

func main() {}
