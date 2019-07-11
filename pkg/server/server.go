package server

import (
	"net/http"
	"os"
	"sort"
)

func RunServer() {

	http.HandleFunc("/", displayEnv)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func displayEnv(rw http.ResponseWriter, _ *http.Request) {

	envStr := "<html><head><title>Service Env</title><body><pre>"
	envs := os.Environ()
	sort.Strings(envs)
	for _, env := range envs {
		envStr += env + "\n"
	}
	envStr += "</body></pre>"

	_, err := rw.Write([]byte(envStr))
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}
