package configs

import (
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
)

var Cloudinary *cloudinary.Cloudinary

func InitCloudinary() {

	var err error
	Cloudinary, err = cloudinary.NewFromParams(
		os.Getenv("CLOUDINARY_NAME"),
		os.Getenv("CLOUDINARY_KEY"),
		os.Getenv("CLOUDINARY_SECRET"))

	if err != nil {
		panic("Error init cloudinary")
	}

}
