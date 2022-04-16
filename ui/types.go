package ui

import (
  "fyne.io/fyne/v2"
  "kishorenewton/pixl/apptype"
  "kishorenewton/pixl/swatch"
)

type AppInit struct {
  PixlWindow fyne.Window
  State *apptype.State
  Swatches []*swatch.Swatch
}
