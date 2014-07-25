package main
import "fmt"

type Shaper interface{
	Area() float32
}

type Square struct{
	side float32
}

func (sq *Square) Area() float32{
	return sq.side*sq.side
}

type Rectangle struct{
	length,width float32
}
func( r Rectangle) Area() float32{
	return r.length*r.width
}

func main(){
	sq1:=new(Square)
	sq1.side=5
	rct:=Rectangle{4,7}	
	var areaIntf Shaper
	areaIntf=sq1
	fmt.Printf("Area of Square is %f\n",areaIntf.Area())
	areaIntf=rct
	fmt.Printf("Area of Rectange is %f\n",areaIntf.Area())
}




















