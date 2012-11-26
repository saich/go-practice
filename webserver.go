package main

import (
	"fmt"
	"net/http"
)

// net/http handles HTTP requests using any value that implements http.Handler interface

// Package http provides HTTP client and server implementations.

// This example will deal with server implementations only..
/* 
package http

type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
}
*/

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "URL: %s\n", r.URL.Path)
	fmt.Fprintf(w, "Method: %s\n", r.Method)
	fmt.Fprintf(w, "Protocol: %s, Major: %v, Minor: %v\n", r.Proto, r.ProtoMajor, r.ProtoMinor)
	fmt.Fprintf(w, "Host: %s\n", r.Host)
	fmt.Fprintf(w, "Remote Address: %s\n", r.RemoteAddr) // Remote Address that sent this request (IP:port)

	fmt.Fprintf(w, "Headers: %v\n", r.Header)
	fmt.Fprintf(w, "User Agent: %s\n", r.UserAgent()) // As set in the header
	fmt.Fprintf(w, "Referrer: %s\n", r.Referer())     // As set in the header

	// Form data (GET, POST)
	r.ParseForm() // r.Form will be filled by this function
	// r.Form contain GET, POST, PUT paramters all combined
	fmt.Fprintf(w, "URL Parameters: %v\n", r.Form) // Parsed Form data

	// Try this: http://localhost:8080/?test=asdasd&test=qwerty
	// NOTE: The duplicate parameters are appended instead of overwritten..
	fmt.Fprintf(w, "Test Parameters: %v\n", r.Form["test"])

	// TODO: To distinguish GET and POST parameters, @see http://code.google.com/p/go/issues/detail?id=3630

	// Reading cookie
	fmt.Fprintf(w, "All cookies: %v\n", r.Cookies()) // Get all cookies
	testCookie, err := r.Cookie("test")              // Get cookie with the given name. err = ErrNoCookie, if cookie not found
	fmt.Fprintf(w, "Cookie with name test: %v, err = %v, isErrNoCookie=%v\n", testCookie, err, err == http.ErrNoCookie)
}

func main() {
	// Handler Functions:
	// If the path ends with "/", all paths below this (/sai, /prasad etc. for /) are handled by this handler
	// In other case, the handler is for specific resoruce only (e.g: /favicon.ico)
	// In case multiple paths can be satisfied by the path specification, the handler with longest path specification
	// will be used. For example, with handlers for /images and /images/thumbnails/, /images/thumbnails.test.png will be handled
	// by the handler of "/images/thumbnails"

	// The handler paths can include the domain names also, which are useful if multiple domains are configured to
	// the same server.

	// Tells that all paths below "/" (basically all requests) will be sent to this handler
	http.HandleFunc("/", handler)

	// Multiple HandleFuncs can be specified, so basically, HandleFunc() acts as routing mechanism

	// Listen on 8080 on any interface
	http.ListenAndServe(":8080", nil) // This function will block until the program is terminated
}

// Go to http://localhost:8000 to see the output
