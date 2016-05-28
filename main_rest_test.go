package main

import (
	"bytes"
	"crypto/tls"
	"net/http"
	"testing"
)

func BenchmarkRESTSetInfo(b *testing.B) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	buf := bytes.NewBufferString(`
		{
			"name":"test",
			"age":1,
			"height":1
		}
	`)
	// run http posts against it
	for i := 0; i < b.N; i++ {
		client.Post("https://localhost:4444", "application/json", buf)
	}
}
