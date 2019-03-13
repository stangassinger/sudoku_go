package main

import "fmt"

type  T_Sudoku_T   [81]int

/*
fn pretty_print(sudoku_ar: SudokuArType) {
	let line_sep = "------+------+------";
	println!("{}", line_sep);
	for i in 0..sudoku_ar.len() {
	    print!("{} ", sudoku_ar[i]);
	    if (i + 1) % 3 == 0 && !((i + 1) % 9 == 0) {
		print!("| ");
	    }
	    if (i + 1) % 9 == 0 {
		println!(" ");
	    }
	    if (i + 1) % 27 == 0 {
		println!("{}", line_sep);
	    }
	}
    }
}
*/

func pretty_print(sudoku_ar *T_Sudoku_T) {
	line_sep := "------+------+------";
	fmt.Println( line_sep )
	for i, v := range *sudoku_ar {
			fmt.Print ( v , " " )
			if (i + 1) % 3 == 0 || (i + 1) % 9 == 0 {
				fmt.Print("| ");
			} 
			
			if (i + 1) % 9 == 0 {
				fmt.Println(" ");
			}
			
			if (i + 1) % 27 == 0 {
					fmt.Println( line_sep);
			}
	}
}



func main() {
	sudoku_ar := T_Sudoku_T {     
        8, 5, 0, 0, 0, 2, 4, 0, 0, 7, 2, 0, 0, 0, 0, 0, 0, 9, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0,
        1, 0, 7, 0, 0, 2, 3, 0, 5, 0, 0, 0, 9, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 0,
        0, 7, 0, 0, 1, 7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 6, 0, 4, 0,
	}
	fmt.Println(sudoku_ar) 

	pretty_print( &sudoku_ar )
}