package parser

import (
	"crawler/engine"
	"regexp"
)
const cityre=`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
func ParserCity(contents []byte) engine.ParserRusult  {
	re,err:= regexp.Compile(cityre)
	checkerr(err)
	match:=re.FindAllSubmatch(contents,-1)
	result:=engine.ParserRusult{}
	for _,x:=range match{
		name:=string(x[2])
		result.Items=append(result.Items,"Users" + string(x[2]))
		result.Requests=append(result.Requests,engine.Request{Url:string(x[1]),ParserFunc: func(bytes []byte) engine.ParserRusult {
			return ParserProfile(bytes,name)
		}})
	}
	return result
}
