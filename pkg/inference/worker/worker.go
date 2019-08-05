package worker

import (
	"bufio"
	"fmt"
	"log"
	// "os"
	"os/exec"
	// "strings"
	"io"
)

func Detect() {
	darknet, err := exec.LookPath("darknet")
	if err != nil {
		log.Fatal("darknet must be installed")
	}

	cmd := exec.Command(darknet, "detect", "cfg/yolov3.cfg", "yolov3.weights")
	r, w := setup(cmd)
	defer w.Close()

	// Start the command!
	err = cmd.Start()
	check(err)

	images := make(chan string)
	quit := make(chan int)
	go func() {
		images <- "data/test.png"
		images <- "data/test1.png"
		images <- "data/test2.png"
		quit <- 0
	}()

	token := []byte("Enter Image Path: ")
	tokenLen := len(token)

	for {
		select {
		case image := <-images:
			line := make([]byte, tokenLen)
			_, err = io.ReadFull(r, line)
			check(err)

			// TODO: Replace with filename received from channel.
			path := []byte(image + "\n")
			w.Write(path)
			fmt.Print("Inferring image:", string(path))

			peek, _ := r.Peek(tokenLen)
			for string(peek) != "Enter Image Path: " {
				label, _, err := r.ReadLine()
				check(err)
				fmt.Println(string(label))
				peek, err = r.Peek(tokenLen)
				check(err)
			}
		case <-quit:
			fmt.Println("quit")
			return
		}
	}

	// cmd.Wait()

}

func setup(cmd *exec.Cmd) (r *bufio.Reader, w io.WriteCloser) {
	// Stdout + stderr
	out, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	r = bufio.NewReader(out)

	// Stdin
	w, err = cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	return r, w
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
