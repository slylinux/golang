package main
import (
	"fmt"
	"net"
	"os"
	"unsafe"
	"encoding/binary"
	"time"
)
type Music struct{
	Id uint32
	Name []byte
	Type []byte
}

func checkConnection(conn net.Conn,err error){
	if(err!=nil){
		return
	}
	fmt.Println("connected with ",conn.RemoteAddr().String())
}
func reconnect() (net.Conn,error){
	conn,err:=net.Dial("tcp","127.0.0.1:8032")
	return conn,err
}

func main(){
	args:=os.Args
	if len(args)!=2{
		println("Usage: ./dialup string")
		return
	}
	data:=make([]byte,1024)
	conn,err:=net.Dial("tcp","127.0.0.1:8032")
	checkConnection(conn,err)
	read:=true
	for read{
		if err!=nil {
			println("socket error,reconnect");
			time.Sleep(1e9);	
			conn,err=reconnect()
			checkConnection(conn,err)
			if err!=nil{
				continue
			}
		}

		_,err=conn.Write([]byte(args[1]))
		if err!=nil{
			continue
		}

		count,err:=conn.Read(data)
		if err!=nil{
			continue
		}
		
		var m Music
		m.Name=make([]byte,128)
		m.Type=make([]byte,128)
		//fmt.Println("sizeof Music",unsafe.Sizeof(m))
		//fmt.Println("offset of id,name,type",unsafe.Offsetof(m.Id),unsafe.Offsetof(m.Name),unsafe.Offsetof(m.Type))
		//fmt.Println("length of id,name,type",unsafe.Sizeof(m.Id),len(m.Name),len(m.Type))
		nameIndex:=int(unsafe.Sizeof(m.Id))
		typeIndex:=nameIndex+len(m.Name)
		m=Music{uint32(binary.LittleEndian.Uint32(data[0:nameIndex])),
			data[nameIndex:typeIndex],data[typeIndex:count]}
		fmt.Printf("server syas:id:%v,name:%s,type:%s\n",uint32(m.Id),string(m.Name),string(m.Type))
		time.Sleep(1e9);
	}
	conn.Close()
}




