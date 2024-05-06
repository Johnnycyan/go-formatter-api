package main

import (
	"fmt"
	"net/http"
	"strconv"

	//"net/url"
	"github.com/dustin/go-humanize"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		argString := r.URL.Query().Get("num")
		if argString == "" {
			fmt.Fprintf(w, "Error: No number provided")
			return
		}
		arg, err := strconv.Atoi(argString)
		if err != nil {
			fmt.Fprintf(w, "Error: %s", err)
			return
		}
		formatted := getFormatted(arg)
		fmt.Fprintf(w, formatted)
	}
	fmt.Println("Listening on http://localhost:8028")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8028", nil)
}

func getFormatted(num int) string {
	return humanize.Comma(int64(num))
}
