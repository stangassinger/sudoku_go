package main

import "fmt"

type T_Sudoku_T [81]int

func check_validity(val int, x int, y int, sudoku_ar *T_Sudoku_T) bool {
	for i := 0; i <= 8; i++ {
		if (sudoku_ar[y*9+i] == val) || (sudoku_ar[i*9+x] == val) {
			return false
		}
	}

	startx := (x / 3) * 3
	starty := (y / 3) * 3
	for i := starty; i <= (starty + 2); i++ {
		for j := startx; j <= (startx + 2); j++ {
			if sudoku_ar[i*9+j] == val {
				return false
			}
		}
	}

	return true
}

func place_number(pos int, sudoku_ar *T_Sudoku_T) bool {
	var ret bool
	if pos == 81 {
		return true
	}

	if sudoku_ar[pos] > 0 {
		ret = place_number(pos+1, sudoku_ar)
		if ret == true {
			return true
		} else {
			return false
		}
	}

	for n := 1; n <= 9; n++ {
		if check_validity(n, pos%9, pos/9, sudoku_ar) == true {
			sudoku_ar[pos] = n
			ret = place_number(pos+1, sudoku_ar)
			if ret == true {
				return true
			}
			sudoku_ar[pos] = 0
		}
	}
	return false

}

func solve(sudoku_ar *T_Sudoku_T) bool {
	return place_number(0, sudoku_ar)
}

func pretty_print(sudoku_ar *T_Sudoku_T) {
	line_sep := "------+------+------"
	fmt.Println(line_sep)
	for i, v := range *sudoku_ar {
		fmt.Print(v, " ")
		if (i+1)%3 == 0 || (i+1)%9 == 0 {
			fmt.Print("| ")
		}

		if (i+1)%9 == 0 {
			fmt.Println(" ")
		}

		if (i+1)%27 == 0 {
			fmt.Println(line_sep)
		}
	}
}

func main() {
	sudoku_ar := T_Sudoku_T{
		8, 5, 0, 0, 0, 2, 4, 0, 0, 7, 2, 0, 0, 0, 0, 0, 0, 9, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		1, 0, 7, 0, 0, 2, 3, 0, 5, 0, 0, 0, 9, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 0,
		0, 7, 0, 0, 1, 7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 6, 0, 4, 0,
	}
	fmt.Println("  ---- Sudoku ----  ")
	pretty_print(&sudoku_ar)
	fmt.Println("Solving...")
	if solve(&sudoku_ar) == false {
		fmt.Println("Unsolvable")
	} else {
		pretty_print(&sudoku_ar)
	}
}
