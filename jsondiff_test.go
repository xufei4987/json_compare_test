package json_compare

import (
	"encoding/json"
	"fmt"
	jsonpatch "github.com/evanphx/json-patch/v5"
	"github.com/wI2L/jsondiff"
	jp "github.com/zucong/jsonpath"
	"reflect"
	"sort"
	"testing"
)

var (
	source = []byte(`{
  "age": 33,
  "value": "Juca",
  "id": "abcd001",
  "address": {
    "street": "11 Rue du Grenier Saint-Lazare",
    "postalCode": "75003-111",
    "city": "Paris",
    "countryCode": "FRA",
    "country": "France",
    "text": "11 Rue du Grenier Saint-Lazare, 75003 Paris, France",
    "text1": "11 Rue du Grenier Saint-Lazare, 75003 Paris, France"
  },
  "friends": [
    {
      "name": "buuki1",
      "age": 21
    },
    {
      "name": "cuuki1",
      "age": 30
    }
  ]
}`)
	target = []byte(`{
  "friends": [
    {
      "name": "buuki2",
      "age": 21
    },
    {
      "name": "cuuki2",
      "age": 30
    }
  ],
  "age": 34,
  "value": "Juca",
  "id": "abcd001",
  "address": {
    "street": "11 Rue du Grenier Saint-Lazare",
    "postalCode": "75003-111",
    "city": "Paris",
    "countryCode": "FRA",
    "country": "France",
    "text": "11 Rue du Grenier Saint-Lazare, 75003 Paris, France",
    "text2": "11 Rue du Grenier Saint-Lazare, 75003 Paris, France"
  }
}`)
)

type Patch struct {
	Op   string
	Path string
}

func NewPatch(op jsondiff.Operation) Patch {
	return Patch{
		Op:   op.Type,
		Path: string(op.Path),
	}
}

func TestJsondiff(t *testing.T) {
	patch, err := jsondiff.CompareJSONOpts(source, target, jsondiff.Equivalent())
	if err != nil {
		// handle error
	}
	for _, op := range patch {
		fmt.Printf("%s\n", op)
	}
}

func TestJsondiff11(t *testing.T) {
	patch, err := jsondiff.CompareJSONOpts(source, target)
	if err != nil {
		// handle error
	}
	lenth := len(patch)
	ops1 := make([]Patch, lenth)
	ops2 := make([]Patch, lenth)
	ops3 := make([]Patch, lenth)
	for i, op := range patch {
		ops1[i] = NewPatch(op)
		ops3[i] = NewPatch(op)
	}
	for i := range patch {
		ops2[i] = NewPatch(patch[lenth-i-1])
	}
	fmt.Println(ops1)
	sort.Slice(ops1, func(i, j int) bool {
		if ops1[i].Op <= ops1[j].Op {
			return true
		}
		if ops1[i].Path <= ops1[j].Path {
			return true
		}
		return false
	})
	fmt.Println(ops1)
	fmt.Println(reflect.DeepEqual(ops1, ops2))
	m1, _ := json.Marshal(ops1)
	m2, _ := json.Marshal(ops2)
	fmt.Println(string(m1))
	fmt.Println(string(m2))

}

func BenchmarkJsondiff(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jsondiff.CompareJSONOpts(source, target, jsondiff.Equivalent())
	}
}

func BenchmarkJsonpatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jsonpatch.Equal(source, target)
	}
}

func BenchmarkJsonpath(b *testing.B) {
	byte1 := []byte(source)
	byte2 := []byte(target)
	var o1 interface{}
	var o2 interface{}
	json.Unmarshal(byte1, &o1)
	json.Unmarshal(byte2, &o2)
	for i := 0; i < b.N; i++ {
		jp.Equal(o1, o2)
	}
}
