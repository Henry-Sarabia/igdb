package igdb

import (
	"github.com/pkg/errors"
	"testing"
)

// Mocked image arguments for testing.
const (
	testImageID    = "dfgkfivjrhcksyymh9vw"
	testImageURL   = "https://images.igdb.com/igdb/image/upload/t_screenshot_med/dfgkfivjrhcksyymh9vw.jpg"
	testImageURL2x = "https://images.igdb.com/igdb/image/upload/t_screenshot_med_2x/dfgkfivjrhcksyymh9vw.jpg"
)

func TestSizedImageURL(t *testing.T) {
	var tests = []struct {
		name    string
		id      string
		size    imageSize
		ratio   int
		wantURL string
		wantErr error
	}{
		{"Non-empty ID and valid single ratio", testImageID, SizeScreenshotMed, 1, testImageURL, nil},
		{"Non-empty ID and valid double ratio", testImageID, SizeScreenshotMed, 2, testImageURL2x, nil},
		{"Non-empty ID and invalid negative ratio", testImageID, SizeScreenshotHuge, -1, "", ErrPixelRatio},
		{"Non-empty ID and invalid triple ratio", testImageID, SizeScreenshotBig, 3, "", ErrPixelRatio},
		{"Empty ID and valid ratio", "", Size1080p, 1, "", ErrBlankID},
		{"Empty ID and invalid 0 ratio", "", SizeMicro, 0, "", ErrBlankID},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			url, err := SizedImageURL(test.id, test.size, test.ratio)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if url != test.wantURL {
				t.Errorf("got: <%v>, want: <%v>", url, test.wantURL)
			}
		})
	}
}

func TestImage_SizedURL(t *testing.T) {
	var tests = []struct {
		name    string
		image   Image
		size    imageSize
		ratio   int
		wantURL string
		wantErr error
	}{
		{"Non-empty ID and valid single ratio", Image{ImageID: testImageID}, SizeScreenshotMed, 1, testImageURL, nil},
		{"Non-empty ID and valid double ratio", Image{ImageID: testImageID}, SizeScreenshotMed, 2, testImageURL2x, nil},
		{"Non-empty ID and invalid negative ratio", Image{ImageID: testImageID}, SizeScreenshotHuge, -1, "", ErrPixelRatio},
		{"Non-empty ID and invalid triple ratio", Image{ImageID: testImageID}, SizeScreenshotBig, 3, "", ErrPixelRatio},
		{"Empty ID and valid ratio", Image{}, Size1080p, 1, "", ErrBlankID},
		{"Empty ID and invalid 0 ratio", Image{}, SizeMicro, 0, "", ErrBlankID},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			url, err := test.image.SizedURL(test.size, test.ratio)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if url != test.wantURL {
				t.Errorf("got: <%v>, want: <%v>", url, test.wantURL)
			}
		})
	}
}
