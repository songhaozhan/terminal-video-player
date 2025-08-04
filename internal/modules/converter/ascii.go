package converter

import (
	"fmt"
	"image"
	"strings"
)

// Quadrant block character mapping based on bit pattern
// bit 0: top-left, bit 1: top-right, bit 2: bottom-left, bit 3: bottom-right
var quadrantChars = []string{
	" ", // 0000: all dark
	"▘", // 0001: top-left only
	"▝", // 0010: top-right only
	"▀", // 0011: top half
	"▖", // 0100: bottom-left only
	"▌", // 0101: left half
	"▞", // 0110: diagonal \
	"▛", // 0111: missing bottom-right
	"▗", // 1000: bottom-right only
	"▚", // 1001: diagonal /
	"▐", // 1010: right half
	"▜", // 1011: missing bottom-left
	"▄", // 1100: bottom half
	"▙", // 1101: missing top-right
	"▟", // 1110: missing top-left
	"█", // 1111: all bright
}

// ImageToASCII converts image to smooth colored block art
func ImageToASCII(img image.Image, width, height int) string {
	bounds := img.Bounds()
	imgWidth := bounds.Max.X
	imgHeight := bounds.Max.Y

	var result strings.Builder
	result.Grow(width * height * 20) // Pre-allocate buffer space

	// Use quadrant block characters for 4x resolution (2x2 pixels per character)
	for y := 0; y < imgHeight; y += 2 {
		for x := 0; x < imgWidth; x += 2 {
			// Sample 2x2 pixel block
			r1, g1, b1, _ := img.At(x, y).RGBA()
			r2, g2, b2, _ := img.At(min(x+1, imgWidth-1), y).RGBA()
			r3, g3, b3, _ := img.At(x, min(y+1, imgHeight-1)).RGBA()
			r4, g4, b4, _ := img.At(min(x+1, imgWidth-1), min(y+1, imgHeight-1)).RGBA()

			// Calculate brightness for each quadrant using luminance formula
			b1f := (0.299*float64(r1) + 0.587*float64(g1) + 0.114*float64(b1)) / 65535.0
			b2f := (0.299*float64(r2) + 0.587*float64(g2) + 0.114*float64(b2)) / 65535.0
			b3f := (0.299*float64(r3) + 0.587*float64(g3) + 0.114*float64(b3)) / 65535.0
			b4f := (0.299*float64(r4) + 0.587*float64(g4) + 0.114*float64(b4)) / 65535.0

			// Convert to 8-bit RGB and average for color
			avgR := uint8((r1 + r2 + r3 + r4) / 4 / 257)
			avgG := uint8((g1 + g2 + g3 + g4) / 4 / 257)
			avgB := uint8((b1 + b2 + b3 + b4) / 4 / 257)

			// Determine which quadrants are "bright" (above average brightness)
			avgBrightness := (b1f + b2f + b3f + b4f) / 4.0
			threshold := avgBrightness * 0.9 // Slightly below average for better contrast

			// Create bitmask for bright quadrants (top-left, top-right, bottom-left, bottom-right)
			quadrantMask := 0
			if b1f > threshold {
				quadrantMask |= 1
			} // Top-left
			if b2f > threshold {
				quadrantMask |= 2
			} // Top-right
			if b3f > threshold {
				quadrantMask |= 4
			} // Bottom-left
			if b4f > threshold {
				quadrantMask |= 8
			} // Bottom-right

			// Map quadrant mask to Unicode block character
			char := quadrantChars[quadrantMask]

			// Output character with averaged color
			if char != " " {
				result.WriteString(fmt.Sprintf("\033[38;2;%d;%d;%dm%s", avgR, avgG, avgB, char))
			} else {
				result.WriteString(char)
			}
		}
		result.WriteString("\033[0m\n") // Reset color at end of line
	}

	return result.String()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
