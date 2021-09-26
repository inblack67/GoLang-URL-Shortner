package shortner

import (
	"fmt"
	"net/http"
)

func GetURLs(rw http.ResponseWriter, r *http.Request) {
	redirects := dummies
	fmt.Println(redirects)
	err := redirects.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Something went wrong", http.StatusInternalServerError)
		return
	}
}
