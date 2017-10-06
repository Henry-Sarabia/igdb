package igdb

import (
	"reflect"
	"testing"
)

const testImageID = "dfgkfivjrhcksyymh9vw"
const testImageURL = "https://images.igdb.com/igdb/image/upload/t_screenshot_med/dfgkfivjrhcksyymh9vw.jpg"
const testImageURL2x = "https://images.igdb.com/igdb/image/upload/t_screenshot_med_2x/dfgkfivjrhcksyymh9vw.jpg"

func TestSizedImageURL(t *testing.T) {
	var imageTests = []struct {
		Name   string
		ID     string
		Size   imageSize
		Ratio  int
		ExpURL string
		ExpErr error
	}{
		{"Non-empty ID and valid single ratio", testImageID, SizeScreenshotMed, 1, testImageURL, nil},
		{"Non-empty ID and valid double ratio", testImageID, SizeScreenshotMed, 2, testImageURL2x, nil},
		{"Empty ID and valid ratio", "", Size1080p, 1, "", ErrEmptyID},
		{"Non-empty ID and invalid negative ratio", testImageID, SizeScreenshotHuge, -1, "", ErrPixelRatio},
		{"Non-empty ID and invalid triple ratio", testImageID, SizeScreenshotBig, 3, "", ErrPixelRatio},
		{"Empty ID and invalid 0 ratio", "", SizeMicro, 0, "", ErrEmptyID},
	}

	for _, it := range imageTests {
		t.Run(it.Name, func(t *testing.T) {
			img := Image{ID: it.ID}
			url, err := img.SizedURL(it.Size, it.Ratio)
			if !reflect.DeepEqual(err, it.ExpErr) {
				t.Fatalf("Expected error '%v', got '%v'", it.ExpErr, err)
			}

			if url != it.ExpURL {
				t.Fatalf("Expected url '%s', got '%s'", it.ExpURL, url)
			}
		})
	}
}
