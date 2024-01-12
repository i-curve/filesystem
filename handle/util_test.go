package handle

import (
	"fmt"
	"testing"
)

func TestListdir(t *testing.T) {
	res, _ := listDir("/bucket")
	for _, v := range res {
		fmt.Printf("%+v\n", v)
	}
}
