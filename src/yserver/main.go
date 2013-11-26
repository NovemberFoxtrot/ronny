package yserver

import (
	"net/http"
	"os"
	"yeasy"
)

func New(port string) {
	wd, err := os.Getwd()
	yeasy.CheckError(err)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir(wd+`/public`))))
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(wd+`/templates`))))

	err = http.ListenAndServe(":"+port, nil)
	yeasy.CheckError(err)
}
