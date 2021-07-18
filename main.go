package main

import (
	"log"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

func ping(ip string) bool{

	if runtime.GOOS == "windows" {
		err := exec.Command("ping", "-n", "1", ip).Run()
		if err != nil {
			// log.Println(err)
			return false
		}

		return true
		
	}

	err := exec.Command("ping", "-c", "1", ip).Run()
	if err != nil {
		// log.Println(err)
		return false
	}

	return true
	
}

func pingThis(i int) {
	if ping("192.168.43." + strconv.Itoa(i)) == true {
		log.Println("192.168.43." + strconv.Itoa(i) + " found!")
	}
}

func main() {
	for i := 1; i <= 255; i++ {
		go pingThis(i)
	}
	for {
		time.Sleep(time.Hour)
	}
	
}