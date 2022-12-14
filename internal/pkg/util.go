//go:generate mockgen -source ./util.go -destination=../mocks/ut.go -package=mock_util
package pkg

import "errors"

func ReverseInt(x interface{}) (int, error) {
	y, ok := x.(int)
	if !ok {
		return 0, errors.New("not int")
	}

	var result int
	for y != 0 {
		result = result*10 + y%10
		if result > 2147483647 || result < -2147483648 {
			return 0, nil
		}

		y /= 10
	}

	return result, nil
}

// ContainsDuplicate - check int array for duplicate. For example [1, 1, 2, 3] -> true
func ContainsDuplicate(nums []int) bool {
	set := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := set[v]; ok {
			return true
		}

		set[v] = struct{}{}
	}

	return false
}

// IsPalindrome - check integer for polindrome. For example 101 -> true, 123 -> false
func IsPalindrome(x int) (bool, bool) {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false, false
	}

	var revertedNum int
	for x > revertedNum {
		last := x % 10
		x /= 10
		revertedNum = revertedNum*10 + last
	}

	return x == revertedNum || x == revertedNum/10, false
}
