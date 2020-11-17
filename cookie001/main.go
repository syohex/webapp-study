package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("cookie")
	if err != nil {
		if err != http.ErrNoCookie {
			fmt.Fprintf(w, "err: %v\n", err)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:  "cookie",
			Value: "cookie_value",
			Path:  "/cookie",
		})
		fmt.Fprintf(w, "Set Cookie\n")
		return
	}

	fmt.Fprintf(w, "Found Cookie header and value is %s\n", c.Value)
}

func main() {
	http.HandleFunc("/cookie", handler)
	http.ListenAndServe(":8080", nil)
}
