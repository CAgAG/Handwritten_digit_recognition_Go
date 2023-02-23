package knn

import (
	"testing"
)

func TestKnn(t *testing.T) {
	type args struct {
		targetPath  string
		datasetPath string
		K           int
	}
	tests := []struct {
		name string
		args args
		want []knnData
	}{
		// TODO: Add test cases.
		{
			name: "01",
			args: args{targetPath: "./9_203.txt", datasetPath: "../trainDigits/", K: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := knn(tt.args.targetPath, tt.args.datasetPath, tt.args.K)

			t.Log(got)
		})
	}
}

func Test_saveRecImage(t *testing.T) {
	type args struct {
		path   string
		N      int
		toPath string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "01",
			args: args{path: "./p6.png", N: 6, toPath: "./test.jpg"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			saveRecImage(tt.args.path, tt.args.N, tt.args.toPath)
		})
	}
}

func TestPredictImage(t *testing.T) {
	type args struct {
		imagePath      string
		toPathDir      string
		datasetPathDir string
		K              int
	}
	tests := []struct {
		name string
		args args
		want []knnData
	}{
		// TODO: Add test cases.
		{
			name: "00",
			args: args{imagePath: "./test_data/p0.png", toPathDir: "./test_data/",
				datasetPathDir: "../trainDigits/", K: 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PredictImage(tt.args.imagePath, tt.args.toPathDir, tt.args.datasetPathDir, tt.args.K)

			t.Log(got)

		})
	}
}

func TestPredictTxt(t *testing.T) {
	type args struct {
		targetPath     string
		datasetPathDir string
		K              int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "00",
			args: args{targetPath: "./test_data/3_1.txt",
				datasetPathDir: "../trainDigits/", K: 6},
		},
		{
			name: "091",
			args: args{targetPath: "./test_data/9_203.txt",
				datasetPathDir: "../trainDigits/", K: 6},
		},
		{
			name: "092",
			args: args{targetPath: "./test_data/9_204.txt",
				datasetPathDir: "../trainDigits/", K: 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PredictTxt(tt.args.targetPath, tt.args.datasetPathDir, tt.args.K)

			t.Log(got)
		})
	}
}

func TestEvalTxtDir(t *testing.T) {
	type args struct {
		targetPathDir  string
		datasetPathDir string
		K              int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "01",
			args: args{targetPathDir: "../testDigits/", datasetPathDir: "../trainDigits/", K: 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EvalTxtDir(tt.args.targetPathDir, tt.args.datasetPathDir, tt.args.K)

			t.Log(got) // 0.982010582010582
		})
	}
}
