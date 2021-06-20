package main

import "fmt"

func main()  {
	letter:=[]string{"A","B","C","D","E"}
	remove:=2

	if remove<len(letter){
		fmt.Println("Before",letter,"Remove",letter[remove])
		//1.先创建一个切片，从0到2
		//2.将2到最后一个元素追加到第一步创建的分片
		letter=append(letter[:remove],letter[remove+1:]...)
		fmt.Println("After",letter)
	}
}
