package solver

import (
	"fmt"
	"strings"
	"testing"
)

func TestSolver(t *testing.T) {

	//	brd := strings.NewReader("1,0,0,0,0,0,0,0,0\n2,0,0,0,0,0,0,0,0\n3,0,0,0,0,0,0,0,0\n4,0,0,0,0,0,0,0,0\n5,0,0,0,0,0,0,0,0\n6,0,0,0,0,0,0,0,0\n7,0,0,0,0,0,0,0,0\n8,0,0,0,0,0,0,0,0\n9,0,0,0,0,0,0,0,0")
	//brd2 := strings.NewReader(
	//	"8,0,0, 0,0,0, 0,0,0\n" +
	//		"0,0,3, 6,0,0, 0,0,0\n" +
	//		"0,0,7, 0,0,0, 0,0,0\n" +
	//		"0,5,0, 0,0,7, 0,0,0\n" +
	//		"0,0,0, 0,4,5, 7,0,0\n" +
	//		"0,0,0, 1,0,0, 0,3,0\n" +
	//		"0,0,1, 0,0,0, 0,6,8\n" +
	//		"0,0,8, 5,0,0, 0,1,0\n" +
	//		"0,0,0, 0,0,0, 4,0,0")

	//Arto Inkala
	brd := strings.NewReader(
		"8,0,0, 0,0,0, 0,0,0\n" +
			"0,0,3, 6,0,0, 0,0,0\n" +
			"0,7,0, 0,9,0, 2,0,0\n" +
			"0,5,0, 0,0,7, 0,0,0\n" +
			"0,0,0, 0,4,5, 7,0,0\n" +
			"0,0,0, 1,0,0, 0,3,0\n" +
			"0,0,1, 0,0,0, 0,6,8\n" +
			"0,0,8, 5,0,0, 0,1,0\n" +
			"0,9,0, 0,0,0, 4,0,0")

	//bArtoInkala := Board{}
	//_ = bArtoInkala.Read(brd)

	b := Board{}
	err := b.Read(brd)

	if err != nil {
		t.Errorf("board was not read  %#v", b)
	}

	b.Print()
	b.Solve(0)
	fmt.Println("Tracebacks", b.Traceback)

	b.Print()
}
