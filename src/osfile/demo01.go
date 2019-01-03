package main

import (
	"io"
	"bufio"
	"fmt"
	"os"
)
func main(){
	//返回file这个结构体的指针
	file,err:=os.Open("d:/test.txt")
	if err!=nil{
		fmt.Println("err:",err)
	}
	//当函数退出是,一定要记得关闭
	defer file.Close()  //及时关闭文件句柄,一面内存泄漏
	//创建一个*Reader 带缓冲的
	reader:=bufio.NewReader(file)
	//循环的读取文件的内容
	for{
	str,err:=reader.ReadString('\n') //读到换行就结束
	if (err==io.EOF){//标示读到文件的末尾
		break;
	}
	fmt.Print(str)
	}
	fmt.Println("文件读取结束")
} 