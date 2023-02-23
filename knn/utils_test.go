package knn

import (
	"fmt"
	"image"
	"testing"
)

func Test_readFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{
			name: "01",
			args: args{path: "./9_203.txt"},
			want: []byte("00"),
		}, {
			name: "02",
			args: args{path: "../9_203.txt"},
			want: []byte("00"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := readFile(tt.args.path)
			t.Logf("\n%s", got)
		})
	}
}

func Test_toVector(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "01",
			args: args{content: readFile("./9_203.txt")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toVector(tt.args.content)

			t.Log(got)
			t.Log(len(got))
		})
	}
}

func Test_knnCompute(t *testing.T) {
	type args struct {
		a []int32
		b []int32
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "01",
			args: args{a: toVector(readFile("./9_203.txt")),
				b: toVector(readFile("./3_1.txt"))},
			want: 0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := knnCompute(tt.args.a, tt.args.b)

			t.Log(got)
		})
	}
}

func Test_walkPath(t *testing.T) {
	type args struct {
		file_dir string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "01",
			args: args{file_dir: "../trainDigits/"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := walkPath(tt.args.file_dir)
			fmt.Println(got)
		})
	}
}

func Test_readGrayImage(t *testing.T) {
	type args struct {
		imagePath string
	}
	tests := []struct {
		name string
		args args
		want image.Image
	}{
		// TODO: Add test cases.
		{
			name: "01",
			args: args{imagePath: "./p6.png"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := readGrayImage(tt.args.imagePath)

			fmt.Println(len(got))
			fmt.Println(got)
		})
	}
}
