package thirdparty

// HexToRGB from https://www.cnblogs.com/dfsxh/articles/11082576.html
func HexToRGB(color int) (red, green, blue uint8) {
	red = uint8(color >> 16)
	green = uint8((color & 0x00FF00) >> 8)
	blue = uint8(color & 0x0000FF)
	return
}
