package main

import (
	"bytes"
	"log"
	"os"

	"github.com/webview/webview"
)

// TODO make this much faster
func readFile(fname string) string {
	f, err := os.Open(fname)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(f)

	return buffer.String()
}

func main() {
	w := webview.New(false)
	defer w.Destroy()

	args := os.Args[1:]
	// TODO maybe add tabs here later but for now this will do
	if len(args) > 1 {
		log.Fatal("Expected 1 argument to hc")
	}

	// default filename is untiled if argument was not provied
	fname := "untitled"
	if len(args) == 1 {
		fname = args[0]
	}
	// set the title of the window to be the name of the file
	w.SetTitle(fname)

	// Set the size here to look more like A4 paper
	w.SetSize(680, 840, webview.HintNone)

	w.SetHtml("")

	if len(args) == 1 {
		w.SetHtml(readFile(fname))
	}

	w.Run()
}
