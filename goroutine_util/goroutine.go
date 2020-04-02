package goroutine_util

import "sync"

// 开启多协程
// wgs定义执行的协程数
// do 定义传入执行方法 接收interface数组
// data 传入目标数据map
func GoFunc(wgs int, do func(...interface{}), data map[interface{}]interface{}) {
	ch := make(chan []interface{}, len(data))
	wg := sync.WaitGroup{}
	wg.Add(wgs)
	go func() {
		defer wg.Done()
		defer close(ch)
		for k, v := range data {
			ch <- []interface{}{k, v}
		}
	}()
	for i := 0; i < wgs-1; i++ {
		go func() {
			defer wg.Done()
			for {
				select {
				case val, ok := <-ch:
					if !ok {
						return
					}
					do(val[0], val[1])
				}
			}
		}()
	}
	wg.Wait()
}

func GoFuncArr(wgs int, do func(...interface{}), data []map[string]interface{}) {
	ch := make(chan []interface{}, len(data))
	wg := sync.WaitGroup{}
	wg.Add(wgs)
	go func() {
		defer wg.Done()
		defer close(ch)
		for k, v := range data {
			ch <- []interface{}{k, v}
		}
	}()
	for i := 0; i < wgs-1; i++ {
		go func() {
			defer wg.Done()
			for {
				select {
				case val, ok := <-ch:
					if !ok {
						return
					}
					do(val[0], val[1])
				}
			}
		}()
	}
	wg.Wait()
}
