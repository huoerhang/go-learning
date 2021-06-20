package main

import "fmt"

func main()  {
	months:=[]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	spring:=months[0:3]
	summer:=months[3:6]
	autumn:=months[6:9]
	winter:=months[9:12]
	fmt.Println(spring,len(spring),cap(spring))
	fmt.Println(summer,len(summer),cap(summer))
	fmt.Println(autumn,len(autumn),cap(autumn))
	fmt.Println(winter,len(winter),cap(winter))
}
