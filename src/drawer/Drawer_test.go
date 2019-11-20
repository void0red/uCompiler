package drawer

import "testing"

func TestNewDrawer(t *testing.T) {
	d := NewDrawer(100, 80)
	d.Draw(1, 1)
	_ = d.Save("test.png")
}
