package main

import "sync"

type ImageFlyWeightFactory struct {
	maps map[string]*ImageFlyWeight
}

var imageFlyWeightFactory *ImageFlyWeightFactory
var mux sync.Once

func NewImageFlyWeightFactory() *ImageFlyWeightFactory {
	mux.Do(func() {
		imageFlyWeightFactory = &ImageFlyWeightFactory{maps: make(map[string]*ImageFlyWeight)}
	})
	return imageFlyWeightFactory
}

func (imf *ImageFlyWeightFactory) Get(filename string) *ImageFlyWeight {
	image := imf.maps[filename]
	if image == nil {
		image = NewImageFlyWeight(filename)
		imf.maps[filename] = image
	}
	return image
}

func (imf *ImageFlyWeightFactory) GetMapCount() int {
	return len(imf.maps)
}
