package knn

import "path/filepath"

type knnData struct {
	file_path string
	score     float64
}

func (d *knnData) File_path() string {
	return d.file_path
}

func (d *knnData) SetFile_path(file_path string) {
	d.file_path = file_path
}

func (d *knnData) Score() float64 {
	return d.score
}

func (d *knnData) SetScore(score float64) {
	d.score = score
}

func (d *knnData) TrueValue() int {
	ret := filepath.Base(d.file_path)[0] - '0'
	return int(ret)
}
