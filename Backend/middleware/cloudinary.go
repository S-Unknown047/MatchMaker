package middleware

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

var cld *cloudinary.Cloudinary
var ctx context.Context

func CreateInstance() {

	CLOUDINARY_URL := os.Getenv("CLOUDINARY_URL")
	var err error
	cld, err = cloudinary.NewFromURL(CLOUDINARY_URL)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func UploadImage(file []byte) (string, string) {
	uploadResult, err := cld.Upload.Upload(context.Background(), file, uploader.UploadParams{
		ResourceType: "image",
		Folder:       "matchmaker",
	})

	if err != nil {
		fmt.Println(err)
		return "", ""
	}

	return uploadResult.PublicID, uploadResult.SecureURL
}
func DeleteImage(ctx context.Context, cld *cloudinary.Cloudinary, publicId string) error {
	resp, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID:     publicId,
		ResourceType: "image",
		Invalidate:   api.Bool(true),
	})

	if err != nil {
		return err
	}

	if resp.Result == "not found" {
		return fmt.Errorf("image with public ID '%s' not found", publicId)
	}

	if resp.Result != "ok" {
		return fmt.Errorf("failed to delete image, API returned: %s", resp.Result)
	}

	return nil
}
