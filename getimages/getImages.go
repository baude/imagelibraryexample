package main

import (
	"fmt"
	"github.com/containers/storage"
	"github.com/projectatomic/libpod/libpod/image"
	"os"
)

func exitError(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func main() {
	// An empty StoreOptions will use defaults, which is /var/lib/containers/storage
	// If you define GraphRoot and RunRoot to a tempdir, it will create a new storage
	storeOpts := storage.StoreOptions{}
	// The first step to using the image library is almost always to create an
	// imageRuntime.
	imageRuntime, err := image.NewImageRuntimeFromOptions(storeOpts)
	if err != nil {
		exitError(err)
	}
	imageInputName := "busybox:glibc"
	// Now we can create a new image object for busybox:glibc which may or may not
	// already be present in the store
	_, err = imageRuntime.New(imageInputName, "",
		"", os.Stdout, nil, image.SigningOptions{})
	if err != nil {
		exitError(err)
	}

	images, err := imageRuntime.GetImages()
	if err != nil {
		exitError(err)
	}

	for _, i := range(images){
		fmt.Println(i.ID())
	}
}