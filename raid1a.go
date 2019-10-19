package main

import "github.com/01-edu/z01"

<<<<<<< HEAD
func Raid1a(x,y int) {
=======
func Raid1a(x, y int) {
>>>>>>> 89b96d79575da0b9b611a588b4837559dfdf6f9d
	if x < 0 || y < 0 {
	} else {
		if x == 0 || y == 0 {
		} else {
			if x == 1 && y == 1 {
				z01.PrintRune('o')
				z01.PrintRune(10)
			} else {
				if y == 1 {
<<<<<<< HEAD
					for i:=1; i <= x; i++ {
=======
					for i := 1; i <= x; i++ {
>>>>>>> 89b96d79575da0b9b611a588b4837559dfdf6f9d
						if i == 1 || i == x {
							z01.PrintRune('o')
						} else {
							z01.PrintRune('-')
						}
					}
					z01.PrintRune(10)
				} else {
<<<<<<< HEAD
					for i:=1; i <= x; i++ {
=======
					for i := 1; i <= x; i++ {
>>>>>>> 89b96d79575da0b9b611a588b4837559dfdf6f9d
						if i == 1 || i == x {
							z01.PrintRune('o')
						} else {
							z01.PrintRune('-')
						}
					}
					z01.PrintRune(10)
<<<<<<< HEAD
					for j:=1; j <= y-2; j++ {
						for k:=1; k <= x; k++ {
=======
					for j := 1; j <= y-2; j++ {
						for k := 1; k <= x; k++ {
>>>>>>> 89b96d79575da0b9b611a588b4837559dfdf6f9d
							if k == 1 || k == x {
								z01.PrintRune('|')
							} else {
								z01.PrintRune(' ')
							}
						}
						z01.PrintRune(10)
					}
<<<<<<< HEAD
					for i:=1; i <= x; i++ {
=======
					for i := 1; i <= x; i++ {
>>>>>>> 89b96d79575da0b9b611a588b4837559dfdf6f9d
						if i == 1 || i == x {
							z01.PrintRune('o')
						} else {
							z01.PrintRune('-')
						}
					}
					z01.PrintRune(10)
<<<<<<< HEAD
				}	 
=======
				}
>>>>>>> 89b96d79575da0b9b611a588b4837559dfdf6f9d
			}
		}
	}
}
