package util

import (
	"fmt"
	"github.com/gookit/goutil/arrutil"
	"testing"
)

func TestName2(t *testing.T) {

	indexs := []int{0, 1, 2}
	for i := range indexs {
		fmt.Printf("%c", indexs[i]+65)
		//strutil.Uint(fmt.Printf("%T", num))
	}

	fmt.Println(arrutil.ConvType(indexs, "string"))
	//fmt.Println('A')
	//fmt.Println(strconv.Atoi("65"))
}
