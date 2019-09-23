package helper

import (
	"encoding/base64"
	"image/jpeg"
	"image/png"
	"strings"
	"image"
	"os"
)

/*
	Convert base64 of image to standard img file
	Currently only support jpeg and png
*/
func Base64ToImg (tmp string, src string) (*os.File, error) {
	imageReader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(src))
	imgData, imgType, err := image.Decode(imageReader)
	if err != nil {
		return nil, err
	}

	img, err := os.Create(tmp)
	if err != nil {
		return nil, err
	}
	defer img.Close()

	if "jpeg" == imgType {
		jpeg.Encode(img, imgData, &jpeg.Options{Quality:75})
	} else if "png" == imgType {
		png.Encode(img, imgData)
	}

	return img, nil
}