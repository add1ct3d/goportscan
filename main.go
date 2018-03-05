package main

// File: goportscan.go
// Date: 2018-03-04
// Author: Matt Weidner <matt.weidner@gmail.com>
// Description: Concurrent full port scanner
import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var target string
	build := 10
	flag.StringVar(&target, "t", "127.0.0.1", "target hostname or IP")
	flag.Parse()
	fmt.Println("goportscan build", build)
	fmt.Println("Scanning", target)
	wg.Add(65536)
	for p := 0; p < 65536; p++ {
		port := fmt.Sprintf("%d", p)
		go func(port string) {
			defer wg.Done()
			//fmt.Println("Trying", port)
			c, e := net.DialTimeout("tcp", target+":"+port, time.Second*3)
			if e == nil {
				fmt.Println(port + "/open")
				c.Close()
			}
		}(port)
		if p%5 == 0 {
			time.Sleep(2 * time.Millisecond)
		}
	}
	wg.Wait()
}
