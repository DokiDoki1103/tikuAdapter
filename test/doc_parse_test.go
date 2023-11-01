package test

import (
	"code.sajari.com/docconv/v2"
	"fmt"
	"testing"
)

func TestParseDoc(t *testing.T) {

	res, err := docconv.ConvertPath("test.pdf")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res)

}
