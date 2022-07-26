package json_compare

import (
	"fmt"
	"github.com/wI2L/jsondiff"
	"testing"
)

var (
	s1 = []byte(`[ {"test1":123} , {"test2":234} ]`)
	t1 = []byte(`[ {"test1":123} , {"test3":234} ]`)
)

func TestJsondiff1(t *testing.T) {
	patch, err := jsondiff.CompareJSONOpts(s1, t1)
	if err != nil {
		// handle error
	}
	for _, op := range patch {
		fmt.Printf("%s\n", op)
	}
}
