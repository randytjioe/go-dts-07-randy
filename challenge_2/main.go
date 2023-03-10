package main

import "fmt"

func main() {
	// Menggunakan perulangan for untuk mencetak nilai i dari 0 hingga 4
	for i := 0; i < 5; i++ {
		fmt.Printf("Nilai i = %d\n", i)
	}

	// Menggunakan perulangan for dan if-else untuk mencetak nilai j dari 0 hingga 10
	for j := 0; j <= 10; j++ {
		if j < 5 {
			fmt.Printf("Nilai j = %d\n", j)
		} else if j == 5 {
			fmt.Printf("Character U+0421 'C' starts at byte position 0\n")
		} else if j == 6 {
			fmt.Printf("Character U+0410 'A' starts at byte position 2\n")
		} else if j == 7 {
			fmt.Printf("Character U+0428 'III' starts at byte position 4\n")
		} else if j == 8 {
			fmt.Printf("Character U+0410 'A' starts at byte position 6\n")
		} else if j == 9 {
			fmt.Printf("Character U+0420 'P' starts at byte position 8\n")
		} else if j == 10 {
			fmt.Printf("Character U+0412 'B' starts at byte position 10\n")
			fmt.Printf("Character U+041E 'O' starts at byte position 12\n")
		}
	}
	for j := 6; j <= 10; j++ {
		fmt.Printf("Nilai j = %d\n", j)
	}

}
