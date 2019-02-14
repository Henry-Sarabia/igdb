package igdb

import (
	"errors"
	"fmt"
)

// Errors returned when creating Image URLs.
var (
	// ErrBlankID occurs when an empty string is used as an argument in a function.
	ErrBlankID = errors.New("igdb: id value empty")
	// ErrPixelRatio occurs when an unsupported display pixel ratio is used as an argument in a function.
	ErrPixelRatio = errors.New("igdb: invalid display pixel ratio")
)

//go:generate gomodifytags -file $GOFILE -struct Image -add-tags json -w

// Image contains the URL, dimensions, and ID of a particular image.
// For more information visit: https://api-docs.igdb.com/#images
type Image struct {
	AlphaChannel bool   `json:"alpha_channel"`
	Animated     bool   `json:"animated"`
	Height       int    `json:"height"`
	ImageID      string `json:"image_id"`
	URL          string `json:"url"`
	Width        int    `json:"width"`
}

// imageSize is the size of an image from the IGDB API. Note that this is not
// the precise size of an image, but rather the maximum possible size of
// the referenced image.
type imageSize string

// Available image sizes supported by the IGDB API
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

// SizedImageURL returns the URL of an image identified by the provided imageID,
// image size, and display pixel ratio. The display pixel ratio only multiplies
// the resolution of the image. The current available ratios are 1 and 2.
//TODO: factor out imageID check
func SizedImageURL(imageID string, size imageSize, ratio int) (string, error) {
	if imageID == "" {
		return "", ErrBlankID
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

	url := fmt.Sprintf("https://images.igdb.com/igdb/image/upload/t_%s%s/%s.jpg", size, dpr, imageID)
	return url, nil
}

// SizedURL returns the URL of this image at the provided image size
// and display pixel ratio. The display pixel ratio only multiplies
// the resolution of the image. The current available ratios are 1 and 2.
func (i Image) SizedURL(size imageSize, ratio int) (string, error) {
	return SizedImageURL(i.ImageID, size, ratio)
}
