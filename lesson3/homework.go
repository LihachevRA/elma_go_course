package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type URLPathType string
type HTTPErrorLocale string

const (
	MainURL  URLPathType = "/"
	LoginURL URLPathType = "/login"
)

const (
	ForbiddenErrorMessage        HTTPErrorLocale = "403 Forbidden"
	NotFoundErrorMessage         HTTPErrorLocale = "404 Not Found"
	MethodNotAllowedErrorMessage HTTPErrorLocale = "405 Method Not Allowed"
	InternalServerErrorMessage   HTTPErrorLocale = "500 Internal Server Error"
)

type LoginForm struct {
	Username string
	Password string
}

func postHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)

	if req.URL.Path == string(MainURL) {
		switch req.Method {
		case "POST":
			if err := req.ParseForm(); err != nil {
				fmt.Printf("ParseForm() err: %v\n", err)
				http.Error(w, string(InternalServerErrorMessage), http.StatusInternalServerError)
				return
			}
			fmt.Fprintf(w, req.Form.Encode())

			return
		default:
			http.Error(w, string(MethodNotAllowedErrorMessage), http.StatusMethodNotAllowed)
			return
		}
	}

	// just for test
	if req.URL.Path == string(LoginURL) {
		switch req.Method {
		case "POST":
			if err := req.ParseForm(); err != nil {
				fmt.Printf("ParseForm() err: %v\n", err)
				http.Error(w, string(InternalServerErrorMessage), http.StatusInternalServerError)
				return
			}

			decoder := json.NewDecoder(req.Body)

			var loginData LoginForm

			if err := decoder.Decode(&loginData); err != nil {
				fmt.Printf("Decode() err: %v\n", err)
				http.Error(w, string(ForbiddenErrorMessage), http.StatusForbidden)
				return
			}

			jsonLoginData, encodingError := json.Marshal(loginData)

			if encodingError != nil {
				fmt.Printf("Json Encoding err: %v\n", encodingError)
				http.Error(w, string(InternalServerErrorMessage), http.StatusInternalServerError)
				return
			}

			fmt.Fprintf(w, string(jsonLoginData))
			return
		default:
			http.Error(w, string(MethodNotAllowedErrorMessage), http.StatusMethodNotAllowed)
			return
		}
	}

	http.Error(w, string(NotFoundErrorMessage), http.StatusNotFound)
	return
}

func main() {
	http.HandleFunc("/", postHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
