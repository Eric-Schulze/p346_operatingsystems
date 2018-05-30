package disk_util

import (
	"fmt"
	"strconv"
	//"io/ioutil"
	//"os"
)

func Format(args []string) {
	if len(args) != 3 {
		fmt.Println("Error: You must enter three arguments - name, rows, and nibbles")
		return
	}

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

	name := args[0]
	if len(name) > ((nibbles - 10) / 2) {
		fmt.Println("Error: The name must be less than " + string((nibbles-10)/2) + " characters")
		return
	}

	d := BuildFormattedDisk(nibbles, rows, name)
	WriteDisk(d)
}

func Directory() {
	d := ReadDisk()
	var fileArray []string

	nextFileHeaderPtr := d.Root.ThirdPointer
	nextFileHeaderId := convertClusterPointerToClusterId(nextFileHeaderPtr)

	for nextFileHeaderId != 0 {
		fileArray = append(fileArray, GetFileClusterData(nextFileHeaderId, d))
		nextFileHeaderPtr = d.Data[nextFileHeaderId].FirstPointer
		nextFileHeaderId = convertClusterPointerToClusterId(nextFileHeaderPtr)
	}

	for _, s := range fileArray {
		fmt.Print(s + "\t\t")
	}
	fmt.Println()
}

func Version() {
	fmt.Println("Shell Program\nBy Eric Schulze\nVersion 1.0.03")
}

func Help() {
	fmt.Println("Help Manual for Shell Program\n")
	fmt.Println("Operation\t\tCommand\t\t\t\tDescription")
	fmt.Println("-----------------------------------------------------------------------------------------------------\n")
	fmt.Println("Version\t\t\t-v\t\t\t\tDisplay Information about this Shell Program, including author and version number\n")
	fmt.Println("Help\t\t\t-h or -help\t\t\tDisplay help information and listing of options for the shell\n")
	fmt.Println("Directory\t\t-dir\t\t\t\tList the directory of files in the disk\n")
	fmt.Println("Format Disk\t\t-format name rows nibbles\tFormat the current shell program disk to a blank disk with a new name")
	fmt.Println("--Options\t\tname\t\t\t\t(required) New name for the disk")
	fmt.Println("\t\t\trows\t\t\t\t(required) Number of rows in the new disk")
	fmt.Println("\t\t\tnibbles\t\t\t\t(required) Number of nibbles in each row of the new disk\n")
	fmt.Println("Rename\t\t\t-rename name\t\t\tUpdate the current name for the disk")
	fmt.Println("--Options\t\tname\t\t\t\t(required) New name for the disk\n")
	fmt.Println("Raw Data\t\t-raw\t\t\t\tDisplay the raw data of the current disk\n")
	fmt.Println("Add File\t\t-add fileName fileData\t\tAdd a new file to the disk with the associated name and data")
	fmt.Println("--Options\t\tfileName\t\t\t(required) Name for the new file")
	fmt.Println("\t\t\tfileData\t\t\t(optional) Text data associated with the file\n")
	fmt.Println("Disk Usage\t\t-du\t\t\t\tDisplay a report for current disk usage, bad and used clusters, and available disk space\n")
	fmt.Println("File Size\t\t-filesize fileName\t\tReport on the current nibble size of the specified file data")
	fmt.Println("--Options\t\tfileName\t\t\t(required) Name for the new file\n")
	fmt.Println("Bad Cluster\t\t-bad\t\t\t\tCreate a bad cluster on the disk; used exclusively for function testing\n")
	fmt.Println()
}

func Rename(args []string) {
	if len(args) != 1 {
		fmt.Println("Error: You must enter one argument: disk name")
		return
	}

	d := ReadDisk()
	nibbles := len(d.Root.Data)
	name := args[0]
	if len(name) > ((nibbles - 10) / 2) {
		fmt.Println("Error: The name must be less than " + string((nibbles-10)/2) + " characters")
		return
	}
	RenameDisk(d, name)
}

func AddFile(args []string) {
	if len(args) == 0 {
		fmt.Println("Error: You must enter at least one argument: file name")
		return
	}

	d := ReadDisk()
	id := GetFileHeaderClusterIdFromFileName(args[0], d)
	if id != 0 {
		fmt.Println("Error: File Name already exists.")
		return
	}

	updatedDisk, newFileId, err := NewFileHeader(args[0], d)
	if newFileId == 0 {
		fmt.Println(err)
		return
	}
	check(err)

	if len(args) == 2 {
		updatedDisk, err = NewFileData(args[1], newFileId, updatedDisk)
		check(err)
	}

	WriteDisk(updatedDisk)
	fmt.Println("File Successfully Added")
}

func FileType(args []string) {
	if len(args) == 0 {
		fmt.Println("Error: You must enter at least one argument: file name")
		return
	}

	d := ReadDisk()
	id := GetFileHeaderClusterIdFromFileName(args[0], d)
	if id == 0 {
		fmt.Println("Error: File Name does not exists.")
		return
	}

	fileString, err := ReadFileData(id, d)
	check(err)
	if fileString == "" {
		fmt.Println("Error: File is empty")
		return
	}

	fmt.Println(fileString)
}

func BadCluster() {
	d := ReadDisk()
	d = CreateBadCluster(d)
	WriteDisk(d)
}

func DiskUsage() {
	d := ReadDisk()
	var availableClusters, badClusters, usedClusters, totalClusters int
	var clusterType int64
	var err error
	var src string

	for _, cluster := range d.Data {
		src = string(cluster.ClusterType)

		clusterType, err = strconv.ParseInt(src, 0, 64)
		check(err)

		switch clusterType {
		case 0:
			usedClusters++
		case 1:
			availableClusters++
		case 2:
			badClusters++
		case 3:
			usedClusters++
		case 4:
			usedClusters++
		}

		totalClusters++
	}

	fmt.Printf("***** Disk Usage *****\n\n")
	fmt.Printf("State\t\t\tCount\t\t\tPercent\n")
	fmt.Printf("---------------------------------------\n")
	fmt.Printf("Used\t\t\t%v\t\t\t%.2f%%\n", usedClusters, (float32(usedClusters) / float32(totalClusters) * 100))
	fmt.Printf("Available\t\t%v\t\t\t%.2f%%\n", availableClusters, (float32(availableClusters)/float32(totalClusters))*100)
	fmt.Printf("Bad\t\t\t%v\t\t\t%.2f%%\n\n", badClusters, (float32(badClusters)/float32(totalClusters))*100)
	fmt.Printf("Total Number of Clusters: %v\n", totalClusters)
	fmt.Printf("Total Number of Unavailable Clusters: %v\n\n", (usedClusters + badClusters))
	if totalClusters <= (usedClusters + badClusters) {
		fmt.Printf("*** DISK FULL ***\n\n")
	} else {
		fmt.Printf("*** %.2f%% DISK AVAILABLE ***\n\n", (float32(availableClusters)/float32(totalClusters))*100)
	}
}

func FileSize(args []string) {
	if len(args) == 0 {
		fmt.Println("Error: You must enter at least one argument: file name")
		return
	}

	d := ReadDisk()
	id := GetFileHeaderClusterIdFromFileName(args[0], d)
	if id == 0 {
		fmt.Println("Error: File Name does not exists.")
		return
	}

	fileString, err := ReadFileData(id, d)
	check(err)
	if fileString == "" {
		fmt.Println("Error: File is empty")
		return
	}

	fmt.Printf("File Size: %v nibbles\n", len(fileString)*2)
}
