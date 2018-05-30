package disk_util

import (
	//"bufio"
	//"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	//"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getHeaderRowText(totalNibbles, totalRows int) string {
	padding := ((totalRows - 1) / 16) + 1
	var s, text string
	for i := 0; i < padding; i++ {
		s += "X"
	}
	text += (s + ": ")
	for i := 1; i < ((totalNibbles-1)/16)+1; i++ {
		text += fmt.Sprintf("%16d", i)
	}
	text += "\n"
	text += (s + ":")
	for i := 0; i < ((totalNibbles-1)/16)+1; i++ {
		for j := 0; j < 16; j++ {
			text += fmt.Sprintf("%X", j)
		}
	}
	text += "\n"

	return text
}

func getClusterWithLabelText(c Cluster, position int, totalRows int) string {
	var text string
	padding := ((totalRows - 1) / 16) + 1
	text += fmt.Sprintf("%0*X:", padding, position)

	for _, b := range c.Data {
		text += (string(b))
	}
	text += "\n"

	return text
}

func WriteDisk(d Disk) {
	b, err := json.Marshal(&d)
	check(err)

	err = ioutil.WriteFile("/home/eric_schulze/dev/go/src/utils/disk_util/disk_data", b, 0644)
	check(err)
}

func ReadDisk() Disk {
	b, err := ioutil.ReadFile("/home/eric_schulze/dev/go/src/utils/disk_util/disk_data")
	check(err)

	var d Disk
	json.Unmarshal(b, &d)

	return d
}
