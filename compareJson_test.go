package json_compare_test

import (
	"fmt"
	"github.com/adrianoribeiro/compareJson/compare"
	"log"
	"testing"
)

var (
	f3 = "data2/a.json"
	f4 = "data2/b.json"
)

func TestCompareJson(t *testing.T) {
	isEqual, err := compare.ExecParallel(f3, f4)
	if err != nil {
		log.Fatal(err)
	}
	if isEqual {
		fmt.Println("The json files are equals")
	} else {
		fmt.Println("The json files aren't equals")
	}
}

