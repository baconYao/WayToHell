package main

import (
	"go-breeders/configuration"
	"os"
	"testing"
)

var testApp application

func TestMain(m *testing.M) {
	testApp = application{
		App: configuration.New(nil),
	}

	os.Exit(m.Run())
}
