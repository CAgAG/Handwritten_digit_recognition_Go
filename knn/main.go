package knn

import (
	"image"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func knn(targetPath string, datasetPathDir string, K int) []knnData {
	x := toVector(readFile(targetPath))
	data_list := []knnData{}

	for _, path := range walkPath(datasetPathDir) {
		data_x := toVector(readFile(path))
		score := knnCompute(x, data_x)

		data := knnData{file_path: path, score: score}
		data_list = append(data_list, data)
	}
	sort.SliceStable(data_list, func(i, j int) bool {
		return data_list[i].score < data_list[j].score
	})
	return data_list[:K]
}

func PredictTxt(targetPath string, datasetPathDir string, K int) int {
	data_list := knn(targetPath, datasetPathDir, K)

	//fmt.Println(data_list)
	count := make([]int, 10)
	for _, data := range data_list {
		count[data.TrueValue()] += 1
	}

	maxValue := 0
	maxIndex := 0
	for i, c := range count {
		if c >= maxValue {
			maxValue = c
			maxIndex = i
		}
	}
	return maxIndex
}

func EvalTxtDir(targetPathDir string, datasetPathDir string, K int) float64 {
	ap := 0.0
	total := 0.0
	for _, targetPath := range walkPath(targetPathDir) {
		data_list := knn(targetPath, datasetPathDir, K)

		count := make([]int, 10)
		for _, data := range data_list {
			count[data.TrueValue()] += 1
		}

		maxValue := 0
		maxIndex := 0
		for i, c := range count {
			if c >= maxValue {
				maxValue = c
				maxIndex = i
			}
		}
		predValue := maxIndex
		filename := strings.Split(filepath.Base(targetPath), ".")[0]
		trueValue, _ := strconv.Atoi(strings.Split(filename, "_")[0])

		if predValue == trueValue {
			ap += 1.0
		}
		total += 1.0
	}
	ap /= total
	return ap
}

func PredictImage(imagePath string, toPathDir string, datasetPathDir string, K int) int {
	filename := strings.Split(filepath.Base(imagePath), ".")[0]
	targetPath := saveRecImage(imagePath, 3, filepath.Join(toPathDir, filename+".jpg"))
	data_list := knn(targetPath, datasetPathDir, K)

	//fmt.Println(data_list)
	count := make([]int, 10)
	for _, data := range data_list {
		count[data.TrueValue()] += 1
	}

	maxValue := 0
	maxIndex := 0
	for i, c := range count {
		if c >= maxValue {
			maxValue = c
			maxIndex = i
		}
	}
	return maxIndex
}

func saveRecImage(path string, N int, toPath string) string {
	input, _ := os.Open(path)
	img, _, err := image.Decode(input)
	img = img.(image.Image)

	if err != nil {
		panic(err)
	}

	image_data := readGrayImage(path)
	rowS := rowSum(image_data)
	colS := colSum(image_data)

	row_min, row_max := thresholdIndex(rowS, N)
	col_min, col_max := thresholdIndex(colS, N)

	window := image.Rect(
		row_min, col_min,
		row_max, col_max)
	dst := imageCrop(img, window)
	dst = resizeNearest(dst, 32, 32)
	err = saveImage(dst, toPath)
	if err != nil {
		panic("save image fail")
	}
	filename := strings.Split(filepath.Base(toPath), ".")[0]
	txtPath := filepath.Join(filepath.Dir(toPath), filename+".txt")
	image2Txt(toPath, txtPath)
	return txtPath
}
