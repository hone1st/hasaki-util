package string

import (
	"fmt"
	"testing"
)

func TestContains(t *testing.T) {
	var s string = "十大看到大家安静到付afadadadasdka单靠打卡"
	fmt.Println(Contains(s, true, "afada", "单靠打卡", "打卡"))
	fmt.Println(Contains(s, true, "afada", "单靠打卡", "打卡1"))
	fmt.Println(Contains(s, false, "afada", "单靠打卡", "打卡1"))
}

func TestReplace(t *testing.T) {
	var s string = "十大看到大家安静到付afadadadasdka单靠打卡"
	fmt.Println(Replace(s, "哈哈", "afada", "单靠打卡", "打卡"))
}


func TestSystemSplit(t *testing.T) {
	var s1 string = "C:\\Users\\admin\\Desktop\\longyou\\application"
	var s2 string = "/bin/sh/sad/adad"
	fmt.Println(SystemSplit(s1))
	fmt.Println(SystemSplit(s2))
}