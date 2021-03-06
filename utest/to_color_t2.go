// -----------------------------------------------------------------------------
// (c) markaum@gmail.com                                            License: MIT
// :v: 2018-04-29 23:42:24 3BF4C5                         [utest/to_color_t2.go]
// -----------------------------------------------------------------------------

package utest

import (
	"fmt"
	"image/color"
	"testing"

	"github.com/balacode/one-file-pdf"
)

// Test_PDF_ToColor_2_ is the second unit test for PDF.ToColor()
func Test_PDF_ToColor_2_(t *testing.T) {
	fmt.Println("Test PDF.ToColor() [2]")
	testCases := []struct {
		description string
		input       string
		color       color.RGBA
		errMsg      string
		errVal      string
	}{
		{
			description: "valid hex",
			input:       "#c83296",
			color:       color.RGBA{200, 50, 150, 255},
		},
		{
			description: "hex with more than seven characters",
			input:       "#c83296XXXXXXX",
			color:       color.RGBA{200, 50, 150, 255},
		},
		{
			description: "invalid hex",
			input:       "#wrongcolor",
			color:       color.RGBA{A: 255},
			errMsg:      "Bad color code",
			errVal:      "#wrongcolor",
		},
		// X is not a valid hex char. Only valid values are: 0-9 and A-F
		{
			description: "hex with an invalid character",
			input:       "#845X76",
			color:       color.RGBA{A: 255},
			errMsg:      "Bad color code",
			errVal:      "#845X76",
		},
		{
			description: "valid color name",
			input:       "MEDIUMPURPLE",
			color:       color.RGBA{147, 112, 219, 255},
		},
		{
			description: "valid lowercase color name",
			input:       "mediumpurple",
			color:       color.RGBA{147, 112, 219, 255},
		},
		{
			description: "unknown color name",
			input:       "picasso",
			color:       color.RGBA{A: 255},
			errMsg:      "Unknown color name",
			errVal:      "picasso",
		},
	}
	for _, test := range testCases {
		var doc pdf.PDF
		t.Run(test.description, func(t *testing.T) {
			color, err := doc.ToColor(test.input)
			var inf = doc.ErrorInfo(err)
			if inf.Msg != test.errMsg {
				t.Fatalf("expected error message %q got %q",
					test.errMsg, inf.Msg)
			}
			if inf.Val != test.errVal {
				t.Fatalf("expected error message %v got %v",
					test.errVal, inf.Val)
			}
			if test.color != color {
				t.Fatalf("expected color %v got %v", test.color, color)
			}
		})
	}
} //                                                         Test_PDF_ToColor_2_

//end
