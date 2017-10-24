package igdb

import (
	"errors"
	"fmt"
)

// Errors returned when creating Image URLs.
var (
	// ErrEmptyID occurs when an empty string is used as an argument in a function.
	ErrEmptyID = errors.New("igdb: id value empty")
	// ErrPixelRatio occurs when an unsupported display pixel ratio is used as an argument in a function.
	ErrPixelRatio = errors.New("igdb: invalid display pixel ratio")
)

// Image contains the URL, dimensions,
// and Cloudinary ID of a particular image.
type Image struct {
	URL    URL    `json:"url"`
	ID     string `json:"cloudinary_id"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// imageSize is the size of the image
// being hosted at a given URL.
type imageSize string

// The precise size of an image is unknown. The
// imageSizes are solely the maximum size of an
// image hosted at a given URL.
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

// SizedImageURL returns the URL of an image identified by the provided ImageID,
// image size, and display pixel ratio. The display pixel ratio only multiplies
// the resolution of the image. The current available ratios are 1 and 2.
func SizedImageURL(ImageID string, size imageSize, ratio int) (string, error) {
	if ImageID == "" {
		return "", ErrEmptyID
	}
	if ratio < 1 || ratio > 2 {
		return "", ErrPixelRatio
	}

	var dpr string

	switch ratio {
	case 1:
		dpr = ""
	case 2:
		dpr = "_2x"
	default:
		return "", ErrPixelRatio
	}

	url := fmt.Sprintf("https://images.igdb.com/igdb/image/upload/t_%s%s/%s.jpg", size, dpr, ImageID)
	return url, nil
}

// SizedURL returns the URL of this image at the provided image size
// and display pixel ratio. The display pixel ratio only multiplies
// the resolution of the image. The current available ratios are 1 and 2.
func (i Image) SizedURL(size imageSize, ratio int) (string, error) {
	return SizedImageURL(i.ID, size, ratio)
}
