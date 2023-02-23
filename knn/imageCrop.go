// copy: https://github.com/disintegration/imaging
package knn

import (
	"image"
	"runtime"
	"sync"
	"sync/atomic"
)

var maxProcs int64

// parallel processes the data in separate goroutines.
func parallel(start, stop int, fn func(<-chan int)) {
	count := stop - start
	if count < 1 {
		return
	}

	procs := runtime.GOMAXPROCS(0)
	limit := int(atomic.LoadInt64(&maxProcs))
	if procs > limit && limit > 0 {
		procs = limit
	}
	if procs > count {
		procs = count
	}

	c := make(chan int, count)
	for i := start; i < stop; i++ {
		c <- i
	}
	close(c)

	var wg sync.WaitGroup
	for i := 0; i < procs; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fn(c)
		}()
	}
	wg.Wait()
}

// imageCrop cuts out a rectangular region with the specified bounds
// from the image and returns the cropped image.
func imageCrop(img image.Image, rect image.Rectangle) *image.NRGBA {
	r := rect.Intersect(img.Bounds()).Sub(img.Bounds().Min)
	if r.Empty() {
		return &image.NRGBA{}
	}
	src := newScanner(img)
	dst := image.NewNRGBA(image.Rect(0, 0, r.Dx(), r.Dy()))
	rowSize := r.Dx() * 4
	parallel(r.Min.Y, r.Max.Y, func(ys <-chan int) {
		for y := range ys {
			i := (y - r.Min.Y) * dst.Stride
			src.scan(r.Min.X, y, r.Max.X, y+1, dst.Pix[i:i+rowSize])
		}
	})
	return dst
}
