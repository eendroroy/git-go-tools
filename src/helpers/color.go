package helpers

import (
  "bytes"
)

type Color struct             { value string }
func NewColor(c string) Color { color := Color{}; color.value = c; return color }
func Black() Color            { return NewColor("\033[30m") }
func Red() Color              { return NewColor("\033[31m") }
func Green() Color            { return NewColor("\033[32m") }
func Yellow() Color           { return NewColor("\033[33m") }
func Blue() Color             { return NewColor("\033[34m") }
func Magenta() Color          { return NewColor("\033[35m") }
func Cyan() Color             { return NewColor("\033[36m") }
func LightGray() Color        { return NewColor("\033[37m") }
func DarkGray() Color         { return NewColor("\033[90m") }
func LightRed() Color         { return NewColor("\033[91m") }
func LightGreen() Color       { return NewColor("\033[92m") }
func LightYellow() Color      { return NewColor("\033[93m") }
func LightBlue() Color        { return NewColor("\033[94m") }
func LightMagenta() Color     { return NewColor("\033[95m") }
func LightCyan() Color        { return NewColor("\033[96m") }
func White() Color            { return NewColor("\033[97m") }

func Colorize(str string, m Color) string {
  var c_str bytes.Buffer
  c_str.WriteString(m.value)
  c_str.WriteString(str)
  c_str.WriteString("\033[m")
  return c_str.String()
}
