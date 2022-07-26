package json_compare

import (
	"fmt"
	jsonpatch "github.com/evanphx/json-patch/v5"
	"testing"
)

func TestJsonpatch(t *testing.T) {
	original := []byte(`
		{
			"age": 24,
			"height": 3.21,
			"name": "John"
			"friends":[
				{"name":"a",age:1},
				{"name":"b",age:2}
			]
		}
	`)
	similar := []byte(`
		{
			"age": 24,
			"height": 3.21,
			"name": "John"
			"friends":[
				{"name":"a",age:1},
				{"name":"b",age:2}
			]
		}
	`)
	different := []byte(`{"name": "Jane", "age": 20, "height": 3.37}`)

	if jsonpatch.Equal(original, similar) {
		fmt.Println(`"original" is structurally equal to "similar"`)
	}

	if !jsonpatch.Equal(original, different) {
		fmt.Println(`"original" is _not_ structurally equal to "different"`)
	}
}
