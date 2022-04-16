package swatch 

import (
  "fyne.io/fyne/v2"
  "fyne.io/fyne/v2/canvas"
)

type SwatchRenderer struct {
  square canvas.Rectangle
  objects []fyne.CanvasObject
  parent *Swatch
}

func (renderer *SwatchRenderer) MinSize() fyne.Size {
  return renderer.square.MinSize()
}

func (renderer *SwatchRenderer) Layout(size fyne.Size) {
  renderer.objects[0].Resize(size)
}
