package goroutine

import (
	"fmt"
	"testing"
)

func TestGoFunc(t *testing.T) {
	GoFunc(5, func(i ...interface{}) {
		fmt.Println("==========================")
		fmt.Println(i[0])
		fmt.Println(i[1])
		fmt.Println("==========================")
	}, map[interface{}]interface{}{
		"1":  "amdkadka",
		"12": "amdkadka",
		"13": "amdkadka",
		"14": "amdkadka",
		"15": "amdkadka",
		"16": "amdkadka",
		"161": "amdkadka",
		"162": "amdkadka",
		"163": "amdkadka",
		"164": "amdkadka",
		"165": "amdkadka",
		"166": "amdkadka",
		"167": "amdkadka",
		"168": "amdkadka",
		"1689": "amdkadka",
	})
}
