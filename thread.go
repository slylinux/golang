package main

import (
	"fmt"
	"sync"
	"runtime"
)

var counter int=0

func count(lock *sync.Mutex){
	lock.Lock()
	counter++
	fmt.Println(counter)
	lock.Unlock()
}
func Count(ch chan int){
	fmt.Println("counting")
	ch<-1
}
func main(){
	lock:=&sync.Mutex{}
	for i:=0;i<10;i++{ go count(lock) }
	for{
		lock.Lock()
		c:=counter
		lock.Unlock()
		runtime.Gosched()
		if(c>=10){
			break
		}
	}
//another way
	chs:=make([]chan int,10)
	for i:=0;i<10;i++{
		chs[i]=make(chan int)
		go Count(chs[i])
	}
	for _,ch:=range (chs){
		<-ch
	}
}
