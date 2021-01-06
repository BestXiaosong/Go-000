package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"testing"
	"time"
)

func TestRollingWindow_IsSuccess(t *testing.T) {
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
