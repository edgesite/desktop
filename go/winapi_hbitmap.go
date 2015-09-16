// build +windows

package desktop

import (
	"image"
	"syscall"
	"unsafe"
)

type HBITMAP uintptr

func HBITMAPPtr(r1, r2 uintptr, err error) HBITMAP {
	LastError = uintptr(err.(syscall.Errno))
	return HBITMAP(r1)
}

func HBITMAPNew(i image.Image) HBITMAP {
	w := i.Bounds().Max.X - i.Bounds().Min.X
	h := i.Bounds().Max.Y - i.Bounds().Min.Y

	screenDC := HDCPtr(GetDC.Call(NULL))
	if screenDC == 0 {
		panic(GetLastErrorString())
	}
	defer ReleaseDC.Call(NULL, Arg(screenDC))

	memDC := HDCPtr(CreateCompatibleDC.Call(Arg(screenDC)))
	defer DeleteDC.Call(Arg(memDC))
	if memDC == 0 {
		panic(GetLastErrorString())
	}

	bmi := &BITMAPINFO{}
	bmi.bmiHeader.biSize = DWORD(unsafe.Sizeof(bmi.bmiHeader))
	bmi.bmiHeader.biWidth = LONG(w)
	bmi.bmiHeader.biHeight = LONG(h)
	bmi.bmiHeader.biPlanes = 1
	bmi.bmiHeader.biBitCount = 32
	bmi.bmiHeader.biCompression = BI_RGB
	bmi.bmiHeader.biSizeImage = DWORD(w * h * 4)

	var ppvBits unsafe.Pointer

	hBitmap := HBITMAPPtr(CreateDIBSection.Call(Arg(memDC), Arg(bmi), Arg(DIB_RGB_COLORS), Arg(&ppvBits), NULL, NULL))
	if hBitmap == 0 {
		panic(GetLastErrorString())
	}

	buf := (*(*[1 << 30]uint32)(ppvBits))[:(w * h)]

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			c := i.At(i.Bounds().Min.X+x, i.Bounds().Min.Y+y)
			r, g, b, a := c.RGBA()
			b = b << 0
			g = g << 8
			r = r << 16
			a = a << 24
			buf[x+(h-y-1)*w] = r | g | b | a
		}
	}

	return hBitmap
}

func (m HBITMAP) Close() {
	DeleteObject.Call(Arg(m))
}
