package main

import (
	"flag"
	"fmt"
	"github.com/boombuler/barcode/qr"
	"github.com/fatih/color"
	"golang.org/x/crypto/ssh/terminal"
	"image"
	"io/ioutil"
	"os"
)

var message string

func init() {
	flag.StringVar(&message, "m", "example", "QR code message")
}

func main() {
	if terminal.IsTerminal(0) {
		// get message from flag
		flag.Parse()
	} else {
		// get message from pipe
		msg, _ := ioutil.ReadAll(os.Stdin)
		message = string(msg)
	}
	// generate qr code from message
	qrCode, _ := qr.Encode(message, qr.L, qr.Auto)

	// output qr code to terminal
	output(qrCode)
}

func output(image image.Image) {
	yMax := image.Bounds().Max.Y
	xMax := image.Bounds().Max.X
	y := 0
	for y < yMax {
		x := 0
		for x < xMax {
			r, _, _, _ := image.At(x, y).RGBA()
			if r < 128 {
				white := color.New(color.BgWhite)
				white.Print("   ")
			}else{
				black := color.New(color.BgBlack)
				black.Print("   ")
			}
			x++
		}
		fmt.Println("")
		y++
	}
}
