package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"sort"
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

func main() {

	rw := &rollingWindow{
		List:       list.New(),
		TimeSecond: 10,
		Size:       5,
		Census:     map[int64]*census{},
	}

	for i := 0; i < 60; i++ {
		if rw.IsSuccess() {
			// TODO
		}
	}

	for i, c := range rw.Census {
		fmt.Println("key:", i, "success:", c.Success, "fail:", c.Fail)
	}

	time.Sleep(10 * time.Second)
	fmt.Println("<<================================>>")

	wg := sync.WaitGroup{}

	for i := 0; i < 30; i++ {
		rw.IsSuccess()
		wg.Add(600)
		for i := 0; i < 600; i++ {
			go func() {
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(1500)))
				wg.Done()
				rw.IsSuccess()

			}()
		}

		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	}

	wg.Wait()
	type kv struct {
		Key   int64
		Value *census
	}
	var ss []kv
	for k, v := range rw.Census {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		//return ss[i].Key > ss[j].Key  // 降序
		return ss[i].Key < ss[j].Key // 升序
	})
	for _, c := range ss {
		fmt.Println("key:", c.Key, "success:", c.Value.Success, "fail:", c.Value.Fail)
	}
}
