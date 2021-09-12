package main

import (
	"bytes"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"

	"vladusenko.io/image-processing/processing"
)

func main() {
	var file *os.File
	var img image.Image
	buff := new(bytes.Buffer)
	var err error

	if file, err = os.Open("sample_images/butterflies.jpg"); err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	if img, err = jpeg.Decode(file); err != nil {
		log.Fatal(err)
	}

	if err = processing.RgbToGrayscale(img, buff); err != nil {
		log.Fatal(err)
	}

	if err = ioutil.WriteFile("output/butterflies.jpg", buff.Bytes(), 0777); err != nil {
		log.Fatal(err)
	}

	log.Println("success")
}
