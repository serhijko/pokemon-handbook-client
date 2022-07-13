package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://localhost:8080/pokemons/4", nil)
	// req.Header.Add("Auth", "Basic YWRtaW46YWRtaW4=")
	req.SetBasicAuth("admin", "admin")

	resp, err := client.Do(req)

	// resp, err := http.Get("http://localhost:8080/pokemons/4")
	if err != nil {
		fmt.Println("error", err)
	} else {
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("error", err)
		}

		fmt.Println(resp.Status)
		fmt.Println(resp.StatusCode)
		fmt.Println(string(body))
	}
}
