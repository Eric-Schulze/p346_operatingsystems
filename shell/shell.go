package main

import (
	"flag"
	//"fmt"
	//"os"
	//"strconv"
	du "utils/disk_util"
)

func main() {
	formatOptPtr := flag.Bool("format", false, "Format the current shell program disk to a blank with a new name")
	dirOptPtr := flag.Bool("dir", false, "List the directory of files in the disk")
	verOptPtr := flag.Bool("v", false, "Display Information about this Shell Program, including author and version number.")
	helpOptPtr := flag.Bool("h", false, "Display help information and listing of options for the shell")
	renameOptPtr := flag.Bool("rename", false, "Update the current name for the disk")
	typeOptPtr := flag.Bool("type", false, "file type shell option")
	duOptPtr := flag.Bool("du", false, "disk usage shell option")
	sizeOptPtr := flag.Bool("filesize", false, "file size shell option")
	rawOptPtr := flag.Bool("raw", false, "Display the raw data of the current disk")
	addOptPtr := flag.Bool("add", false, "Add a new file to the current disk")
	badOptPtr := flag.Bool("bad", false, "Create a bad cluster at the next available cluster (used for function testing)")

	flag.Parse()

	args := flag.Args()

	if *formatOptPtr {
		du.Format(args)
		d := du.ReadDisk()
		du.PrintDisk(d)
	} else if *dirOptPtr {
		du.Directory()
	} else if *verOptPtr {
		du.Version()
	} else if *helpOptPtr {
		du.Help()
	} else if *renameOptPtr {
		du.Rename(args)
	} else if *typeOptPtr {
		du.FileType(args)
	} else if *duOptPtr {
		du.DiskUsage()
	} else if *sizeOptPtr {
		du.FileSize(args)
	} else if *rawOptPtr {
		d := du.ReadDisk()
		du.PrintDisk(d)
	} else if *addOptPtr {
		du.AddFile(args)
	} else if *badOptPtr {
		du.BadCluster()
	}
}
