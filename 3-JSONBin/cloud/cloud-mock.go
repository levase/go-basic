package cloud

import (
	"encoding/json"
	"time"
)

type CloudStorage struct {
	url string
}

func NewCloudStorage(url string) *CloudStorage {
	return &CloudStorage{
		url: url,
	}
}

func (cs *CloudStorage) Read() ([]byte, error) {
	mockData := map[string]interface{}{
		"Bins": []map[string]any{
			{
				"id":        "mock1",
				"name":      "Mock Bin 1",
				"createdAt": time.Now(),
				"private":   false,
			},
			{
				"id":        "mock2",
				"name":      "Mock Bin 2",
				"createdAt": time.Now(),
				"private":   true,
			},
		},
	}
	return json.Marshal(mockData)
}

func (cs *CloudStorage) Write(a any) error {
	return nil
}
