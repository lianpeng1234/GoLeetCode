package main

type TwoSum struct {
	maps map[int]int8
}

func Constructor170() TwoSum {
	return TwoSum{make(map[int]int8)}
}

func (this *TwoSum) Add(number int) {
	this.maps[number] = 1
}

func (this *TwoSum) Find(value int) bool {
	if _, ok := this.maps[value]; ok {
		return ok
	}

	for i := range this.maps {
		if _, ok := this.maps[value-i]; ok {
			return ok
		}
	}

	return false
}
