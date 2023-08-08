package main

import "testing"

func TestHelloWorld(t *testing.T) {
    if HelloWorld() != "Hello World!!" {
        t.Fatal("Test fail")
    }
}
