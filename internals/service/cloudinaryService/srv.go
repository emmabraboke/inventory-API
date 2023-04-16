package cloudinaryService

import (
	"context"
	"inventory/internals/entity/userEntity"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"log"
)

type cloudinarySrv struct {
	cloudURI string
}

type CloudinaryService interface {
	ImageUpload(file userEntity.ImageFile) (*string, error)
	CloudinaryInstance() *cloudinary.Cloudinary
}

func NewCloudinarySrv(cloudURI string) CloudinaryService {
	return &cloudinarySrv{cloudURI: cloudURI}
}

func (t *cloudinarySrv) CloudinaryInstance() *cloudinary.Cloudinary {
	cld, err := cloudinary.NewFromURL(t.cloudURI)
	if err != nil {
		log.Fatal(err)
	}

	return cld
}

func (t *cloudinarySrv) ImageUpload(file userEntity.ImageFile) (*string, error) {
	ctx := context.Background()
	cld := t.CloudinaryInstance()
   

	uploadResult, err := cld.Upload.Upload(ctx, file.File, uploader.UploadParams{
		ResourceType: "image",
	})

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &uploadResult.SecureURL, nil
}
