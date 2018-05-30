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
	rows, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Error: The second argument must be an integer value for the number of rows in the disk")
		return
	}
	nibbles, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Error: The third argument must be an integer value for the number of nibbles in each row")
		return
	}

	d := disk.BuildFormattedDisk(nibbles, rows, name)
	disk.PrintDisk(d)

}
