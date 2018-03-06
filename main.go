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
	build := 15
	fmt.Println("goportscan build", build, " Matt Weidner @jopawp")
	fmt.Println("   https://github.com/mattweidner/goportscan")
	flag.StringVar(&target, "t", "", "target hostname or IP")
	flag.Parse()
	fmt.Println("")
	if target == "" {
		fmt.Println("Please specify a target with the -t flag.")
		fmt.Println("  ex: goportscan -t testfire.net")
		fmt.Println("      goportscan -t 192.168.0.1")
		return
	}
	fmt.Println("Scanning", target, " ports 1-65535")
	wg.Add(65535)
	for p := 1; p < 65536; p++ {
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
