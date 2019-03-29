package main

import (
	"fmt"
	"os/exec"
	"reflect"
	"strings"
	"sync"
)

func main() {

	// dateCmd := exec.Command("date")
	// dateOut, err := dateCmd.Output()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("> date")
	// fmt.Println(string(dateOut))

	k8sFilterKey := "-o"
	k8sFilterExpression := "jsonpath='{range .items[*]} {.metadata.name} {.status.startTime}  {end}'"
	getMoonPods := "kubectl get pod -n moon"

	wg := new(sync.WaitGroup)
	wg.Add(1)

	go exeCmd(getMoonPods, k8sFilterKey, k8sFilterExpression, wg)

	wg.Wait()
}

func exeCmd(cmd string, k8sFilterKey string, k8sFilterExpression string, wg *sync.WaitGroup) {
	fmt.Println("given string is ", cmd)
	// splitting entryPoint => main command, shParts => rest of the command
	shParts := strings.Split(cmd, " ")
	entryPoint := shParts[0]
	shParts = shParts[1:len(shParts)]
	shParts = append(shParts, k8sFilterKey, k8sFilterExpression)
	fmt.Println(shParts)

	out, err := exec.Command(entryPoint, shParts...).Output()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", out)

	// filter CLI pout
	filterK8sOutput(out)
	wg.Done() // Need to signal to waitgroup that this goroutine is done
}

func filterK8sOutput(out []byte) {

	fmt.Println(reflect.TypeOf(out))
	ConvertToString := strings.Trim(string(out), "'")
	fmt.Println(reflect.TypeOf(ConvertToString))
	fmt.Println("OUTPUT: ", ConvertToString)
	TrimString := strings.Trim(ConvertToString, " ")
	fmt.Println(len(TrimString))
	fmt.Println(ConvertToString)
	SplitString := strings.Split(ConvertToString, "   ")
	fmt.Println(len(SplitString))
	fmt.Println("OLOLO: ", SplitString[0])

}
