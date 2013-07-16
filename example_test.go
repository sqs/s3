package s3_test

import (
	"fmt"
	"github.com/kr/s3"
	"log"
	"net/http"
	"os"
	"strings"
)

func Example() {
	client := new(s3.Client)
	s3.DefaultKeys = &s3.Keys{
		AccessKey: os.Getenv("S3_ACCESS_KEY"),
		SecretKey: os.Getenv("S3_SECRET_KEY"),
	}
	data := strings.NewReader("hello, world")
	resp, err := client.Put("https://example.s3.amazonaws.com/foo", data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.StatusCode)
}

func ExampleClient_Do() {
	client := new(s3.Client)
	data := strings.NewReader("hello, world")
	req, err := http.NewRequest("PUT", "https://example.s3.amazonaws.com/foo", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("X-Amz-Acl", "public-read")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.StatusCode)
}
