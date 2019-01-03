package parser

import (
	"crawler/engine"
	"regexp"
	"crawler/model"
	"strconv"
)
   var weightre=regexp.MustCompile(`<span class="label">体重：</span><span field="">([^<]+)</span>`)
   var genderre=regexp.MustCompile(`<span class="label">性别：</span><span field="">([^<]+)</span>`)
   var agere=regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
   var heightre=regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)
   var incomere=regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
   var marrigere=regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
   var edncationre=regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
   var hokoure=regexp.MustCompile(`<td><span class="label">工作地：</span>([^<]+)</td>`)
   var occupationre=regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
   //var xinzuore=regexp.MustCompile(`<td><span class="label">星座：</span>([^<]+)</td>`)
   var carre=regexp.MustCompile(`<span class="label">是否购车：</span><span field="">([^<]+)</span>`)
   var housere=regexp.MustCompile(`<span class="label">住房条件：</span><span field="">([^<]+)</span>`)
func ParserProfile(contains []byte,name string) engine.ParserRusult {
     profile:= model.Profile{}
     profile.Name=name
     age,err:=strconv.Atoi(extrastring(contains,agere))
    if err==nil{
  	   profile.Age=age
    }
	height,err:=strconv.Atoi(extrastring(contains,heightre))
	if err==nil{
		profile.Height=height
	}
	weiht,err:=strconv.Atoi(extrastring(contains,weightre))
	if err!=nil{
		profile.Weight=weiht
	}
    profile.Marriage=extrastring(contains,marrigere)
    profile.Gender=extrastring(contains,genderre)
	profile.House=extrastring(contains,heightre)
	profile.Car=extrastring(contains,carre)
	profile.Occupation=extrastring(contains,occupationre)
	profile.Education=extrastring(contains,edncationre)
	profile.Hokou=extrastring(contains,hokoure)
	result:=engine.ParserRusult{Items:[]interface{}{profile}}
	return result
}
func extrastring (contents []byte ,re *regexp.Regexp) string{
   match:=re.FindSubmatch(contents)
   if len(match)>=2{
   	return string(match[1])
   }else {
   	return ""
   }
}
