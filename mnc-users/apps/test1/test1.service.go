package test1

import (
	"fmt"
	"strings"
	"time"
)

type Test1Service struct{}

func NewTest1Service() *Test1Service {
	return &Test1Service{}
}

func (s *Test1Service) FindMatchString(n int, stringsList []string) interface{} {
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if strings.ToLower(stringsList[i]) == strings.ToLower(stringsList[j]) {
				return []int{i + 1, j + 1}
			}
		}
	}
	return false
}

func (s *Test1Service) CalculateChange(totalBelanja, dibayarkan int) interface{} {
	if dibayarkan < totalBelanja {
		return false
	}
	pecahan := []int{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}
	kembalian := dibayarkan - totalBelanja
	kembalian = (kembalian / 100) * 100
	result := map[int]int{}
	for _, p := range pecahan {
		if kembalian >= p {
			result[p] = kembalian / p
			kembalian %= p
		}
	}
	return result
}

func (s *Test1Service) IsValidBracketSequence(input string) bool {
	stack := []rune{}
	pairs := map[rune]rune{'{': '}', '[': ']', '<': '>'}
	for _, char := range input {
		if closing, found := pairs[char]; found {
			stack = append(stack, closing)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != char {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func (s *Test1Service) CheckLeave(holidayQuota int, joinDateStr, leaveDateStr string, duration int) (bool, string) {
	joinDate, _ := time.Parse("2006-01-02", joinDateStr)
	leaveDate, _ := time.Parse("2006-01-02", leaveDateStr)
	startLeaveDate := joinDate.AddDate(0, 0, 180)
	n := leaveDate.YearDay() - startLeaveDate.YearDay()
	if leaveDate.Before(startLeaveDate) {
		return false, "Belum 180 hari sejak tanggal join"
	}
	annualLeave := 14 - holidayQuota
	availableLeave := (n * annualLeave) / 365
	if availableLeave < duration {
		return false, fmt.Sprintf("Hanya boleh mengambil %d hari cuti", availableLeave)
	}
	if duration > 3 {
		return false, "Cuti maksimal 3 hari berturut-turut"
	}
	return true, "Dapat mengambil cuti"
}