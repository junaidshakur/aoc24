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
	enabled := true
	for s.Scan() {
		line := s.Text()
		lineSum, en := scan_and_sum(line, enabled)
		sum += lineSum
		enabled = en
	}
	fmt.Println("Multiplication sum:", sum)

}

func scan_and_sum(line string, i_enabled bool) (int64, bool) {
	var sum int64 = 0
	validMulCount := 0
	enabled := i_enabled

	tok := strings.Split(line, "mul(")

	enabledIndex := strings.LastIndex(tok[0], "do()")
	disabledIndex := strings.LastIndex(tok[0], "don't()")
	if disabledIndex > enabledIndex {
		enabled = false
	}

	println("zero tok :: enabledIndex:", enabledIndex, ", disabledIndex:", disabledIndex)

	for i := 1; i < len(tok); i++ {
		if enabled {
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
		enabledIndex := strings.LastIndex(tok[i], "do()")
		disabledIndex := strings.LastIndex(tok[i], "don't()")
		if disabledIndex > enabledIndex {
			enabled = false
		} else if enabledIndex > disabledIndex {
			enabled = true
		}
		//println(i, " tok :: enabledIndex:", enabledIndex, ", disabledIndex:", disabledIndex)
	}
	println("valid Mul count", validMulCount)
	return sum, enabled
}
