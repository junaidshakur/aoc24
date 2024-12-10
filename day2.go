package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func day2_task() {
	pureSafeReportCount := 0
	minorSafeReportCount := 0

	f, err := os.Open("./inputs/2.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		reportData := strings.Split(s.Text(), " ")
		isValid, canBeValid := is_safe(reportData)
		if isValid {
			pureSafeReportCount++
		} else if canBeValid {
			if is_safe_bf_main(reportData) {
				minorSafeReportCount++
			}
		}
	}

	fmt.Println("Total Safe Count:", pureSafeReportCount+minorSafeReportCount, ", Pure Safe count:", pureSafeReportCount, ", Minor Safe count:", minorSafeReportCount)
}

func is_safe(data []string) (bool, bool) {
	last, _ := strconv.Atoi(data[0])
	second, _ := strconv.Atoi(data[1])
	upword := last-second > 0
	switchDirection := 0
	diffFalseCount := 0

	for i := 1; i < len(data); i++ {
		current, _ := strconv.Atoi(data[i])
		diff := last - current

		if upword {
			if diff < 0 {
				switchDirection++
				upword = false
			}
		} else {
			if diff > 0 {
				switchDirection++
				upword = true
			}
		}

		if diff < 0 {
			diff *= -1
		}
		if diff < 1 || diff > 3 {
			diffFalseCount++
		}
		last = current
	}
	if diffFalseCount == 0 && switchDirection == 0 {
		return true, true
	} else if diffFalseCount > 2 || switchDirection > 2 {
		return false, false
	} else {
		return false, true
	}

}

func is_safe_bf_main(data []string) bool {

	for i := 0; i < len(data); i++ {
		if is_safe_bf_loop(data, i) {
			return true
		}
	}
	return false
}

func is_safe_bf_loop(data []string, k int) bool {

	firstIndex := 0
	secondIndex := 1
	if k == 0 {
		firstIndex = 1
		secondIndex = 2
	}
	if k == 1 {
		secondIndex = 2
	}
	last, _ := strconv.Atoi(data[firstIndex])
	second, _ := strconv.Atoi(data[secondIndex])
	upword := last-second > 0

	i := 1
	if k == 0 {
		i = 2
	}
	for ; i < len(data); i++ {
		if i == k {
			continue
		}
		current, _ := strconv.Atoi(data[i])
		diff := last - current

		if upword {
			if diff < 0 {
				return false
			}
		} else {
			if diff > 0 {
				return false
			}
		}

		if diff < 0 {
			diff *= -1
		}
		if diff < 1 || diff > 3 {
			return false
		}
		last = current
	}

	return true
}
