package engine
import(
	"crawler/fetcher"
	"log"
)
type SampleEngine struct {

}
func (e SampleEngine) Run(seeds ...Request)  {
	var Requests []Request
	for _ , r :=range seeds{
		Requests=append(Requests,r)
	}
	for len(Requests)>0{
		r:=Requests[0]
		Requests=Requests[1:]
		ParseRult,err:=worker(r)
		if(err!=nil){
          continue
		}
		//jiang所有东西展开送进去
		Requests=append(Requests,ParseRult.Requests...)
		for _ ,item:=range ParseRult.Items{
           log.Printf("Got item %v",item)
		}
	}
}
func  worker (r Request) (ParserRusult,error ){
	body,err:=fetcher.Fetch(r.Url)
	if err!=nil{
		log.Printf("Feath error"+"fetcher url %s:%v",r.Url,err)
		return ParserRusult{},err
	}
	ParseRult:=r.ParserFunc(body)
	return ParseRult,nil
}