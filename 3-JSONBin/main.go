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

	fmt.Println("Current bins:")
	for _, bin := range binList.Bins {
		fmt.Printf("- %s (ID: %s, Created: %s, Private: %v)\n",
			bin.Name, bin.ID, bin.CreatedAt.Format("2006-01-02 15:04:05"), bin.Private)
	}
}
