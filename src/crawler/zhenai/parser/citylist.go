package parser
import(
	"regexp"
	"crawler/engine"
)
const citylistRe=`<a href="(http://city.zhenai.com/[a-z0-9]+)"[^>]*>([^<]+)</a>`
func ProcessCity(mess []byte) (engine.ParserRusult){
	re,err:= regexp.Compile(citylistRe)
	checkerr(err)
	match:=re.FindAllSubmatch(mess,-1)
	result:=engine.ParserRusult{}
	for _,x:=range match{
		result.Items=append(result.Items,"City" + string(x[2]))
		result.Requests=append(result.Requests,engine.Request{Url:string(x[1]),ParserFunc:ParserCity})
	}
	return result
 }
 func checkerr(e error){
	if e!=nil{
		panic(e)
	}
 } 