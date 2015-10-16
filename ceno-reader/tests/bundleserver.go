package main

import (
    "encoding/json"
    "encoding/base64"
    "net/http"
    "net/url"
    "time"
    "fmt"
)

// The port number that the RSS reader expects the bundler server to be listening on.
const BSPort string = "3094"

// Fields that the bundler server is expected to respond with
type BundleResponse struct {
    Url     string `json:"url"`
    Created string `json:"created"`
    Bundle  string `json:"bundle"`
}

/**
 * Run an HTTP server that will act like a bundle server so that we can test
 * that the reader is interacting with it proplerly.
 */
func main() {
    http.HandleFunc("/", reportBundleCreation)
    fmt.Println("Running mock bundle server on http://localhost" + BSPort)
    if err := http.ListenAndServe(BSPort, nil); err != nil {
        panic(err)
    }
}

/**
 * Parse a base64-encoded URL from the query string of a GET request for "/"
 * and report a successful status if the provided URL is a valid URL.
 */
func reportBundleCreation(w http.ResponseWriter, r *http.Request) {
    qs := r.Url.Query()
    b64Url, found := qs["url"]
    if !found {
        w.StatusCode = 400
        w.Write([]byte("No url parameter provided in query string"))
        return
    }
    reqUrlBytes, err := base64.StdEncoding.DecodeString(b64Url)
    if err != nil {
        w.StatusCode = 400
        w.Write([]byte("Could not base64-decode provided URL"))
        return
    }
    requestedUrl = string(reqUrlBytes)
    _, parseErr := url.Parse(requestedUrl)
    if parseErr != nil {
        w.StatusCode = 400
        w.Write([]byte("Invalid URL provided for bundling"))
    } else {
        marshalled, _ := json.Marshal(BundleResponse{
            Url: requestedUrl,
            Created: time.Now().Format(time.UnixDate),
            Bundle: "SOMEBLOBOFDATAREPRESENTINGABUNDLE",
        })
        w.StatusCode = 200
        w.Write(marshalled)
    }
}
