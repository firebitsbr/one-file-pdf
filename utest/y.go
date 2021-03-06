// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-05-13 01:54:23 57FF6D                                   [utest/y.go]
// -----------------------------------------------------------------------------

package utest

import (
	"fmt"
	"testing"

	"github.com/balacode/one-file-pdf"
)

// Test_PDF_Y_ is the unit test for PDF.Y()
func Test_PDF_Y_(t *testing.T) {
	fmt.Println("Test PDF.Y()")
	//
	// Y of new PDF must be -1
	func() {
		var doc pdf.PDF
		TEqual(t, doc.Y(), -1)
	}()
	func() {
		var doc = pdf.NewPDF("A4")
		TEqual(t, doc.Y(), -1)
	}()
	// SetY() sets the property?
	func() {
		var doc pdf.PDF
		doc.SetY(321)
		TEqual(t, doc.Y(), 321)
	}()
	func() {
		var doc = pdf.NewPDF("A4")
		doc.SetY(654)
		TEqual(t, doc.Y(), 654)
	}()
	// -------------------------------------------------------------------------
	// Test PDF generation
	func() {
		var doc = pdf.NewPDF("A4")
		doc.SetCompression(false).
			SetUnits("cm").
			SetXY(1, 10).
			SetFont("Times-Bold", 20).
			DrawText("X=1 Y=10")
		const expect = `
		%PDF-1.4
		1 0 obj <</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R
		/Resources <</Font <</FNT1 5 0 R>> >> >>
		endobj
		4 0 obj <</Length 91>> stream
		BT /FNT1 20 Tf ET
		0.000 0.000 0.000 rg
		0.000 0.000 0.000 RG
		BT 28 558 Td (X=1 Y=10) Tj ET
		endstream
		endobj
		5 0 obj <</Type/Font/Subtype/Type1/Name/FNT1
		/BaseFont/Times-Bold
		/Encoding/StandardEncoding>>
		endobj
		xref
		0 6
		0000000000 65535 f
		0000000010 00000 n
		0000000056 00000 n
		0000000130 00000 n
		0000000228 00000 n
		0000000369 00000 n
		trailer
		<</Size 6/Root 1 0 R>>
		startxref
		471
		%%EOF
		`
		ComparePDF(t, doc.Bytes(), expect)
		// doc.SaveFile("~~Test_PDF_Y_.pdf")
	}()
} //                                                                 Test_PDF_Y_

//end
