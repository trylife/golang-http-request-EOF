package main

import (
    "fmt"
    "net/http"
    "runtime"
    "time"
)

func main() {
    for i := 0; i < 10000; i++ {
        go func() {
            fmt.Print("new .. ")
            tr := &http.Transport{
                MaxIdleConns:       10,
                IdleConnTimeout:    5 * time.Second,
                DisableCompression: true,
                DisableKeepAlives: false,
            }

            client := &http.Client{Transport: tr}
            resp, err := client.Get("http://localhost:8000")
            fmt.Print(http.StatusOK)

            time.Sleep(2 * time.Second)

            if err != nil {
                fmt.Println(err)
                return
            }

            fmt.Println(http.StatusOK)
            if resp.StatusCode == http.StatusOK {
                time.Sleep(2 * time.Second)
                return
            }

            resp.Body.Close()
            fmt.Println("go num", runtime.NumGoroutine())
            time.Sleep(2 * time.Second)
        }()
    }
}