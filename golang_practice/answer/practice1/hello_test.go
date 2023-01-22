package main

import (
	"fmt"
	"testing"
)

func main() {
	fmt.Println(Hello("aa"))
}

func Hello(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func TestHello(t *testing.T) {
	if Hello("Gopher") != "Hello, Gopher" {
		t.Error("Should be 'Hello, Gopher'")
	}
}
