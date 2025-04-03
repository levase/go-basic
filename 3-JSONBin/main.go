package main

import (
	"homework/JSONBin/bins"
	"homework/JSONBin/file"
	"homework/JSONBin/storage"
	"log"
)

type Config struct {
	StoragePath string
	ImportPath  string
}

func main() {
	cfg := Config{
		StoragePath: "./output/data.json",
		ImportPath:  "./output/somefile.json",
	}

	storage, err := storage.NewStorage(cfg.StoragePath)
	if err != nil {
		log.Fatalf("Failed to initialize storage: %v", err)
	}

	binList, err := storage.ReadJSON()
	if err != nil {
		log.Fatalf("Failed to read bins: %v", err)
	}

	data, err := file.ReadFile(cfg.ImportPath)
	if err != nil {
		log.Println(err)
	}

	if len(data) > 0 {
		binList = append(binList, data...)
	}

	// Test bin. Can be deleted
	bin, err := bins.NewBin("5", "John", false)
	if err != nil {
		log.Fatalf("Failed to create bin: %v", err)
	}

	binList = append(binList, bin)
	if err := storage.SaveJSON(&binList); err != nil {
		log.Fatalf("Failed to save bins: %v", err)
	}
}
