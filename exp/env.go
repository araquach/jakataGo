package main

import "os"
import "strings"
import "fmt"

func main() {

	// To set a key/value pair, use `os.Setenv`. To get a
	// value for a key, use `os.Getenv`. This will return
	// an empty string if the key isn't present in the
	// environment.
	os.Setenv("DB_PASSWORD", "password")
	os.Setenv("DB_USER", "root")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	// Use `os.Environ` to list all key/value pairs in the
	// environment. This returns a slice of strings in the
	// form `KEY=value`. You can `strings.Split` them to
	// get the key and value. Here we print all the keys.
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair)
	}

	fmt.Println(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"))
}