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
