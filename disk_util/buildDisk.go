package disk_util

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
)

type Cluster struct {
	ClusterId     int
	Data          []rune
	ClusterType   rune
	FirstPointer  []rune
	SecondPointer []rune
	ThirdPointer  []rune
}

type Disk struct {
	Data []Cluster
	Name string
	Root Cluster
}

//Build Disk
func buildRootCluster(rowSize int, name string, availClusterPointer, badClusterPointer, fileClusterPointer []rune) Cluster {
	c := Cluster{Data: make([]rune, rowSize), ClusterType: rune('0'), ClusterId: 0}
	for i, _ := range c.Data {
		c.Data[i] = rune('0')
	}

	//pointers
	c.Data[1] = availClusterPointer[0]
	c.Data[2] = availClusterPointer[1]
	c.FirstPointer = c.Data[1:3]
	c.Data[3] = badClusterPointer[0]
	c.Data[4] = badClusterPointer[1]
	c.SecondPointer = c.Data[3:5]
	c.Data[5] = fileClusterPointer[0]
	c.Data[6] = fileClusterPointer[1]
	c.ThirdPointer = c.Data[5:7]

	//disk_name
	name_hex := hex.EncodeToString([]byte(name))
	for i, letter := range name_hex {
		c.Data[i+7] = rune(letter)
	}
	return c
}

func BuildNewCluster(size, id int, designatedType rune, first []rune) Cluster {
	c := Cluster{Data: make([]rune, size), ClusterId: id, ClusterType: designatedType, FirstPointer: first}
	for i, _ := range c.Data {
		c.Data[i] = rune('0')
	}
	c.Data[0] = c.ClusterType
	c.Data[1] = c.FirstPointer[0]
	c.Data[2] = c.FirstPointer[1]

	return c
}

func BuildFormattedDisk(rowSize, diskSize int, diskName string) Disk {
	availPointer := []rune{'0', '1'}
	badPointer := []rune{'0', '0'}
	filePointer := []rune{'0', '0'}
	d := Disk{Data: make([]Cluster, diskSize), Name: diskName}
	d.Root = buildRootCluster(rowSize, diskName, availPointer, badPointer, filePointer)
	d.Data[0] = d.Root

	var s string
	for i := 1; i < diskSize; i++ {
		if i < (diskSize - 1) {
			s = fmt.Sprintf("%0*x", 2, (i + 1))
		} else {
			s = "00"
		}
		d.Data[i] = BuildNewCluster(rowSize, i, '1', []rune(s))
	}

	return d
}

//Disk Operations
func RenameDisk(d Disk, name string) {
	name_hex := hex.EncodeToString([]byte(name))
	for i, letter := range name_hex {
		d.Root.Data[i+7] = rune(letter)
	}

	fmt.Println(name)
	WriteDisk(d)
}

func NewFileHeader(fileName string, d Disk) (newDisk Disk, newFileHeaderId int, err error) {
	if (len(fileName) * 2) > (len(d.Root.Data) - 7) {
		newDisk = d
		err_string := fmt.Sprintf("Error: File name is too long. Name can only be %v characters long.", (len(d.Root.Data)-7)/2)
		err = errors.New(err_string)
		newFileHeaderId = 0
		return
	}

	//Get Next Available Cluster and Update Root Pointer
	availableFileHeaderPtr := d.Root.FirstPointer
	availableFileHeaderId := convertClusterPointerToClusterId(availableFileHeaderPtr)
	d.Root.Data[1] = d.Data[availableFileHeaderId].FirstPointer[0]
	d.Root.Data[2] = d.Data[availableFileHeaderId].FirstPointer[1]
	d.Root.FirstPointer = d.Root.Data[1:3]
	d.Data[0] = d.Root

	//Get last file header in linked list
	nextFileHeaderPtr := d.Root.ThirdPointer
	nextFileHeaderId := convertClusterPointerToClusterId(nextFileHeaderPtr)
	lastFileHeaderId := nextFileHeaderId

	for nextFileHeaderId != 0 {
		lastFileHeaderId = nextFileHeaderId
		nextFileHeaderPtr = d.Data[nextFileHeaderId].FirstPointer
		nextFileHeaderId = convertClusterPointerToClusterId(nextFileHeaderPtr)
	}

	//Add new file header to link list
	if lastFileHeaderId == 0 {
		d.Root.Data[5] = availableFileHeaderPtr[0]
		d.Root.Data[6] = availableFileHeaderPtr[1]
		d.Root.ThirdPointer = d.Root.Data[5:7]
		d.Data[0] = d.Root
	} else {
		d.Data[lastFileHeaderId].Data[1] = availableFileHeaderPtr[0]
		d.Data[lastFileHeaderId].Data[2] = availableFileHeaderPtr[1]
		d.Data[lastFileHeaderId].FirstPointer = d.Data[lastFileHeaderId].Data[1:3]
	}

	//Input data into available cluster
	var fileHeader = d.Data[availableFileHeaderId]
	fileHeader.Data[0] = rune('3')
	d.Data[availableFileHeaderId].ClusterType = fileHeader.Data[0]
	fileHeader.Data[1] = rune('0')
	fileHeader.Data[2] = rune('0')
	d.Data[availableFileHeaderId].FirstPointer = fileHeader.Data[1:3]
	fileHeader.Data[3] = rune('0')
	fileHeader.Data[4] = rune('0')
	d.Data[availableFileHeaderId].SecondPointer = fileHeader.Data[3:5]

	name_hex := hex.EncodeToString([]byte(fileName))
	for i, letter := range name_hex {
		fileHeader.Data[i+5] = rune(letter)
	}

	newDisk = d
	err = nil
	newFileHeaderId = availableFileHeaderId
	return
}

func NewFileData(fileString string, fileHeaderId int, d Disk) (updatedDisk Disk, err error) {
	//Input data into available cluster, over flowing if necessary
	var availableLength = len(d.Root.Data) - 5
	var fileData = []byte(fileString)
	var dataLength = len(fileData) * 2
	var currentFileClusterId = fileHeaderId
	var currentFileClusterPtr []rune
	var referringClusterId int
	var dataEndPoint = 0
	var dataStartPoint = 0
	var fileData_partial []byte

	for i := 0; dataStartPoint < (dataLength / 2); i++ {
		referringClusterId = currentFileClusterId
		currentFileClusterId, currentFileClusterPtr, d = getNextAvailableCluster_UpdateRoot(d)

		if referringClusterId == fileHeaderId {
			d.Data[referringClusterId].Data[3] = currentFileClusterPtr[0]
			d.Data[referringClusterId].Data[4] = currentFileClusterPtr[1]
			d.Data[referringClusterId].SecondPointer = d.Data[referringClusterId].Data[3:5]
		} else {
			d.Data[referringClusterId].Data[1] = currentFileClusterPtr[0]
			d.Data[referringClusterId].Data[2] = currentFileClusterPtr[1]
			d.Data[referringClusterId].FirstPointer = d.Data[referringClusterId].Data[1:3]
		}

		d.Data[currentFileClusterId].Data[0] = rune('4')
		d.Data[currentFileClusterId].ClusterType = d.Data[currentFileClusterId].Data[0]
		d.Data[currentFileClusterId].Data[1] = rune('0')
		d.Data[currentFileClusterId].Data[2] = rune('0')
		d.Data[currentFileClusterId].FirstPointer = d.Data[currentFileClusterId].Data[1:3]

		if ((dataLength / 2) - dataStartPoint) <= availableLength/2 {
			dataEndPoint = dataLength / 2
		} else {
			dataEndPoint = dataStartPoint + (availableLength / 2)
		}

		dataEndPoint = dataEndPoint
		fileData_partial = fileData[dataStartPoint:dataEndPoint]
		file_hex := hex.EncodeToString(fileData_partial)
		for i, letter := range file_hex {
			d.Data[currentFileClusterId].Data[i+3] = rune(letter)
		}

		dataStartPoint = dataEndPoint
	}

	updatedDisk = d
	err = nil
	return
}

func AppendToFileData(fileData string, fileHeaderId int, d Disk) (updatedDisk Disk, err error) {
	return
}

func ReadFileData(fileHeaderId int, d Disk) (fileData string, err error) {
	nextClusterId := convertClusterPointerToClusterId(d.Data[fileHeaderId].SecondPointer)
	var marker1, marker2, string_hex string
	var int1, int2 int

	for nextClusterId != 0 {
		marker1 = string(d.Data[nextClusterId].Data[3])
		marker2 = string(d.Data[nextClusterId].Data[4])
		string_hex = ""
		for i := 1; marker1 != "0" || marker2 != "0"; i++ {
			string_hex = string_hex + marker1 + marker2
			int1 = (i * 2) + 3
			int2 = (i * 2) + 4
			marker1 = string(d.Data[nextClusterId].Data[int1])
			marker2 = string(d.Data[nextClusterId].Data[int2])
		}

		resultString, err := hex.DecodeString(string_hex)
		check(err)

		fileData = fileData + string(resultString)

		nextClusterId = convertClusterPointerToClusterId(d.Data[nextClusterId].FirstPointer)
	}
	return
}

func GetFileClusterData(clusterId int, d Disk) string {
	cluster := d.Data[clusterId]
	nextRune := "01"
	var name_hex string

	for i := 0; nextRune != "00"; i++ {
		nextRune = string(cluster.Data[(2*i)+5]) + string(cluster.Data[(2*i)+6])
		if nextRune != "00" {
			name_hex += nextRune
		}
	}

	dest, err := hex.DecodeString(name_hex)
	check(err)

	return string(dest)
}

func CreateBadCluster(d Disk) Disk {
	availableBadClusterId, availableBadClusterPtr, d := getNextAvailableCluster_UpdateRoot(d)
	//Get last bad cluster in linked list
	nextBadClusterPtr := d.Root.SecondPointer
	nextBadClusterId := convertClusterPointerToClusterId(nextBadClusterPtr)
	lastBadClusterId := nextBadClusterId

	for nextBadClusterId != 0 {
		lastBadClusterId = nextBadClusterId
		nextBadClusterPtr = d.Data[nextBadClusterId].FirstPointer
		nextBadClusterId = convertClusterPointerToClusterId(nextBadClusterPtr)
	}

	//Add new bad cluster to link list
	if lastBadClusterId == 0 {
		d.Root.Data[3] = availableBadClusterPtr[0]
		d.Root.Data[4] = availableBadClusterPtr[1]
		d.Root.SecondPointer = d.Root.Data[3:5]
		d.Data[0] = d.Root
	} else {
		d.Data[lastBadClusterId].Data[1] = availableBadClusterPtr[0]
		d.Data[lastBadClusterId].Data[2] = availableBadClusterPtr[1]
		d.Data[lastBadClusterId].FirstPointer = d.Data[lastBadClusterId].Data[1:3]
	}

	//Input data into available bad cluster
	var fileHeader = d.Data[availableBadClusterId]
	fileHeader.Data[0] = rune('2')
	d.Data[availableBadClusterId].ClusterType = fileHeader.Data[0]
	fileHeader.Data[1] = rune('0')
	fileHeader.Data[2] = rune('0')
	d.Data[availableBadClusterId].FirstPointer = fileHeader.Data[1:3]

	return d
}

//Printing
func PrintCluster(c Cluster) {
	for _, b := range c.Data {
		fmt.Print(b)
	}
}

func PrintHeaderRows(totalNibbles, totalRows int) {
	padding := ((totalRows - 1) / 16) + 1
	var s string
	for i := 0; i < padding; i++ {
		s += "X"
	}
	fmt.Print(s + ": ")
	for i := 1; i < ((totalNibbles-1)/16)+1; i++ {
		fmt.Printf("%16d", i)
	}
	fmt.Println()
	fmt.Print(s + ":")
	for i := 0; i < ((totalNibbles-1)/16)+1; i++ {
		for j := 0; j < 16; j++ {
			fmt.Printf("%X", j)
		}
	}
	fmt.Println()
}

func PrintClusterWithLabel(c Cluster, position int, totalRows int) {
	//lbl := fmt.Sprintf("%X", position)
	padding := ((totalRows - 1) / 16) + 1
	fmt.Printf("%0*X:", padding, position)

	for _, b := range c.Data {
		fmt.Print(string(b))
	}
	fmt.Println()
}

func PrintDisk(d Disk) {
	PrintHeaderRows(len(d.Root.Data), len(d.Data))
	for i, c := range d.Data {
		PrintClusterWithLabel(c, i, len(d.Data))
	}
}

//Util Functions
func convertClusterPointerToClusterId(pointer []rune) int {
	src := "0x" + string(pointer[0]) + string(pointer[1])

	id, e := strconv.ParseInt(src, 0, 64)
	check(e)

	return int(id)
}

func convertClusterIdToClusterPointer(id int) []rune {
	src := []byte(strconv.Itoa(id))

	dst := make([]byte, hex.EncodedLen(len(src)))
	n := hex.Encode(dst, src)

	var ptr []rune
	if n == 1 {
		ptr[0] = rune('0')
		ptr[1] = rune(dst[0])
	} else {
		ptr[0] = rune(dst[0])
		ptr[1] = rune(dst[1])
	}

	return ptr
}

func GetFileHeaderClusterIdFromFileName(fileName string, d Disk) int {
	nextFileHeaderPtr := d.Root.ThirdPointer
	nextFileHeaderId := convertClusterPointerToClusterId(nextFileHeaderPtr)
	var currentFileName string

	for nextFileHeaderId != 0 {
		currentFileName = GetFileClusterData(nextFileHeaderId, d)
		if fileName == currentFileName {
			return nextFileHeaderId
		}

		nextFileHeaderPtr = d.Data[nextFileHeaderId].FirstPointer
		nextFileHeaderId = convertClusterPointerToClusterId(nextFileHeaderPtr)
	}

	return 0
}

func getNextAvailableCluster_UpdateRoot(d Disk) (availableFileHeaderId int, returnFileHeaderPtr []rune, updatedDisk Disk) {
	//Get Next Available Cluster and Update Root Pointer
	returnFileHeaderPtr = make([]rune, 2)
	returnFileHeaderPtr[0] = d.Root.FirstPointer[0]
	returnFileHeaderPtr[1] = d.Root.FirstPointer[1]
	availableFileHeaderId = convertClusterPointerToClusterId(returnFileHeaderPtr)
	d.Root.Data[1] = d.Data[availableFileHeaderId].FirstPointer[0]
	d.Root.Data[2] = d.Data[availableFileHeaderId].FirstPointer[1]
	d.Root.FirstPointer = d.Root.Data[1:3]
	d.Data[0] = d.Root

	updatedDisk = d
	return
}
