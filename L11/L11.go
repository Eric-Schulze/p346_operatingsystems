package L11

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type thread struct {
	name  string
	delay int
}

func MakeThreads(args []string) ([]thread, error) {
	var numOfThreads = len(args) / 2
	threads := make([]thread, numOfThreads)
	var name string
	for i, _ := range args {
		if i%2 == 1 {
			delay, err := strconv.Atoi(args[i])
			num := i / 2
			if err != nil {
				var ending string
				switch num + 1 {
				case 1:
					ending = "st"
				case 2:
					ending = "nd"
				case 3:
					ending = "rd"
				default:
					ending = "th"
				}

				return make([]thread, 0), fmt.Errorf("The delay for the %v%v thread is not an integer\n", num+1, ending)
			}
			t := thread{name, delay}
			threads[num] = t
		} else {
			name = args[i]
		}
	}
	return threads, nil
}

func PrintLoop(t thread, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Thread: %s - Count %d\n", t.name, i+1)
		time.Sleep(time.Duration(t.delay) * time.Millisecond)
	}
	wg.Done()
}

func main() {
	fmt.Println("Start Time: ", time.Now().Format(time.UnixDate))
	args := os.Args[1:]

	if len(args)%2 == 1 {
		fmt.Println("Missing Argument: Each thread must have a name and delay.")
		return
	}

	var wg sync.WaitGroup
	threads, err := MakeThreads(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, _ := range threads {
		wg.Add(1)
		go PrintLoop(threads[i], &wg)
	}

	wg.Wait()
	fmt.Println("All Threads Completed.")
	fmt.Println("End Time: ", time.Now().Format(time.UnixDate))
}
