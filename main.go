package main

import (
	"strings"
	"regexp"
)

func main(){
	
}

func calculate(name1, name2 string) int{
	
	names := strings.ToLower(name1+name2)
	var numbers []int //slice which holds occurrences of the single letter in names

	for len(names)>1{
		c := names[0] //occurrence of first letter will be counted
		occ := len(regexp.MustCompile(string(c)).FindAllStringIndex(names,-1)) //counts occurrence of c in names
		numbers = append(numbers, occ)

		names = regexp.MustCompile(string(c)).ReplaceAllString(names,"") //removes all cs in names
	}

	if len(names) == 1{
		numbers = append(numbers,1)
	}
	
	return -1 //TODO right return (reduce numbers)
}

func reduceNumbers(numbers1, numbers2 []int)int{
	var result int

	if len(numbers1) <= 2 && len(numbers1) == 0{
		//TODO
	}else if len(numbers1) >= 2{
		//TODO
	}else if len(numbers1) == 1{
		//TODO
	}else if len(numbers1) == 0 && len(numbers2) > 0{
		//TODO
	}

	return -1 //TODO right return
}