package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func getPath(args []string) (string, error) {
	if len(args) == 2 {
		return args[1], nil
	}
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return dir, nil
}

func main() {
	args := os.Args[1:]
	port := args[0]
	dir, _ := getPath(args)
	fileServer := http.FileServer(http.Dir(dir))
	http.Handle("/", fileServer)

	fmt.Printf("Starting server at %s...", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
