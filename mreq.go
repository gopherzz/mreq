package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	url := flag.String("u", "localhost:80", "Url to request")
	method := flag.String("m", "GET", "Request Method")
	content := flag.String("c", "", "Request Content-Type")
	body := flag.String("b", "text/plain", "Request Body")
	scheme := flag.String("s", "http", "Scheme")
	flag.Parse()

	if len(os.Args) == 1 {
		flag.PrintDefaults()
		return
	}

	if !strings.HasPrefix(*url, "http://") {
		*url = *scheme + "://" + *url
	}

	req, err := http.NewRequest(*method, *url, strings.NewReader(*body))
	if err != nil {
		panic(err.Error())
	}
	req.Header.Set("Content-Type", *content)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)

	buf := make([]byte, 1024)
	for {
		n, err := resp.Body.Read(buf)
		if err != nil {
			break
		}
		println(buf)
		os.Stdout.Write(buf[:n])
	}
}
