package drawer

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

type Point struct {
	X, Y float64
}

type Drawer struct {
	Parameter float64
	Origin    *Point
	RotAngle  float64
	Scale     *Point
	Color     *color.RGBA
	img       *image.RGBA
}

func NewDrawer(width, height int) *Drawer {
	d := Drawer{
		Parameter: 0,
		Origin:    new(Point),
		RotAngle:  0,
		img:       image.NewRGBA(image.Rect(0, 0, width, height)),
	}
	d.Scale = &Point{X: 1, Y: 1}
	d.Color = &color.RGBA{R: 255, G: 255, B: 255, A: 255}
	draw.Draw(d.img, d.img.Bounds(), &image.Uniform{C: color.White}, image.Point{X: 0, Y: 0}, draw.Src)
	return &d
}

func (d *Drawer) pointTrans(oldX, oldY int) (x, y int) {
	return oldX, d.img.Rect.Max.Y - oldY
}

func (d *Drawer) Draw(x, y float64) {
	newX, newY := d.pointTrans(int(x), int(y))
	d.img.Set(newX, newY, d.Color)
}

//func (d *Drawer) Draw(x, y float64) {
//
//}

func (d *Drawer) Save(path string) error {
	imgFile, _ := os.Create(path)
	defer imgFile.Close()
	return png.Encode(imgFile, d.img)
}
