package json_compare_test

import (
	"fmt"
	"github.com/iostrovok/go-jsoncompare/jsoncompare"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
)

var f1 = "data1/a.json"
var f2 = "data1/b.json"

func TestJsoncompare(t *testing.T) {
	b1, e1 := os.ReadFile(f1)
	if e1 != nil {
		log.Fatalln(e1)
	}

	b2, e2 := os.ReadFile(f2)
	if e2 != nil {
		log.Fatalln(e2)
	}

	list, err := jsoncompare.Compare(b1, b2)
	if err != nil {
		log.Fatalln(err)
	}

	leftOnly, rightOnly, noEqual, goodList := jsoncompare.SplitBySide(list)

	printList("GOOD: ", goodList)
	printList("Left Only: ", leftOnly)
	printList("Right Only: ", rightOnly)
	printList("No Equal: ", noEqual)
}

func printList(suff string, list []*jsoncompare.PathDiff) {
	for i, v := range list {
		sing := "!="
		if v.IsEqual {
			sing = "=="
		}
		fmt.Printf("%d. %s. %s <%s> %s\n", i, suff, v.PathLeft, sing, v.PathRight)
	}
}

func loadUrl(url string) ([]byte, error) {
	resp, err_get := http.Get(url)
	if err_get != nil {
		return nil, err_get
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}