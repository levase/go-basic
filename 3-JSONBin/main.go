package main

import (
	"fmt"
	"homework/JSONBin/bins"
	"homework/JSONBin/config"
	"homework/JSONBin/file"
	"homework/JSONBin/storage"
	"log"
)

func main() {
	if err := config.LoadEnv(".env"); err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("Failed to create config: ", err)
	}
	fmt.Println(cfg)

	fileStorage, err := file.NewFileStorage("./output/local_bin_storage.json")
	if err != nil {
		log.Fatalf("Failed to create file storage: %v", err)
	}

	jsonStorage := storage.NewStorage(fileStorage)

	// cloudStorage := cloud.NewCloudStorage("https://example.com")
	// jsonStorage := storage.NewStorage(cloudStorage)

	binList, err := jsonStorage.LoadBins()
	if err != nil {
		log.Printf("Failed to load bins: %v", err)
	}

	// Тестовое создание и добавление нового bin
	newBin, err := bins.NewBin("1", "Test bin", false)
	if err != nil {
		log.Fatalf("Failed to create bin: %v", err)
	}

	if err := jsonStorage.AddBin(newBin); err != nil {
		log.Fatalf("Failed to add bin: %v", err)
	}

	// Тестовое добавление второго bin
	secondBin, err := bins.NewBin("2", "Test bin 2", true)
	if err != nil {
		log.Fatalf("Failed to create second bin: %v", err)
	}

	if err := jsonStorage.AddBin(secondBin); err != nil {
		log.Fatalf("Failed to add second bin: %v", err)
	}

	fmt.Println("Current bins:")
	for _, bin := range binList.Bins {
		fmt.Printf("- %s (ID: %s, Created: %s, Private: %v)\n",
			bin.Name, bin.ID, bin.CreatedAt.Format("2006-01-02 15:04:05"), bin.Private)
	}
}
