package main

import "fmt"
import "log"
import "os/exec"

func main() {
	path, err := exec.LookPath("mongodump")
	if err != nil {
		log.Fatal("mongodump could not be found")
	}
	fmt.Printf("mongodump is available at %s\n", path)

	dumpCmd := exec.Command("mongodump", "--host", "10.10.1.103", "--port", "27017", "--archive=file.txt")

	dumpOut, err := dumpCmd.Output()
	fmt.Println(string(dumpOut))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dumpOut))
}
