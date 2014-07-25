package main

import (
	"fmt"
	"sort"
)

var (
	barVal=map[string]int{"alpha":34,"beta":56,"charlie": 23,"delta": 87, "echo": 56, 
		                  "foxtrot": 12, "golf": 34, "hotel": 16,
		                  "indio": 87, "juliet": 65, "kilo": 43, "lima": 98}
)

func main(){
	s:="the string "
	print(s)
	fmt.Println("unsorted:",barVal)
	keys:=make([]string,len(barVal))
	i:=0
	for k,_:=range barVal{
		keys[i]=k
		i++
	}
	sort.Strings(keys)
	fmt.Println("sourted:")
	for _,k:=range keys{
		fmt.Println(k,barVal[k])
	}
	var invMap=make(map[int]string,len(barVal))
	for k,v:=range barVal{
		invMap[v]=k
	}
	fmt.Println("inverted:",invMap)
	values:=make([]int,len(invMap))
	i=0
	for k,_:=range invMap{
		values[i]=k
		i++
	}
	sort.Ints(values)
	fmt.Println("sorted:")
	for _,k:=range values{
		fmt.Println(k,invMap[k])
	}
}
