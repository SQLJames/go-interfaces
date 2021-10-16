package main

import (
	"go-interfaces-demo/storage/driver/filesystem"
	"go-interfaces-demo/storage/driver"
)

func main()  {
	FS := filesystem.New()
	storageDriver := driver.New(FS)
	storageDriver.Insert("test",[]byte("data used for demo purposes"))
	storageDriver.Read("test")
}
