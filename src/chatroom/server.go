package main
import (
	"fmt"
	"os"
	"net"
)
func checkerror(err error){
	if err!=nil{
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
func processInfo(conn net.Conn){
  buffer:=make([]byte, 1024)
  defer conn.Close()
  for{
	  _,err:=conn.Read(buffer)
	  checkerror(err)
	  fmt.Println(conn)
	  fmt.Println("has reviced"+string(buffer))
  }
}
func main(){
	listen_socket,err:=net.Listen("tcp","127.0.0.1:9090")
	if err!=nil{
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer listen_socket.Close()
	for{
		conn, err:= listen_socket.Accept()
		checkerror(err)
		go processInfo(conn)
	}
}