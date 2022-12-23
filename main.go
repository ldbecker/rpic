package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"time"
)

func listFiles(dir string) ([]string, error) {
	xstat, err := os.Stat(dir)
	if err != nil {
		return []string{}, err
	}
	if !xstat.IsDir() && xstat.Mode().IsRegular() {
		return []string{path.Join(dir, xstat.Name())}, nil
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		return []string{}, err
	}
	filtFiles := make([]string, 0)
	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".") {
			continue
		}
		if file.IsDir() {
			dirf, err1 := listFiles(path.Join(dir, file.Name()))
			if err1 != nil {
				return []string{}, err1
			}
			filtFiles = append(filtFiles, dirf...)
			continue
		}

		filtFiles = append(filtFiles, path.Join(dir, file.Name()))
	}
	return filtFiles, nil
}

func rpic(filtFiles []string, rr *rand.Rand) error {

	fc := len(filtFiles)
	fi := rr.Float64() * float64(fc)
	fn := filtFiles[int(fi)]
	fmt.Println(fn)
	//os.StartProcess(fmt.Sprintf("open -a Preview %v", fn), []string{}, nil)
	//cmd := exec.Command(fmt.Sprintf("open -a Preview %v", fn))
	cmd := exec.Command("open", "-a", "/Applications/VLC.app", fn)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%v\n", err)
		panic(err)
		return rpic(filtFiles, rr)
	}
	///      s sfmt.Println(fn)
	time.Sleep(time.Duration(200) * time.Millisecond)
	return nil
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
	var ff []string
	xtt := time.Now()
	dstr := fmt.Sprintf("./%v-%v-%v-%v.json", path.Base(dir), xtt.Year(), xtt.Month(), xtt.Day())
	fmt.Printf("dstr = %v\n", dstr)
	if strings.HasSuffix(dir, ".json") {
		jb, errs := os.ReadFile(dir)
		if errs !=nil{
			panic(err)
		}
		json.Unmarshal(jb, &ff)
	} else {
		if _, ok := os.Stat(dstr); ok == nil {
			jbb, errx := os.ReadFile(dstr)
			if errx != nil {
				panic(errx)
			}
			errx = json.Unmarshal(jbb, &ff)
			if errx != nil {
				panic(errx)
			}
		} else {
			ff, err = listFiles(dir)
			if err != nil {
				panic(err)
			}
			jbo, _ := json.Marshal(ff)
			os.WriteFile(dstr, jbo, 0755)
		}
	}


	if num == 0 {
		jbl, jerr := json.Marshal(ff)
		if jerr != nil {
			panic(err)
		}
		os.WriteFile(dstr, jbl, 0755)
		for _, e := range ff {
			fmt.Println(e)
		}
		return
	}
	fmt.Printf("%v of %v files\n", num, len(ff))
	for i := 0; i < num; i++ {
		err = rpic(ff, rr)
		if err != nil {
			fmt.Printf("rpic err: %v\n", err)
			panic(err)
		}
	}
}
