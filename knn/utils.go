package knn

import (
	"bufio"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/fs"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func readFile(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func toVector(content string) []int32 {
	s_list := strings.Split(content, "\n")
	for i, o := range s_list {
		s_list[i] = strings.TrimSpace(o)
	}
	out := strings.Join(s_list, "")

	ret := make([]int32, len(out))
	for i, o := range out {
		ret[i] = o - '0'
	}
	return ret
}

func knnCompute(a []int32, b []int32) float64 {
	if len(a) != len(b) {
		panic("a and b must be same length")
	}

	ret := 0.0
	for i := 0; i < len(a); i++ {
		tp := math.Abs(float64(a[i]) - float64(b[i]))
		tp = math.Pow(tp, 2)
		ret += tp
	}
	return math.Pow(ret, 0.5)
}

func walkPath(file_dir string) []string {
	files := []string{}

	if err := filepath.WalkDir(file_dir, func(path string, d fs.DirEntry, err error) error {
		if filepath.Ext(path) == ".txt" {
			files = append(files, path)
		}
		return nil
	}); err != nil {
		panic(err)
	}
	return files
}

func readGrayImage(imagePath string) [][]int {
	f, err := os.Open(imagePath)
	if err != nil {
		panic("File is not exist")
	}
	ext := filepath.Ext(imagePath)
	var img_list image.Image
	if ext == ".jpg" {
		img, err := jpeg.Decode(f)
		if err != nil {
			panic("File is not a jpg image")
		}
		img_list = img
	} else if ext == ".png" {
		img, err := png.Decode(f)
		if err != nil {
			panic("File is not a png image")
		}
		img_list = img
	}

	border := img_list.Bounds()
	list := [][]int{}
	dx := border.Dx()
	dy := border.Dy()

	for x := 0; x < dx; x++ {
		t_list := []int{}
		for y := 0; y < dy; y++ {
			_, g, _, _ := img_list.At(x, y).RGBA()
			ele := int(g >> 8) // 转为 RGB
			ele /= 255         // 二值化
			if ele == 1 {
				ele = 0
			} else {
				ele = 1
			}
			t_list = append(t_list, ele)
		}
		list = append(list, t_list)
	}
	return list
}

func image2Txt(imagePath string, toTxtPath string) {
	image_list := readGrayImage(imagePath)

	file, err := os.Create(toTxtPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	content := ""
	for _, i_list := range image_list {
		for _, ele := range i_list {
			content += strconv.Itoa(ele)
		}
		content += "\n"
	}
	_, err = file.WriteString(content)
	if err != nil {
		return
	}
}

func rowSum(maxtrix [][]int) []int {
	rowLength := len(maxtrix)
	colLength := len(maxtrix[0])

	ret := []int{}
	for i := 0; i < rowLength; i++ {
		rowS := 0
		for j := 0; j < colLength; j++ {
			rowS += maxtrix[i][j]
		}
		ret = append(ret, rowS)
	}
	return ret
}

func colSum(maxtrix [][]int) []int {
	rowLength := len(maxtrix)
	colLength := len(maxtrix[0])

	ret := []int{}
	for i := 0; i < colLength; i++ {
		colS := 0
		for j := 0; j < rowLength; j++ {
			colS += maxtrix[j][i]
		}
		ret = append(ret, colS)
	}
	return ret
}

func thresholdIndex(data []int, N int) (int, int) {
	length := len(data)
	start := 0
	end := length - 1

	for i := 0; i < length; i++ {
		if data[i] >= N {
			start = i
			break
		}
	}

	for i := length - 1; i >= 0; i-- {
		if data[i] >= N {
			end = i
			break
		}
	}
	return start, end
}

func saveImage(img image.Image, filename string) error {
	outFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outFile.Close()
	b := bufio.NewWriter(outFile)
	err = jpeg.Encode(b, img, nil)
	if err != nil {
		return err
	}
	err = b.Flush()
	if err != nil {
		return err
	}
	return nil
}

// resizeNearest is a fast nearest-neighbor resize, no filtering.
func resizeNearest(img image.Image, width, height int) *image.NRGBA {
	dst := image.NewNRGBA(image.Rect(0, 0, width, height))
	dx := float64(img.Bounds().Dx()) / float64(width)
	dy := float64(img.Bounds().Dy()) / float64(height)

	src := newScanner(img)
	parallel(0, height, func(ys <-chan int) {
		for y := range ys {
			srcY := int((float64(y) + 0.5) * dy)
			dstOff := y * dst.Stride
			for x := 0; x < width; x++ {
				srcX := int((float64(x) + 0.5) * dx)
				src.scan(srcX, srcY, srcX+1, srcY+1, dst.Pix[dstOff:dstOff+4])
				dstOff += 4
			}
		}
	})

	return dst
}
