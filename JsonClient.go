package main
import (
	"fmt"
	"net"
	"encoding/json"
)

type Person struct{
	Name string
	Id   int
	Lks []string
}
type Book struct {
	Title string
	Authors []string
	Publisher string
	IsPublished bool
	Price float32
}

func main(){

	buf:=make([]byte,512)
	conn,err:=net.Dial("tcp","127.0.0.1:50000")
	if err!=nil{
		fmt.Println("Dial Error:",err.Error())
		return
	}
	person := Person{"sly",2008,[]string{"hello","world","I am ","sly!"}}
	fmt.Println("original values:",person)
	buf,err=json.Marshal(person)
	if err!=nil{
		fmt.Println("Marshal Error:",err.Error())
		return
	}
	var person1 Person
	err=json.Unmarshal(buf,&person1)
	if err!=nil{
		fmt.Println("Unmarshal Error:",err.Error());
		return
	}
	fmt.Println("after Unmarshal Person1:",person1);
	fmt.Println("data write to server:",string(buf));
	conn.Write(buf)
	conn.Close()
}










