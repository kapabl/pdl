package filetimes

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type FileTimes struct {
	filename string
	entries  map[string]int64
}

func NewFileTimes(filename string) FileTimes {
	var result FileTimes
	result.filename = filename
	result.entries = make(map[string]int64)
	return result
}

func (registry *FileTimes) Read() error {
	_, err := os.Stat(registry.filename)
	if err != nil {
		return nil
	}
	bytes, readErr := os.ReadFile(registry.filename)
	if readErr != nil {
		return readErr
	}
	decodeErr := json.Unmarshal(bytes, &registry.entries)
	if decodeErr != nil {
		registry.entries = make(map[string]int64)
		return decodeErr
	}
	return nil
}

func (registry *FileTimes) Write() error {
	dir := filepath.Dir(registry.filename)
	if dir != "" {
		os.MkdirAll(dir, 0o755)
	}
	payload, encodeErr := json.Marshal(registry.entries)
	if encodeErr != nil {
		return encodeErr
	}
	return os.WriteFile(registry.filename, payload, 0o644)
}

func (registry *FileTimes) AddFile(file string) error {
	info, statErr := os.Stat(file)
	if statErr != nil {
		return statErr
	}
	registry.entries[file] = info.ModTime().UnixMilli()
	return nil
}

func (registry *FileTimes) IsFileModified(file string) (bool, error) {
	info, statErr := os.Stat(file)
	if statErr != nil {
		return false, statErr
	}
	stored, ok := registry.entries[file]
	if !ok {
		return true, nil
	}
	result := info.ModTime().UnixMilli() != stored
	return result, nil
}

func (registry *FileTimes) Touch(file string) error {
	return registry.AddFile(file)
}

func CurrentTimestamp() int64 {
	result := time.Now().UnixMilli()
	return result
}
