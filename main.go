package main

import (
	"Handwritten_digit_recognition_Go/knn"
	"fmt"
)

func main() {
	// predict txt file
	pred := knn.PredictTxt("./testDigits/9_18.txt", "./trainDigits/", 6)
	fmt.Println(pred)

	// predict image
	pred = knn.PredictImage("./images/p0.png", "./images/", "./trainDigits/", 6)
	fmt.Println(pred)

	// evaluate knn AP
	ap := knn.EvalTxtDir("./testDigits/", "./trainDigits/", 6)
	fmt.Println(ap) //0.982010582010582
}
