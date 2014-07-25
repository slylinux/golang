package main

import ("fmt"
	"time"
	"log"
	"os"
)

var logfile os.File
var logger *log.Logger

func add(x,y int){
	z:=x+y	
	logger.Println("z=",z)
}

func main(){
	logfile,err:=os.OpenFile("file.log",os.O_RDWR|os.O_CREATE,0)
	defer logfile.Close()
	logger=log.New(logfile, "\r\n", log.Ldate|log.Ltime|log.Llongfile)
	logger.Println("main app start!")

	if err!=nil{
		println("log file create error");
		return
	}
	t0:=time.Now()
	for i:=0;i<1000;i++{
		time.Sleep(100);
		go add(i,i+2)
}
	t1:=time.Now()
	time.Sleep(10e9);
	fmt.Println("time spent=",t1.Sub(t0));
}














