package main

type ImageFlyWeight struct {
	data string
}

func NewImageFlyWeight(filename string) *ImageFlyWeight {
	return &ImageFlyWeight{data: filename}
}

func (imw *ImageFlyWeight) GetData() string {
	return imw.data
}
