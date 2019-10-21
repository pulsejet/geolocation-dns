package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

type GeoHandleTestCase struct {
	IP   string `json:"ip"`
	Want string `json:"want"`
}

func TestGeoHandle(t *testing.T) {
	SetupEnvironment("config_test.json")

	// Test cases to run
	var cases []GeoHandleTestCase

	// Open config file
	jsonFile, err := os.Open("intranet_tests.json")
	if err != nil {
		log.Fatal(err)
	}

	// Read config
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &cases)

	for _, test := range cases {
		t.Run(fmt.Sprintf("GeoHandle %s", test.IP), func(t *testing.T) {
			got := GeoHandle(test.IP)
			if got != test.Want {
				t.Errorf("GeoHandle(%s) = %s; want %s", test.IP, got, test.Want)
			}
		})
	}
}
