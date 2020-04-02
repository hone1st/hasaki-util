package array_util

import (
	"fmt"
	"testing"
)

func TestArrayMultiSort(t *testing.T) {

	/**
			    ['name' => 'book', 'price' => 100],
	            ['name' => 'pen', 'price' => 50],
	            ['name' => 'dress', 'price' => 500],
	            ['name' => 'shoes', 'price' => 700],
	            ['name' => 'chocolate', 'price' => 200]
	*/

	dest := make([]map[string]interface{}, 0)
	dest = append(dest, map[string]interface{}{
		"name":  "book",
		"price": 100,
	})
	dest = append(dest, map[string]interface{}{
		"name":  "pen",
		"price": 50,
	})
	dest = append(dest, map[string]interface{}{
		"name":  "dress",
		"price": 500,
	})
	dest = append(dest, map[string]interface{}{
		"name":  "shoes",
		"price": 700,
	})
	dest = append(dest, map[string]interface{}{
		"name":  "chocolate",
		"price": 200,
	})

	//sortArr := ArrayColumnInt(dest, "price")
	//sortArr1 := ArrayColumnString(dest, "name")
	//fmt.Println(sortArr)
	//fmt.Println(sortArr1)
	//sort.Ints(sortArr)
	//sort.Strings(sortArr1)
	//fmt.Println(sortArr1)
	//fmt.Println(sortArr)

	fmt.Println(ArrayMultiSort("name", MultiSortDesc, dest))
	//ArrayMultiSort(sort,1)
}
