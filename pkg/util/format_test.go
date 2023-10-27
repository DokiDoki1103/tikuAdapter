package util

import (
	"fmt"
	"github.com/antlabs/strsim"
	"testing"
)

func TestName2(t *testing.T) {
	fmt.Println(strsim.Compare("2008年9月", "2008", strsim.Simhash()))
	//fmt.Println('A')
	//fmt.Println(strconv.Atoi("65"))
}
