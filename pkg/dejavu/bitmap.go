package dejavu

import (
	"image"
	"image/draw"
)

const (
	Width      = 7
	Height     = 10
	FullHeight = 12
)

func DrawText(canvas draw.Image, text string, x, y int, src image.Image) {
	x -= Width
	maskp := image.Point{}
	zerop := image.Point{0, 0}

	for _, r := range text {
		x += Width

		if r < 0x20 && r >= 0x7f {
			continue
		}

		maskp.Y = (int(r) - 0x20) * FullHeight
		draw.DrawMask(canvas, image.Rect(x, y-Height, x+Width, y+FullHeight-Height), src, zerop, mask, maskp, draw.Over)
	}
}

var mask = &image.Alpha{
	Stride: Width,
	Rect:   image.Rectangle{Max: image.Point{Width, 95 * FullHeight}},
	Pix: []byte{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 245, 0, 0, 0, 0, 0, 0, 214, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 0, 255, 0, 0, 0, 0, 255, 0, 255, 0, 0, 0, 0, 255, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 148, 52, 152, 48, 0, 0, 0, 196, 3, 198, 2, 0, 112, 255, 255, 255, 255, 255, 0, 0, 88, 112, 92, 108, 0, 0, 255, 255, 255, 255, 255, 112, 0, 0, 198, 1, 199, 0, 0, 0, 32, 168, 36, 164, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 107, 231, 255, 255, 255, 0, 0, 247, 38, 255, 0, 0, 0, 0, 181, 160, 255, 56, 0, 0, 0, 3, 70, 255, 165, 169, 0, 0, 0, 0, 255, 45, 243, 0, 0, 251, 240, 255, 228, 93, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 138, 247, 138, 0, 0, 0, 0, 246, 59, 244, 0, 0, 6, 0, 138, 243, 134, 30, 146, 94, 0, 0, 20, 137, 134, 19, 0, 0, 78, 147, 31, 138, 247, 136, 0, 3, 0, 0, 246, 59, 244, 0, 0, 0, 0, 138, 243, 134, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 151, 248, 255, 0, 0, 0, 0, 240, 31, 0, 0, 0, 0, 4, 185, 147, 0, 0, 0, 0, 149, 86, 230, 47, 248, 0, 0, 245, 7, 104, 189, 187, 0, 0, 219, 135, 14, 218, 63, 0, 0, 58, 215, 231, 150, 174, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 168, 0, 0, 0, 0, 0, 93, 115, 0, 0, 0, 0, 0, 181, 52, 0, 0, 0, 0, 0, 235, 14, 0, 0, 0, 0, 0, 252, 2, 0, 0, 0, 0, 0, 234, 15, 0, 0, 0, 0, 0, 179, 54, 0, 0, 0, 0, 0, 91, 117, 0, 0, 0, 0, 0, 5, 170, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 170, 4, 0, 0, 0, 0, 0, 118, 89, 0, 0, 0, 0, 0, 54, 178, 0, 0, 0, 0, 0, 16, 230, 0, 0, 0, 0, 0, 4, 250, 0, 0, 0, 0, 0, 17, 230, 0, 0, 0, 0, 0, 56, 178, 0, 0, 0, 0, 0, 119, 89, 0, 0, 0, 0, 0, 169, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 161, 28, 255, 29, 160, 0, 0, 28, 167, 255, 167, 27, 0, 0, 29, 168, 255, 169, 28, 0, 0, 160, 27, 255, 28, 159, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 255, 255, 255, 255, 255, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 64, 192, 0, 0, 0, 0, 0, 192, 64, 0, 0, 0, 0, 64, 192, 0, 0, 0, 0, 0, 192, 64, 0, 0, 0, 0, 64, 192, 0, 0, 0, 0, 0, 192, 64, 0, 0, 0, 0, 64, 192, 0, 0, 0, 0, 0, 192, 64, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 33, 203, 252, 199, 30, 0, 0, 174, 130, 6, 132, 168, 0, 0, 239, 21, 0, 24, 233, 0, 0, 254, 0, 195, 4, 252, 0, 0, 238, 22, 0, 25, 233, 0, 0, 172, 135, 10, 137, 165, 0, 0, 30, 197, 248, 194, 27, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 255, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 94, 220, 251, 209, 67, 0, 0, 171, 45, 5, 88, 233, 0, 0, 0, 0, 0, 30, 249, 3, 0, 0, 0, 22, 202, 139, 0, 0, 0, 51, 223, 132, 1, 0, 0, 100, 200, 54, 0, 0, 0, 0, 255, 255, 255, 255, 255, 12, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 83, 213, 250, 205, 62, 0, 0, 167, 40, 5, 81, 230, 0, 0, 0, 0, 12, 90, 213, 0, 0, 0, 255, 255, 236, 49, 0, 0, 0, 0, 9, 88, 212, 0, 0, 157, 31, 10, 89, 236, 0, 0, 92, 217, 250, 209, 71, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 122, 255, 0, 0, 0, 0, 49, 234, 255, 0, 0, 0, 7, 212, 90, 255, 0, 0, 0, 142, 168, 0, 255, 0, 0, 0, 254, 255, 255, 255, 255, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 240, 248, 185, 32, 0, 0, 0, 0, 20, 136, 198, 0, 0, 0, 0, 0, 9, 248, 0, 0, 0, 0, 15, 125, 198, 0, 0, 255, 255, 247, 185, 35, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 149, 237, 255, 255, 0, 0, 143, 206, 38, 0, 0, 0, 0, 232, 60, 0, 0, 0, 0, 0, 254, 136, 247, 229, 86, 0, 0, 239, 72, 6, 68, 235, 0, 0, 177, 76, 9, 69, 232, 0, 0, 35, 199, 248, 219, 77, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 235, 0, 0, 0, 0, 0, 150, 132, 0, 0, 0, 0, 18, 239, 25, 0, 0, 0, 0, 120, 164, 0, 0, 0, 0, 4, 227, 51, 0, 0, 0, 0, 88, 196, 0, 0, 0, 0, 0, 200, 84, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 84, 221, 254, 220, 81, 0, 0, 241, 68, 6, 72, 238, 0, 0, 207, 72, 8, 75, 203, 0, 0, 55, 238, 255, 237, 51, 0, 0, 219, 70, 6, 75, 215, 0, 0, 241, 74, 8, 76, 237, 0, 0, 86, 222, 252, 220, 79, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 86, 226, 252, 201, 35, 0, 0, 238, 68, 5, 76, 176, 0, 0, 237, 66, 8, 77, 235, 0, 0, 84, 224, 245, 138, 251, 0, 0, 0, 0, 0, 68, 225, 0, 0, 0, 1, 45, 211, 131, 0, 0, 255, 254, 230, 138, 6, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 98, 201, 0, 0, 35, 142, 207, 126, 28, 0, 0, 252, 150, 5, 0, 0, 0, 0, 34, 138, 207, 130, 30, 0, 0, 0, 0, 8, 94, 199, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 201, 98, 8, 0, 0, 0, 0, 29, 126, 206, 140, 35, 0, 0, 0, 0, 6, 152, 251, 0, 0, 30, 130, 206, 136, 33, 0, 0, 199, 94, 7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 254, 217, 123, 0, 0, 0, 0, 0, 45, 243, 0, 0, 0, 0, 48, 229, 106, 0, 0, 0, 0, 217, 85, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 125, 239, 234, 72, 0, 0, 87, 202, 22, 85, 216, 0, 0, 191, 71, 129, 206, 253, 0, 0, 241, 16, 229, 77, 255, 0, 0, 254, 2, 252, 8, 255, 0, 0, 239, 19, 228, 79, 255, 0, 0, 187, 84, 127, 198, 255, 0, 0, 79, 221, 44, 0, 0, 0, 0, 0, 111, 231, 225, 2, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 44, 255, 44, 0, 0, 0, 0, 116, 253, 116, 0, 0, 0, 0, 188, 156, 188, 0, 0, 0, 11, 243, 24, 243, 11, 0, 0, 76, 255, 255, 255, 76, 0, 0, 148, 118, 0, 120, 148, 0, 0, 220, 44, 0, 46, 220, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 254, 222, 83, 0, 0, 255, 0, 0, 59, 238, 0, 0, 255, 0, 5, 74, 218, 0, 0, 255, 255, 255, 247, 73, 0, 0, 255, 0, 3, 73, 217, 0, 0, 255, 0, 2, 65, 240, 0, 0, 255, 255, 253, 221, 88, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 6, 149, 239, 243, 119, 0, 0, 138, 173, 22, 15, 138, 0, 0, 230, 31, 0, 0, 0, 0, 0, 254, 2, 0, 0, 0, 0, 0, 230, 31, 0, 0, 0, 0, 0, 137, 175, 25, 18, 141, 0, 0, 6, 142, 233, 241, 119, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 235, 151, 10, 0, 0, 255, 0, 26, 164, 148, 0, 0, 255, 0, 0, 29, 229, 0, 0, 255, 0, 0, 6, 250, 0, 0, 255, 0, 0, 31, 227, 0, 0, 255, 2, 31, 168, 144, 0, 0, 255, 253, 229, 144, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 255, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 255, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 255, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 255, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 6, 149, 239, 241, 116, 0, 0, 138, 173, 21, 14, 140, 0, 0, 230, 31, 0, 0, 0, 0, 0, 254, 1, 0, 255, 255, 0, 0, 231, 25, 0, 0, 255, 0, 0, 146, 160, 17, 26, 255, 0, 0, 10, 158, 241, 235, 122, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 255, 0, 0, 255, 0, 0, 0, 255, 0, 0, 255, 0, 0, 0, 255, 0, 0, 255, 255, 255, 255, 255, 0, 0, 255, 0, 0, 0, 255, 0, 0, 255, 0, 0, 0, 255, 0, 0, 255, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 255, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 255, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 5, 252, 0, 0, 0, 138, 16, 77, 219, 0, 0, 0, 109, 237, 232, 85, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 105, 182, 8, 0, 255, 0, 105, 182, 8, 0, 0, 255, 105, 182, 8, 0, 0, 0, 255, 232, 123, 0, 0, 0, 0, 255, 15, 199, 75, 0, 0, 0, 255, 0, 31, 222, 40, 0, 0, 255, 0, 0, 73, 216, 17, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 114, 0, 116, 255, 0, 0, 255, 214, 1, 217, 255, 0, 0, 255, 196, 124, 192, 255, 0, 0, 255, 90, 253, 88, 255, 0, 0, 255, 0, 0, 0, 255, 0, 0, 255, 0, 0, 0, 255, 0, 0, 255, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 116, 0, 0, 255, 0, 0, 255, 216, 12, 0, 255, 0, 0, 255, 124, 122, 0, 255, 0, 0, 255, 13, 219, 16, 255, 0, 0, 255, 0, 120, 130, 255, 0, 0, 255, 0, 12, 222, 255, 0, 0, 255, 0, 0, 118, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 37, 206, 252, 204, 34, 0, 0, 180, 119, 6, 125, 173, 0, 0, 240, 16, 0, 22, 235, 0, 0, 254, 0, 0, 4, 252, 0, 0, 240, 17, 0, 23, 235, 0, 0, 179, 123, 9, 129, 171, 0, 0, 35, 201, 248, 198, 31, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 253, 217, 78, 0, 0, 255, 0, 4, 80, 234, 0, 0, 255, 0, 6, 85, 234, 0, 0, 255, 255, 251, 212, 75, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 37, 206, 252, 204, 34, 0, 0, 180, 119, 6, 125, 175, 0, 0, 240, 16, 0, 22, 235, 0, 0, 254, 0, 0, 4, 252, 0, 0, 240, 17, 0, 23, 233, 0, 0, 178, 123, 9, 129, 166, 0, 0, 34, 201, 252, 229, 25, 0, 0, 0, 0, 8, 165, 78, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 252, 213, 76, 0, 0, 255, 0, 5, 81, 235, 0, 0, 255, 0, 5, 76, 218, 0, 0, 255, 254, 255, 226, 38, 0, 0, 255, 0, 10, 177, 128, 0, 0, 255, 0, 0, 30, 236, 32, 0, 255, 0, 0, 0, 120, 178, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 64, 206, 250, 220, 90, 0, 0, 232, 93, 7, 32, 159, 0, 0, 220, 127, 58, 10, 0, 0, 0, 36, 149, 204, 237, 91, 0, 0, 0, 0, 0, 43, 233, 0, 0, 165, 37, 8, 93, 236, 0, 0, 93, 220, 251, 212, 77, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 255, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 255, 0, 0, 255, 0, 0, 0, 255, 0, 0, 255, 0, 0, 0, 255, 0, 0, 255, 0, 0, 0, 255, 0, 0, 254, 0, 0, 1, 251, 0, 0, 220, 77, 6, 79, 219, 0, 0, 65, 213, 250, 213, 64, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 222, 42, 0, 44, 220, 0, 0, 152, 112, 0, 116, 148, 0, 0, 80, 184, 0, 188, 76, 0, 0, 13, 241, 20, 241, 10, 0, 0, 0, 192, 148, 188, 0, 0, 0, 0, 120, 251, 116, 0, 0, 0, 0, 48, 255, 46, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 238, 22, 0, 0, 24, 236, 0, 200, 60, 229, 227, 64, 196, 0, 160, 118, 219, 223, 118, 156, 0, 120, 200, 164, 168, 200, 116, 0, 80, 254, 108, 112, 254, 76, 0, 40, 255, 52, 56, 255, 36, 0, 5, 243, 5, 7, 243, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 180, 106, 0, 110, 180, 0, 0, 34, 217, 15, 225, 35, 0, 0, 0, 132, 199, 138, 0, 0, 0, 0, 48, 255, 58, 0, 0, 0, 0, 172, 159, 178, 0, 0, 0, 53, 208, 3, 218, 55, 0, 0, 188, 94, 0, 98, 188, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 190, 100, 0, 102, 188, 0, 0, 58, 224, 13, 226, 54, 0, 0, 0, 182, 199, 178, 0, 0, 0, 0, 49, 255, 48, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 248, 0, 0, 0, 0, 1, 190, 118, 0, 0, 0, 0, 110, 198, 2, 0, 0, 0, 38, 236, 38, 0, 0, 0, 3, 199, 114, 0, 0, 0, 0, 120, 196, 2, 0, 0, 0, 0, 248, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 192, 64, 0, 0, 0, 0, 0, 64, 192, 0, 0, 0, 0, 0, 0, 192, 64, 0, 0, 0, 0, 0, 64, 192, 0, 0, 0, 0, 0, 0, 192, 64, 0, 0, 0, 0, 0, 64, 192, 0, 0, 0, 0, 0, 0, 192, 64, 0, 0, 0, 0, 0, 64, 192, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 68, 249, 66, 0, 0, 0, 18, 201, 50, 201, 16, 0, 0, 162, 57, 0, 58, 160, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 255, 255, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 163, 36, 0, 0, 0, 0, 0, 19, 155, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 251, 208, 67, 0, 0, 0, 0, 1, 59, 229, 0, 0, 112, 228, 254, 255, 254, 0, 0, 248, 69, 8, 78, 255, 0, 0, 126, 235, 237, 116, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 140, 248, 219, 56, 0, 0, 255, 98, 8, 98, 210, 0, 0, 255, 8, 0, 10, 250, 0, 0, 255, 102, 11, 99, 206, 0, 0, 255, 139, 245, 213, 52, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 53, 210, 254, 255, 0, 0, 0, 214, 106, 6, 0, 0, 0, 0, 254, 8, 0, 0, 0, 0, 0, 211, 108, 8, 0, 0, 0, 0, 49, 205, 252, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 58, 219, 248, 139, 255, 0, 0, 215, 97, 8, 102, 255, 0, 0, 254, 8, 0, 12, 255, 0, 0, 211, 98, 11, 106, 255, 0, 0, 56, 214, 245, 135, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 45, 203, 252, 220, 64, 0, 0, 208, 89, 5, 69, 217, 0, 0, 253, 255, 255, 255, 253, 0, 0, 201, 80, 10, 0, 0, 0, 0, 37, 187, 246, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 105, 245, 255, 0, 0, 0, 0, 233, 45, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 255, 255, 255, 255, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 57, 220, 248, 135, 255, 0, 0, 213, 100, 8, 98, 255, 0, 0, 254, 8, 0, 10, 255, 0, 0, 211, 105, 12, 101, 255, 0, 0, 53, 215, 246, 134, 252, 0, 0, 0, 0, 3, 85, 212, 0, 0, 0, 255, 253, 217, 61, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 113, 240, 240, 101, 0, 0, 255, 104, 7, 71, 235, 0, 0, 255, 3, 0, 1, 255, 0, 0, 255, 0, 0, 0, 255, 0, 0, 255, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 255, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 56, 235, 0, 0, 0, 0, 255, 234, 95, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 22, 192, 94, 0, 0, 255, 43, 209, 65, 0, 0, 0, 255, 208, 233, 44, 0, 0, 0, 255, 21, 81, 221, 23, 0, 0, 255, 0, 0, 97, 200, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 232, 49, 0, 0, 0, 0, 0, 100, 242, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 205, 169, 223, 165, 0, 0, 255, 42, 255, 43, 240, 0, 0, 255, 0, 255, 0, 254, 0, 0, 255, 0, 255, 0, 255, 0, 0, 255, 0, 255, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 113, 240, 240, 101, 0, 0, 255, 104, 7, 71, 235, 0, 0, 255, 3, 0, 1, 255, 0, 0, 255, 0, 0, 0, 255, 0, 0, 255, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 61, 217, 254, 215, 59, 0, 0, 217, 98, 8, 102, 211, 0, 0, 254, 8, 0, 12, 250, 0, 0, 215, 102, 12, 106, 210, 0, 0, 59, 211, 250, 210, 56, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 140, 248, 218, 56, 0, 0, 255, 98, 8, 98, 209, 0, 0, 255, 8, 0, 10, 250, 0, 0, 255, 102, 11, 99, 209, 0, 0, 255, 139, 245, 213, 53, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 60, 219, 248, 136, 255, 0, 0, 214, 98, 8, 102, 255, 0, 0, 254, 8, 0, 12, 255, 0, 0, 214, 98, 11, 106, 255, 0, 0, 57, 214, 245, 132, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 141, 246, 206, 0, 0, 0, 255, 111, 9, 154, 0, 0, 0, 255, 5, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 112, 224, 253, 255, 255, 0, 0, 244, 121, 42, 6, 0, 0, 0, 77, 177, 217, 242, 129, 0, 0, 0, 0, 3, 67, 243, 0, 0, 255, 255, 251, 217, 98, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 255, 255, 255, 255, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 250, 34, 0, 0, 0, 0, 0, 142, 246, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 255, 0, 0, 255, 0, 0, 0, 255, 0, 0, 255, 0, 0, 5, 255, 0, 0, 235, 72, 9, 106, 255, 0, 0, 100, 236, 236, 106, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 208, 68, 0, 70, 206, 0, 0, 110, 164, 0, 168, 106, 0, 0, 18, 239, 34, 239, 16, 0, 0, 0, 170, 205, 166, 0, 0, 0, 0, 70, 255, 68, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 224, 32, 0, 34, 222, 0, 0, 160, 100, 255, 98, 156, 0, 0, 94, 178, 200, 176, 90, 0, 0, 28, 246, 112, 245, 24, 0, 0, 0, 217, 27, 215, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 150, 160, 0, 162, 148, 0, 0, 4, 187, 147, 182, 3, 0, 0, 0, 52, 255, 50, 0, 0, 0, 9, 209, 179, 206, 8, 0, 0, 158, 163, 0, 164, 156, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 200, 68, 0, 88, 196, 0, 0, 88, 166, 0, 206, 76, 0, 0, 4, 219, 84, 212, 1, 0, 0, 0, 120, 244, 98, 0, 0, 0, 0, 20, 236, 8, 0, 0, 0, 3, 110, 140, 0, 0, 0, 0, 255, 207, 22, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 242, 0, 0, 0, 0, 105, 240, 56, 0, 0, 0, 89, 237, 53, 0, 0, 0, 74, 234, 52, 0, 0, 0, 0, 245, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 162, 252, 0, 0, 0, 0, 0, 251, 21, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 2, 66, 233, 0, 0, 0, 0, 255, 255, 113, 0, 0, 0, 0, 2, 80, 235, 0, 0, 0, 0, 0, 2, 255, 0, 0, 0, 0, 0, 0, 250, 24, 0, 0, 0, 0, 0, 156, 249, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 251, 158, 0, 0, 0, 0, 0, 24, 247, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 235, 62, 2, 0, 0, 0, 0, 114, 254, 255, 0, 0, 0, 0, 237, 76, 2, 0, 0, 0, 0, 255, 1, 0, 0, 0, 0, 26, 247, 0, 0, 0, 0, 0, 249, 152, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 146, 247, 150, 20, 116, 0, 0, 111, 10, 121, 240, 142, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	},
}