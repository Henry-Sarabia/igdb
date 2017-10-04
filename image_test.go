package igdb

import (
	"errors"
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
		{"invalid image ID", "", Size1080p, 1, "", errors.New("invalid empty image ID")},
		{"invalid negative ratio", testImageID, SizeScreenshotHuge, -1, "", errors.New("invalid pixel display ratio")},
		{"invalid positive ratio", testImageID, SizeScreenshotBig, 3, "", errors.New("invalid pixel display ratio")},
		{"invalid image ID and ratio", "", SizeMicro, 0, "", errors.New("invalid empty image ID")},
	}

	for _, it := range imageTests {
		t.Run(it.Name, func(t *testing.T) {
			img := Image{ID: it.ID}
			url, err := img.SizedURL(it.Size, it.Ratio)
			if err != nil {
				if err.Error() != it.ExpErr.Error() {
					t.Fatalf("Expected error '%v', got '%v'", it.ExpErr, err.Error())
				}
			}

			if url != it.ExpURL {
				t.Fatalf("Expected url '%s', got '%s'", it.ExpURL, url)
			}
		})
	}
}
