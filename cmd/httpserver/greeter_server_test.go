package main_test

import (
	"testing"
	"github.com/ddrouin/specifications"
)

func TestGreeterServer(t *testing.T) {
	driver := go_specs_greet.Driver{BaseURL: "http://localhost:8080"}
	specifications.GreetSpecification(t, nil)
}
