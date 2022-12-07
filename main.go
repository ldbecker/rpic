package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func rpic(dir string, rr *rand.Rand) {

	files, err := os.ReadDir(dir)
	if err != nil {
		fj := append([]string{"-a", "Preview"}, dir)
		cmd := exec.Command("open", fj...)
		var out bytes.Buffer
		cmd.Stdout = &out
		err = cmd.Run()
		if err != nil {
			panic(err)
		}
		return
	}
	filtFiles := make([]string, 0)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if strings.HasPrefix(file.Name(), ".") {
			continue
		}
		filtFiles = append(filtFiles, file.Name())
	}
	fc := len(filtFiles)
	if fc == 0 {
		fmt.Println("no files")
		return
	}
	fi := rr.Int() % fc
	fn := fmt.Sprintf("%v/%v", dir, filtFiles[fi])
	//os.StartProcess(fmt.Sprintf("open -a Preview %v", fn), []string{}, nil)
	//cmd := exec.Command(fmt.Sprintf("open -a Preview %v", fn))
	cmd := exec.Command("open", "-a", "Preview", fn)
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println(fn)
	//time.Sleep(time.Duration(200) * time.Millisecond)
}

func main() {
	argz := os.Args
	//fmt.Println("hi")

	dir := argz[len(argz)-2]
	xnum := argz[len(argz)-1]
	num := 0
	anum, err := strconv.Atoi(xnum)
	if err != nil {
		dir = xnum
		num = 1
	} else {
		num = anum
	}
	rr := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i < num; i++ {
		rpic(dir, rr)
	}

}
