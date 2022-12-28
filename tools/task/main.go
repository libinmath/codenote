package main

import (
	"log"
	"os"
	"time"
)

/* 调用示例:*/
func main() {
	log.Println("...开始执行任务...")

	timeout := 20 * time.Second // 任务执行允许的总时长，防止hang
	r := New(timeout)

	r.Add(createTask(1), createTask(2), createTask(5))

	if err := r.Start(); err != nil {
		switch err {
		case ErrTimeOut:
			log.Println(err)
			os.Exit(1) //退出
		case ErrInterruput:
			log.Println(err)
			os.Exit(2)
		default:
			break
		}
	}
	log.Println("...任务执行结束...")
}

func createTask(param int) func() {
	return func() {
		log.Printf("正在执行任务%d", param)
		time.Sleep(time.Duration(param) * time.Second)
	}
}
