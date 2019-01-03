/*
.匹配除换行符以外的任意字符
^字符串的开始
$字符串的结束
*/
package main

import (
	"fmt"
	"regexp"
)
func main(){
	// isok,str:=regexp.MatchString("","helio1")
	// fmt.Println(isok,str)
	// reg:=regexp.MustCompile("[a-zA-Z]{3}")
	// str:=reg.FindAllString("heloo",2)
	// fmt.Println(str)
	/*
^$分别表示匹配的开始和结束，界定我们正则表达式的范围。
[\d]{4}表示我们要正好匹配4位数字，因为年份是4位，所以我们定义为匹配4位。后面的月份和天是2位，所以定义为2位。
[\w-]匹配字符串和中杠，加号(+)表示匹配1个或者多个。
然后他们都加了括号()，意味着我们要提取这些字符串。
	*/
	// flysnowRegexp :=
	//  regexp.MustCompile(`^http://www.flysnow.org/([\d]{4})/([\d]{2})/([\d]{2})/([\w-]+).html$`)
	// params :=
	//  flysnowRegexp.FindStringSubmatch("http://www.flysnow.org/2018/01/20/golang-goquery-examples-selector.html")

    // for _,param :=range params {
	// fmt.Println(param)
	  Rege:=regexp.MustCompile(`河[\p{Han}]*大[\p{Han}\w]+路`)
	  params:=Rege.FindAllStringSubmatch("我们河南大学文化路都是文化人,当时河大学林路要好好学习文化知识",2)
	  fmt.Println(params)
	  fmt.Println("学习"=="学习好")
	 }
	