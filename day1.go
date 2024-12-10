package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func day1_task1() {
	arr1 := [1000]int{}
	arr2 := [1000]int{}
	i := 0

	f, err := os.Open("./inputs/1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		txt := strings.Split(s.Text(), " ")
		//println(len(txt), txt[0], txt[3])
		arr1[i], _ = strconv.Atoi(txt[0])
		arr2[i], _ = strconv.Atoi(txt[3])
		i++
	}

	sort.Ints(arr1[:])
	sort.Ints(arr2[:])
	sum := 0
	for i := 0; i < 1000; i++ {

		diff := arr1[i] - arr2[i]
		if diff < 0 {
			diff *= -1
		}
		//println(diff, arr1[i], arr2[i])
		sum += diff
	}
	println("Sum of difference is:", sum)
}

func day1_task2() {
	arr1 := [1000]int{}
	map1 := make(map[int]int)
	i := 0

	f, err := os.Open("./inputs/1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		txt := strings.Split(s.Text(), " ")
		//println(len(txt), txt[0], txt[3])
		arr1[i], _ = strconv.Atoi(txt[0])
		num2, _ := strconv.Atoi(txt[3])
		map1[num2] = (map1[num2] + 1)
		i++
	}

	similaritySum := 0
	for i := 0; i < 1000; i++ {

		c := map1[arr1[i]]
		if c > 0 {
			similaritySum += (c * arr1[i])
		}
		//println(c, arr1[i])
	}

	println("Similarity Sum is:", similaritySum)

}
