package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main(){

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Your Name: ")
	name1,_ := reader.ReadString('\n')
	fmt.Print("Enter Your Crushe's Name: ")
	name2,_ := reader.ReadString('\n')
	name1 = name1[:len(name1)-1]
	name2 = name2[:len(name2)-1]

	result1,err1 := calculate(name1,name2)
	result2,err2 := calculate(name2,name1)

	result := (float32(result1)+float32(result2))/2
	if err1==nil && err2 == nil {
		fmt.Printf("The compatibility of %s and %s is %.1f%%.",name1,name2,result)
	}else{
		fmt.Println(err1)
	}
}

func calculate(name1, name2 string) (int, error){
	
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
	
	return reduceNumbers(numbers,[]int{})
}

func reduceNumbers(numbers1, numbers2 []int)(int,error){
	var result int

	if len(numbers1) <= 2 && len(numbers2) == 0{
		switch len(numbers1){
		case 2: result =  numbers1[0]*10+numbers1[1]
		case 1: result = numbers1[0]
		case 0: result = 0
		default: result = -1
		}
	}else if len(numbers1) >= 2{
		numbers2 = append(numbers2, numbers1[0]+numbers1[len(numbers1)-1])
		numbers1 = numbers1[1:len(numbers1)-1] //shift & pop
		return reduceNumbers(numbers1,numbers2)
	}else if len(numbers1) == 1{
		numbers2 = append(numbers2,numbers1[0])
		numbers1 = []int{}
		return reduceNumbers(numbers2,numbers1)
	}else if len(numbers1) == 0 && len(numbers2) > 0{
		return reduceNumbers(numbers2,numbers1)
	}

	if result == -1{
		return result, errors.New("something went wrong")
	}else if result > 99{
		//int to string
		resultStr := strconv.Itoa(result)
		//string to int slice
		var numbers = []int{}
		for _,e := range resultStr{
			i,_ := strconv.Atoi(string(e))
			numbers = append(numbers,i)
		}
		//reduce again
		return reduceNumbers(numbers,[]int{})
	}
	
	return result, nil
}