package util

import (
	"fmt"
	"os"
	"testing"
)

func TestMkDir(t *testing.T) {
	MakDir("/temp/ttt.jpg")
}

func TestDel(t *testing.T) {
	// os.Remove("/data/temp/aaa/bbb")
	// fmt.Print(path.Dir("/data/temp/aaa/bbb")
	// DelDir("/data/temp/aaa/bbb")
	fmt.Println(os.Remove("/data"))
}
