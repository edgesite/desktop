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

func (m HBITMAP) Close() {
	DeleteObject.Call(Arg(m))
}

//
// Bitmap
//

type Bitmap struct {
	h HBITMAP
	i image.Image
}

func BitmapNew(i image.Image) *Bitmap {
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
			r = r << 0
			g = g << 8
			b = b << 16
			a = a << 24
			buf[x+(h-y-1)*w] = r | g | b | a
		}
	}

	return &Bitmap{hBitmap, i}
}

func (m *Bitmap) Close() {
	m.h.Close()
}

func (m *Bitmap) Size() SIZE {
	w := m.i.Bounds().Max.X - m.i.Bounds().Min.X
	h := m.i.Bounds().Max.Y - m.i.Bounds().Min.Y
	return SIZE{LONG(w), LONG(h)}
}

func (m *Bitmap) Draw(x LONG, y LONG, hdcDst HDC) {
	hbm := m.h
	cx := m.Size().cx
	cy := m.Size().cy
	
    hdcSrc := HDCPtr(CreateCompatibleDC.Call(Arg(hdcDst)))
	if hdcSrc == 0 {
		panic(GetLastErrorString())
	}
    old := HANDLEPtr(SelectObject.Call(Arg(hdcSrc), Arg(hbm)))

    bld := BLENDFUNCTION{}
    bld.BlendOp = AC_SRC_OVER
    bld.BlendFlags = 0
    bld.SourceConstantAlpha = 255
    bld.AlphaFormat = AC_SRC_ALPHA

    if !BOOLPtr(AlphaBlend.Call(Arg(hdcDst), Arg(x), Arg(y), Arg(cx), Arg(cy), Arg(hdcSrc), NULL, NULL, Arg(cx), Arg(cy), Arg(&bld))).Bool() {
        panic(GetLastErrorString())
	}

    SelectObject.Call(Arg(hdcSrc), Arg(old))
	
    if !BOOLPtr(DeleteDC.Call(Arg(hdcSrc))).Bool() {
		panic(GetLastErrorString())
	}
}
