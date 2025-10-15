package ast

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadDocumentFile(filePath string) (Document, error) {
	var result Document
	fileHandle, openErr := os.Open(filePath)
	if openErr != nil {
		return result, fmt.Errorf("unable to open AST document %s: %w", filePath, openErr)
	}
	defer fileHandle.Close()
	decoder := json.NewDecoder(fileHandle)
	decodeErr := decoder.Decode(&result)
	if decodeErr != nil {
		return result, fmt.Errorf("unable to decode AST document %s: %w", filePath, decodeErr)
	}
	return result, nil
}

func LoadDocumentBytes(payload []byte) (Document, error) {
	var result Document
	decodeErr := json.Unmarshal(payload, &result)
	if decodeErr != nil {
		return result, fmt.Errorf("unable to decode AST payload: %w", decodeErr)
	}
	return result, nil
}
