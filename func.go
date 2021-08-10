package gTools

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func GetNowTimeStr() string{
	return time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
}

func GetLocalIP() string{
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil{
		log.Printf("get local addr err:%v", err)
		return ""
	}
	localIp := strings.Split(conn.LocalAddr().String(), ":")[0]
	return localIp
}

func GetHostName() string{
	name, _ := os.Hostname()
	return "v1.0.2" + name
}

func main(){
	fmt.Println(GetNowTimeStr())
	fmt.Println(GetLocalIP())
	fmt.Println(GetHostName())
}