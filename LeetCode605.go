package main

import (
	"fmt"
)

func CanPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	length := len(flowerbed)
	if length == 1 {
		if n == 1 && flowerbed[0] == 0 {
			return true
		}
		return false
	}

	for i := 0; i < len(flowerbed); i++ {
		if n <= 0 {
			break
		}
		if flowerbed[i] != 1 {
			if i == 0 {
				if flowerbed[i+1] != 1 {
					flowerbed[i] = 1
					n--
				}
			} else if i == len(flowerbed)-1 {
				if flowerbed[i-1] != 1 {
					flowerbed[i] = 1
					n--
				}
			} else if flowerbed[i-1] != 1 && flowerbed[i+1] != 1 {
				flowerbed[i] = 1
				n--
			}
		}
	}
	fmt.Println(flowerbed)
	return n == 0
}
