package util

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/matsuyoshi30/song2"
	"github.com/nfnt/resize"
)

func IsImage(contentType string) bool {
	return strings.Split(contentType, "/")[0] == "image"
}

func ReadFile(file io.Reader, fileName string) {
	stream, readErr := ioutil.ReadAll(file)
	if readErr != nil {
		fmt.Printf("error from file %v", readErr)
	}

	fileErr := ioutil.WriteFile(fileName, stream, 0644)
	if fileErr != nil {
		fmt.Printf("file err %v", fileErr)
	}
}

func CalculateHash(file io.Reader, fileName string) string {
	tempFile, openErr := os.Open(fileName)
	if openErr != nil {
		fmt.Printf("Error opening file: %v", openErr)
	}

	encryption := sha256.New()
	if _, err := io.Copy(encryption, tempFile); err != nil {
		log.Fatal(err)
	}
	hash := fmt.Sprintf("%x\n", encryption.Sum(nil))

	defer tempFile.Close()

	return hash
}

func CreateImage(img image.Image, name string, blur bool, width *uint) (int64, image.Image, error) {
	var blurDegree float64
	var imageWidth uint
	var quality int

	if blur {
		blurDegree = float64(img.Bounds().Size().X / 16)
		quality = 95
	} else {
		blurDegree = 0
		quality = 50
	}

	if width != nil {
		imageWidth = *width
	} else {
		imageWidth = uint(img.Bounds().Size().X)
	}

	file, err := os.Create(fmt.Sprintf("./%v", name))
	if err != nil {
		fmt.Println(err)
		return 0, nil, err
	}
	defer file.Close()

	blurred := song2.GaussianBlur(img, blurDegree)
	resized := resize.Resize(imageWidth, 0, blurred, resize.Lanczos3)

	if err := jpeg.Encode(file, resized, &jpeg.Options{Quality: quality}); err != nil {
		fmt.Println(err)
		return 0, nil, err
	}

	stat, _ := os.Stat(fmt.Sprintf("./%v", name))

	return stat.Size(), resized, nil
}

func ReadImage(fileName string, fileSize int64) *bytes.Reader {
	f, openErr := os.Open(fileName)
	if openErr != nil {
		fmt.Printf("Error opening file: %v", openErr)
	}

	defer f.Close()

	buffer := make([]byte, fileSize)

	_, readErr := f.Read(buffer)
	if readErr != nil {
		fmt.Printf("Error reading file: %v", readErr)
	}

	fileBytes := bytes.NewReader(buffer)

	return fileBytes
}

func ConvertFileToImage(fileName string, fileSize int64) (image.Image, error) {
	fileBytes := ReadImage(fileName, fileSize)
	img, _, err := image.Decode(fileBytes)
	if err != nil {
		fmt.Printf("Error decoding image: %v", err)
	}

	return img, err
}
