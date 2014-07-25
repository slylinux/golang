package main
import "fmt"

func InsertSort(values []int){
	if len(values)<2{
		fmt.Println("too short to sort!")
		return
	}
	if values[0]>values[1]{
		values[0],values[1]=values[1],values[0]
	}
	for i:=2;i<len(values);i++{
		if values[i-1]>values[i]{			
			for j:=i-1;j>=0;j--{
				if(values[j]>values[j+1]){
					values[j],values[j+1]=values[j+1],values[j]
				}else {
					break
				}				
			}
		}
	}
}

func main(){
	values :=[]int{7,2,5,4,10,9,17,8,1,3,12}
	fmt.Println("before sort:",values)
	InsertSort(values)
	fmt.Println("after InsertSort:",values)

	values =[]int{8,2,5,9,1,3,12}
	fmt.Println("before sort:",values)
	InsertSort(values)
	fmt.Println("after InsertSort:",values)

	values =[]int{2,7,10,3,1,89,76,4,5,100,300,230,200,198,17}
	fmt.Println("before sort:",values)
	InsertSort(values)
	fmt.Println("after InsertSort:",values)



}
