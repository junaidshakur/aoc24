package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func day3_task1() {

	f, err := os.Open("./inputs/3.txt")
	if err != nil {
		log.Fatal(err)
	}

	var sum int64 = 0
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		sum += scan_and_sum(line)
	}
	fmt.Println("Multiplication sum:", sum)

}

func scan_and_sum(line string) int64 {
	var sum int64 = 0
	validMulCount := 0

	tok := strings.Split(line, "mul(")
	for i := 1; i < len(tok); i++ {

		subTok := strings.Split(tok[i], ")")
		if len(subTok) > 0 && len(subTok[0]) <= 7 {
			nums := strings.Split(subTok[0], ",")
			if len(nums) == 2 {
				num1, err := strconv.Atoi(nums[0])
				if err == nil {
					num2, err := strconv.Atoi(nums[1])
					if err == nil {
						fmt.Println(num1, num2)
						sum += int64(num1 * num2)
						validMulCount++
					}
				}
			}
		}
	}
	return sum
}
