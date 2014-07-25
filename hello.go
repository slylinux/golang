// hello
package main

import (
	"fmt"
	"flag"
	"runtime"
	"time"
)

func main() {
	flag.Parse()	
	for i:=0;i<flag.NArg();i++{
		fmt.Println("arg[",i,"]",flag.Arg(i))
	}
	flag.Parse();
	hi:="hello world!\n"
	print(hi)
	fmt.Println("num of cpus =",runtime.NumCPU());
	fmt.Println("now time=",time.Now());
//	for{
//		t0:=time.Now();
//		time.Sleep(250000);
//		t1:=time.Now();
//		fmt.Println("timer interval=",t1.Sub(t0));
//	}
}



















