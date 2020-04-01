package int

import (
	"fmt"
	"strconv"
)

// left   左操作符
// right  右操作符
// scales 保留的小数点数
func bc(left, right string, scales ...int) (string, float64, float64, error) {
	var scale int = 0
	if len(scales) > 0 {
		scale = scales[0]
	}
	leftF, err := strconv.ParseFloat(left, 64)
	if err != nil {
		return "", 0, 0, err
	}
	rightF, err := strconv.ParseFloat(right, 64)
	if err != nil {
		return "", 0, 0, err
	}
	format := fmt.Sprintf("0.%df", scale)

	return format, leftF, rightF, nil
}

//字符串加法
func BcAdd(left, right string, scales ...int) (string, error) {
	format, leftF, rightF, err := bc(left, right, scales...)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%"+format, leftF+rightF), nil
}

//字符串除法
func BcDiv(left, right string, scales ...int) (string, error) {
	format, leftF, rightF, err := bc(left, right, scales...)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%"+format, leftF/rightF), nil
}

//字符串减法
func BcSub(left, right string, scales ...int) (string, error) {
	format, leftF, rightF, err := bc(left, right, scales...)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%"+format, leftF-rightF), nil
}

//字符串乘法
func BcMul(left, right string, scales ...int) (string, error) {
	format, leftF, rightF, err := bc(left, right, scales...)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%"+format, leftF*rightF), nil
}
