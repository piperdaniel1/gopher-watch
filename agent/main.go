package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/mem"
)

func main() {
	fmt.Println("This is the agent package.")

	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(v)
}
