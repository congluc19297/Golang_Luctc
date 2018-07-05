package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/alexedwards/scs"
)

// Initialize a new encrypted-cookie based session manager and store it in a global
// variable. In a real application, you might inject the session manager as a
// dependency to your handlers instead. The parameter to the NewCookieManager()
// function is a 32 character long random key, which is used to encrypt and
// authenticate the session cookies.
var sessionManager = scs.NewCookieManager("u46IpCV9y5Vlur8YvODJEhgOY8m9JVE4")

func main() {
	// Set up your HTTP handlers in the normal way.
	mux := http.NewServeMux()
	mux.HandleFunc("/put", putHandler)
	mux.HandleFunc("/get", getHandler)
	mux.HandleFunc("/delete", removeHandler)
	// Wrap your handlers with the session manager middleware.
	http.ListenAndServe(":4000", sessionManager.Use(mux))
}

func putHandler(w http.ResponseWriter, r *http.Request) {
	// Load the session data for the current request. Any errors are deferred
	// until you actually use the session data.
	session := sessionManager.Load(r)

	// Use the PutString() method to add a new key and associated string value
	// to the session data. Methods for many other common data types are also
	// provided. The session data is automatically saved.
	err := session.PutString(w, "message", "Hello world!")
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	// Load the session data for the current request.
	session := sessionManager.Load(r)

	// Use the GetString() helper to retrieve the string value for the "message"
	// key from the session data. The zero value for a string is returned if the
	// key does not exist.
	message, err := session.GetString("message")
	fmt.Println("message không tồn tại: ", message)
	fmt.Println("err không tồn tại messege: ", message)
	if message == "" {
		io.WriteString(w, "Không tôn tại session này")
	}
	if err != nil {
		fmt.Println("GetString err: ", err)
		http.Error(w, err.Error(), 500)
	}

	io.WriteString(w, message)
}

func removeHandler(w http.ResponseWriter, r *http.Request) {
	session := sessionManager.Load(r)

	err := session.Remove(w, "message")
	if err != nil {
		// io.WriteString(w, )
		fmt.Println("removeHander err: ", err)

		http.Error(w, err.Error(), 500)
	}

	io.WriteString(w, "remove thành công")
}
