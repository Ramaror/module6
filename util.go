package main

import "errors"

type Num struct {
	Number int
}

func (x Num) X() int {
	return x.Number * 1
}

type x interface {
	x() int
	s() error
}

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