package igdb

import (
	"errors"
	"fmt"
)

// Image holds an image's URL, dimensions,
// and its Cloudinary ID.
type Image struct {
	URL    URL    `json:"url"`
	ID     string `json:"cloudinary_id"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// imageSize is the size of the image
// being hosted at a given URL.
type imageSize string

// The precise size of an image is unknown, these
// imageSizes are simply the maximum size of an
// image at a given URL.
const (
	// SizeCoverSmall is sized at 90x128.
	SizeCoverSmall imageSize = "cover_small"
	// SizeCoverBig is sized at 227x320.
	SizeCoverBig imageSize = "cover_big"
	// SizeScreenshotMed is sized at 569x320.
	SizeScreenshotMed imageSize = "screenshot_med"
	// SizeScreenshotBig is sized at 889x500
	SizeScreenshotBig imageSize = "screenshot_big"
	// SizeScreenshotHuge is sized at 1280x720.
	SizeScreenshotHuge imageSize = "screenshot_huge"
	// SizeLogoMed is sized at 284x160.
	SizeLogoMed imageSize = "logo_med"
	// SizeMicro is sized at 35x35.
	SizeMicro imageSize = "micro"
	// SizeThumb is sized at 90x90.
	SizeThumb imageSize = "thumb"
	// Size720p is sized at 1280x720.
	Size720p imageSize = "720p"
	// Size1080p is sized at 1920x1080.
	Size1080p imageSize = "1080p"
)

// SizedImageURL returns a URL addressed to an image denoted by the given ImageID
// at the given size. Additional display pixel ratio images are available by
// passing in an integer. The available ratios are 1 and 2.
func SizedImageURL(ImageID string, size imageSize, ratio int) (string, error) {
	if ImageID == "" {
		return "", errors.New("invalid empty image ID")
	}
	if ratio < 1 || ratio > 2 {
		return "", errors.New("invalid pixel display ratio")
	}

	var dpr string

	switch ratio {
	case 1:
		dpr = ""
	case 2:
		dpr = "_2x"
	}

	url := fmt.Sprintf("https://images.igdb.com/igdb/image/upload/t_%s%s/%s.jpg", size, dpr, ImageID)
	return url, nil
}

// SizedURL returns a URL addressed to this image at the given size
// in the given display pixel ratio. The available ratios are 1 and 2.
func (i Image) SizedURL(size imageSize, ratio int) (string, error) {
	return SizedImageURL(i.ID, size, ratio)
}
