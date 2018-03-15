package main

import (
	"fmt"
	"os"
	"strconv"
	disk "utils/disk_util"
)

func main() {
	//Retrieve command line arguments
	args := os.Args[1:]
	if len(args) != 3 {
		fmt.Println("Error: You must enter three arguments - name, rows, and nibbles")
		return
	}

	name := args[0]
	rows, _ := strconv.Atoi(args[1])
	nibbles, _ := strconv.Atoi(args[2])

	d := disk.BuildFormattedDisk(nibbles, rows, name)
	disk.PrintDisk(d)

}
