package main

import "sort"

type aStructure struct {
	person string
	height int
	weight int
}

func main() {
	myslise := make([]aStructure, 0)

	myslise = append(myslise, aStructure{"dd", 10, 20})

	sort.Slice(myslise, func(i, j int) bool {
		return myslise[i].height > myslise[j].height
	})

}
