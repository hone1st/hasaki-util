package array_util

import (
	"reflect"
	"sort"
	"strings"
)

type MultiSortType int

const (
	MultiSortAsc  MultiSortType = 0
	MultiSortDesc MultiSortType = 1
)

type MultiSort struct {
	Key   interface{}
	Value map[string]interface{}
	Flag  bool
}

type MultiSortList []*MultiSort

func (m MultiSortList) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
func (m MultiSortList) Len() int      { return len(m) }
func (m MultiSortList) Less(i, j int) bool {
	var sort bool = m[i].Flag
	kind1 := reflect.TypeOf(m[i].Key)
	kind2 := reflect.TypeOf(m[j].Key)

	if kind1 != kind2 {
		return sort
	}
	switch kind1.Kind() {
	case reflect.String:
		va := strings.Compare(m[i].Key.(string), m[j].Key.(string))
		if va == -1 {
			sort = !sort
		}
		break
	case reflect.Int:
		if m[i].Key.(int) < m[j].Key.(int) {
			sort = !sort
		}
		break
	case reflect.Float64:
		if m[i].Key.(float64) < m[j].Key.(float64) {
			sort = !sort
		}
	default:
		break

	}
	return sort
}

func ArrayMultiSort(key string, flag MultiSortType, params []map[string]interface{}) []map[string]interface{} {
	sortType := false
	if flag == MultiSortDesc {
		sortType = true
	}
	m := make(MultiSortList, len(params))
	for i, v := range params {
		m[i] = &MultiSort{
			Value: v,
			Flag:  sortType,
		}
		if val, ex := v[key]; ex {
			m[i].Key = val
		} else {
			m[i].Key = ""
		}

	}
	sort.Sort(m)

	re := make([]map[string]interface{}, 0)

	for _, v := range m {
		re = append(re, v.Value)
	}
	return re
}

// 获取一维数组的某个key的值
func ArrayColumnInt(mp []map[string]interface{}, key string) []int {
	var result = make([]int, 0)
	for _, k := range mp {
		if v, ex := k[key]; ex && reflect.TypeOf(v).Kind() == reflect.Int {
			result = append(result, v.(int))
		}
	}
	return result
}

// 获取一维数组的某个key的值
func ArrayColumnFloat(mp []map[string]interface{}, key string) []float64 {
	var result = make([]float64, 0)
	for _, k := range mp {
		if v, ex := k[key]; ex && reflect.TypeOf(v).Kind() == reflect.Float64 {
			result = append(result, v.(float64))
		}
	}
	return result
}

// 获取一维数组的某个key的值
func ArrayColumnString(mp []map[string]interface{}, key string) []string {
	var result = make([]string, 0)
	for _, k := range mp {
		if v, ex := k[key]; ex && reflect.TypeOf(v).Kind() == reflect.String {
			result = append(result, v.(string))
		}
	}
	return result
}
