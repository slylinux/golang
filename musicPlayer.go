package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"bufio"
)

type Music struct{
	Id string
	Name string
	Type string
}

type MusicManager struct{
	musics []Music
}

func (m *MusicManager)List(){
	for _,m := range m.musics{
		fmt.Println(m.Id,m.Name)
	}
}
func (m *MusicManager)Add(music *Music){
	m.musics=append(m.musics,*music)
}
func (m *MusicManager)Remove(index int) *Music{
	if index <0 || index>len(m.musics)-1{
		fmt.Println("Remove index out of range.")
		return nil
	}
	removedMusic:=&m.musics[index]
	m.musics=append(m.musics[:index],m.musics[index+1:]...)
	return removedMusic
}
func NewMusicManager() *MusicManager{
	return &MusicManager{make([]Music,0)}
}

var musicManager *MusicManager
var id int
func handleCommands(tokens []string){
	switch tokens[0]{
	case "add":
		if len(tokens)<3 {
			println("you must specify the name,id,type for a music.")
			break;
		}
		musicManager.Add(&Music{strconv.Itoa(id),tokens[1],tokens[2]})
		id++
	case "remove":
		if len(tokens)<2 {
			println("please give the index for the music.");
			break;
		}
		index,_:=strconv.Atoi(tokens[1])
		musicManager.Remove(index)
	case "list":
		musicManager.List()
	default:
         println("unsupported command.");
	}
}
func main(){
	id=0
	fmt.Println("Usage:\nlist------list all the musics in the library.\n"+
		"add Name type------add a music\n"+"remove index------remove the music at index")
	musicManager=NewMusicManager()
	reader:=bufio.NewReader(os.Stdin)
	for{
		fmt.Print("Enter Commands->")
		l,_,_:=reader.ReadLine()
		line:=string(l)
		if line=="q"||line=="exit"{
			break
		}
		tokens:=strings.Split(line," ")
		handleCommands(tokens)
	}
	
}