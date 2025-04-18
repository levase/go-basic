package file

import (
	"encoding/json"
	"errors"
	"fmt"
	"homework/JSONBin/bins"
	"os"
	"path/filepath"
	"strings"
)

var (
	ErrNotJSONFile = errors.New("file is not a JSON file")
)

func ReadFile(path string) (bins.BinList, error) {
	if !isJSONFile(path) {
		return nil, fmt.Errorf("%w: '%s'", ErrNotJSONFile, path)
	}
       if _, err := os.Stat(path); err != nil {
        if os.IsNotExist(err) {
            return nil, fmt.Errorf("file '%s' does not exist: %w", path, err)
        }
        return nil, fmt.Errorf("failed to check file '%s': %w", path, err)
    }


	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", path, err)
	}

	if len(data) == 0 {
		return nil, errors.New("file is empty")
	}

	var binList bins.BinList
	if err := json.Unmarshal(data, &binList); err != nil {
		return nil, fmt.Errorf("failed to unmarshal data: %w", err)
	}
	return binList, nil
}

func isJSONFile(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	return ext == ".json"
}
