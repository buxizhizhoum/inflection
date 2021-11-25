package test

import (
	"fmt"
	"github.com/buxizhizhoum/inflection"
	"testing"
)

func TestUnderscore(t *testing.T) {
	testCases := make(map[string]string)
	testCases["aA"] = "a_a"
	testCases["aaA"] = "aa_a"
	testCases["aAA"] = "a_aa"
	testCases["my_case1"] = "my_case1"
	testCases["ILoveGolangAndJSONSoMuch"] = "i_love_golang_and_json_so_much"

	for k, v := range testCases {
		res := inflection.Underscore(k)
		if res != v {
			t.Errorf("TestUnderscore Failed, input: %s,output: %s", k, v)
		}
		fmt.Println(res)
	}

}

func TestCamelizeExpectedFirstUpper(t *testing.T) {
	testCases := make(map[string]string)
	testCases["a_a"] = "AA"
	testCases["aa_a"] = "AaA"
	testCases["aaa"] = "Aaa"

	for k, v := range testCases {
		res := inflection.Camelize(k, true)
		if res != v {
			t.Errorf("TestCamelize Failed, input: %s,output: %s", k, v)
		}
	}
}

func TestCamelizeExpectedFirstLower(t *testing.T) {
	testCases := make(map[string]string)
	testCases["a_a"] = "aA"
	testCases["aa_a"] = "aaA"
	testCases["aaa"] = "aaa"

	for k, v := range testCases {
		res := inflection.Camelize(k, false)
		if res != v {
			t.Errorf("TestCamelize Failed, input: %s,output: %s", k, v)
		}
	}
}
