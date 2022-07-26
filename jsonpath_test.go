package json_compare

import (
	"encoding/json"
	jp "github.com/zucong/jsonpath"
	"testing"
)

func TestJsonpathGet(t *testing.T) {
	name := "Filter expression with addition"
	expr := `$[?(@.key+50==100)]`
	jsondata := `[{"key": 60}, {"key": 50}, {"key": 10}, {"key": -50}, {"key+50": 100}]`
	jsonObj := jp.ConvertToJsonObj(jsondata)
	j, err := jp.New(name, expr)
	if err != nil {
		t.Errorf("[⛔️parser error] when create jsonpath(%s)=%s: %v", name, expr, err)
		return
	}
	j.InitData(jsonObj)
	jsonpathResult, _ := j.Get()
	resultJsonBytes, _ := json.Marshal(jsonpathResult)
	t.Log(string(resultJsonBytes))
}

func TestJsonpathSet(t *testing.T) {
	name := "multi-level virtual elements with data"
	expr := `$.a.b[:].e1`
	data := `{"a":{"b":[{"e1":"1","e2":"2"},{"e2":"2"}]}}`
	var change interface{} = 0
	j, err := jp.New(name, expr)
	if err != nil {
		t.Fatalf("cannot parse jsonpath")
	}
	jsonObj := jp.ConvertToJsonObj(data)
	j.InitData(jsonObj)
	j.Set(change)
	marshal1, _ := json.Marshal(j.Data())
	marshal2, _ := json.Marshal(jsonObj)
	t.Logf("success: %s", marshal1)
	t.Logf("success: %s", marshal2)

}

func TestJsonpathEqual(t *testing.T) {
	data1 := `{"a":{"b":[{"e1":1,"e2":"2"},{"e1":3,"e2":"4"}]}}`
	data2 := `{"a":{"b":[{"e1":3,"e2":"4"},{"e1":1,"e2":"2"}]}}`
	var obj1 interface{}
	var obj2 interface{}
	json.Unmarshal([]byte(data1), &obj1)
	json.Unmarshal([]byte(data2), &obj2)
	t.Log(jp.Equal(obj1, obj2))
}
