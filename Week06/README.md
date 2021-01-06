## 作业
参考 Hystrix 实现一个滑动窗口计数器。

## 代码：
```go
package main
import (
	"container/list"
	"sync"
	"time"
)

type census struct {
	Success int
	Fail    int
}

type rollingWindow struct {
	mu         sync.Mutex
	List       *list.List
	TimeSecond int64
	Size       int
	Census     map[int64]*census
}

// 滑动窗口放入当前访问时间戳
func (w *rollingWindow) Add(now int64) {
	w.List.PushFront(now)
}

// 移除滑动窗口最后一个
func (w *rollingWindow) Remove() {
	w.List.Remove(w.List.Back())
}

func (w *rollingWindow) IsSuccess() bool {
	w.mu.Lock()
	now := time.Now().Unix()
	res := false
	if w.List.Len() >= w.Size { // 达到限流次数  判断是否在限制时间内
		if (now - w.List.Back().Value.(int64)) >= w.TimeSecond {
			w.Add(now)
			w.Remove()
			res = true
		}
	} else { //未达到限流次数  写入访问时间戳到滑动窗口中
		res = true
		w.Add(now)
	}

	v1 := w.Census[now]
	if v1 == nil {
		v1 = new(census)
	}
	if res {
		v1.Success++
	} else {
		v1.Fail++
	}
	w.Census[now] = v1
	w.mu.Unlock()
	return res
}
```
## 运行测试
```
=== RUN   TestRollingWindow_IsSuccess
key: 1609939355 success: 5 fail: 55
<<================================>>
key: 1609939355 success: 5 fail: 55
key: 1609939365 success: 5 fail: 185
key: 1609939366 success: 0 fail: 1054
key: 1609939367 success: 0 fail: 2082
key: 1609939368 success: 0 fail: 1291
key: 1609939369 success: 0 fail: 876
key: 1609939370 success: 0 fail: 1070
key: 1609939371 success: 0 fail: 1825
key: 1609939372 success: 0 fail: 1412
key: 1609939373 success: 0 fail: 1535
key: 1609939374 success: 0 fail: 1874
key: 1609939375 success: 5 fail: 1500
key: 1609939376 success: 0 fail: 1650
key: 1609939377 success: 0 fail: 787
key: 1609939378 success: 0 fail: 620
key: 1609939379 success: 0 fail: 259
--- PASS: TestRollingWindow_IsSuccess (24.10s)
PASS

Process finished with exit code 0
```