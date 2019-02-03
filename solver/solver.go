package solver

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
)

//Board to play with
type Board struct {
	Cells      [9][9]int
	Contraints [9][9]CellConstain
	Traceback  int
}

type CellConstain struct {
	Allowed [10]bool
}

type Assumption struct {
	I int
	J int
	V int
}

func (b *Board) Print() {

	for i := 0; i < 9; i++ {

		fmt.Println(b.Cells[i][0], b.Cells[i][1], b.Cells[i][2], "|", b.Cells[i][3], b.Cells[i][4], b.Cells[i][5], "|", b.Cells[i][6], b.Cells[i][7], b.Cells[i][8])
		if i == 2 || i == 5 {
			fmt.Println("------+-------+------")
		}
	}

}

func (b *Board) Read(r io.Reader) error {

	cr := csv.NewReader(r)
	cr.FieldsPerRecord = 9
	i := 0

	for {
		j := 0
		line, crerr := cr.Read()
		//fmt.Println(line)
		if crerr != nil {
			//fmt.Println(crerr)
			if i > 8 {
				return nil
			}
			return crerr
		}

		for j < 9 {

			b.Cells[i][j], _ = strconv.Atoi(strings.Trim(line[j], " "))
			j = j + 1
		}
		i++
	}
}

//FindNextAmptyCell fine the next empty cell on the board
func (b *Board) FindNextEmptyCell() (r int, c int) {

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b.Cells[i][j] == 0 {
				return i, j

			}
		}
	}

	return -1, -1
}

//FindNextAmptyCell fine the next empty cell on the board
func (b *Board) IsValid() (bool, bool) {

	zeros := false

	for i := 0; i < 9; i++ {
		seen := make([]bool, 10)

		for j := 0; j < 9; j++ {

			v := b.Cells[i][j]
			if v > 0 {
				if seen[v] {
					return false, zeros
				}
				seen[v] = true
			} else {
				zeros = true
			}

		}
	}

	for j := 0; j < 9; j++ {
		seen := make([]bool, 10)

		for i := 0; i < 9; i++ {

			v := b.Cells[i][j]
			if v > 0 {
				if seen[v] {
					return false, zeros
				}
				seen[v] = true
			} else {
				zeros = true
			}

		}
	}

	for s := 0; s < 9; s++ {
		i0 := (s / 3) * 3
		j0 := (s % 3) * 3

		seen := make([]bool, 10)

		for i := i0; i < i0+3; i++ {
			for j := j0; j < j0+3; j++ {

				v := b.Cells[i][j]
				if v > 0 {
					if seen[v] {
						return false, zeros
					}
					seen[v] = true
				} else {
					zeros = true
				}

			}

		}
	}

	return true, zeros
}

func (b *Board) GetConstrain(ix int, jx int) []int {

	var res []int

	//fmt.Printf(" (%d,%d,?) ", ix, jx)

	//fmt.Printf("i %d ", len(b.Cells))

	if b.Cells[ix][jx] != 0 {
		return res
	}

	//row constrain
	seen := make([]bool, 10)
	for j := 0; j < 9; j++ {

		if j != jx {
			v := b.Cells[ix][j]
			if v > 0 {
				seen[v] = true
			}
		}
	}

	//col constrain
	for i := 0; i < 9; i++ {

		if i != ix {
			v := b.Cells[i][jx]
			if v > 0 {
				seen[v] = true
			}
		}
	}

	//sec constrain
	i0 := (ix / 3) * 3
	j0 := (jx / 3) * 3

	for i := i0; i < i0+3; i++ {
		for j := j0; j < j0+3; j++ {

			if i != ix && j != jx {
				v := b.Cells[i][j]
				if v > 0 {
					seen[v] = true
				}
			}
		}
	}

	for z := 1; z < 10; z++ {
		if !seen[z] {
			res = append(res, z)
		}
	}

	return res
}

func (b *Board) MakeAssumptions(ix int, jx int, t int) []Assumption {
	var ass []Assumption

	b.Cells[ix][jx] = t

	ass = append(ass, Assumption{ix, jx, t})

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if !(ix == i && jx == j) {
				func(ik int, jk int) {
					con := b.GetConstrain(ik, jk)
					if len(con) == 1 {
						b.Cells[ik][jk] = con[0]
						ass = append(ass, Assumption{ik, jk, con[0]})
					}

				}(i, j)
			}
		}
	}

	return ass
}

func (b *Board) UndoAssumptions(ass []Assumption) {

	b.Traceback++
	for _, a := range ass {
		b.Cells[a.I][a.J] = 0

	}
	if b.Traceback%1000 == 0 {

		//fmt.Println("Tracebacks", b.Traceback)
		//b.Print()
	}
}

func (b *Board) Solve(d int) bool {

	//fmt.Printf("depth %d \n", d)
	//b.Print()
	for {
		i, j := b.FindNextEmptyCell()

		if i == -1 {
			return false
		}

		//fmt.Printf("try (%d,%d)\n", i, j)

		con := b.GetConstrain(i, j)
		//fmt.Printf("constrains  (%d,%d) %v \n", i, j, con)

		for _, t := range con {
			//for t := 1; t < 10; t++ {

			ass := b.MakeAssumptions(i, j, t)

			val, zeros := b.IsValid()

			if val {

				if !zeros || b.Solve(d+1) {
					return true
				}
			}

			b.UndoAssumptions(ass)

		}

		//fmt.Printf("traceback (%d,%d)\n", i, j)

		return false

	}

}
