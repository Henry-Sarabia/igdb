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
		{"happy path", testImageID, SizeScreenshotMed, 1, testImageURL, nil},
		{"happy path 2x", testImageID, SizeScreenshotMed, 2, testImageURL2x, nil},
		{"invalid image ID", "", Size1080p, 1, "", ErrEmptyID},
		{"invalid negative ratio", testImageID, SizeScreenshotHuge, -1, "", ErrPixelRatio},
		{"invalid positive ratio", testImageID, SizeScreenshotBig, 3, "", ErrPixelRatio},
		{"invalid image ID and ratio", "", SizeMicro, 0, "", ErrEmptyID},
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
