package main

import (
	"os"
	"patterns_go/models"
	"testing"
)

var testApp application

func TestMain(m *testing.M) {

	testApp = application{
		Models: *models.New(nil),
	}
	os.Exit(m.Run())
}
