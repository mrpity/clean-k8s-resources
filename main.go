package main

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"
)

func main() {

	dateCmd := exec.Command("date")
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	fmt.Println("pity")
	fmt.Printf("%q\n", strings.Split("a,b", ","))

	exec.Command("sh", "-c", "echo 1 > /tmp/1.txt")
	//fmt.Printf("%s\n\n", outtt.Run())

	wg := new(sync.WaitGroup)
	wg.Add(3)

	x := []string{"ls#-la"}
	go exeCmd(x[0], wg)

	wg.Wait()
}

func exeCmd(cmd string, wg *sync.WaitGroup) {
	fmt.Println("given string is ", cmd)
	// splitting entryPoint => main command, shParts => rest of the command
	shParts := strings.Split(cmd, "#")
	entryPoint := shParts[0]
	//fmt.Println(entryPoint)
	shParts = shParts[1:len(shParts)]
	//fmt.Println(shParts)

	out, _ := exec.Command(entryPoint, shParts...).Output()
	// if err != nil {
	// 	fmt.Printf("%s", err)
	// }
	fmt.Printf("%s", out)
	wg.Done() // Need to signal to waitgroup that this goroutine is done
}
