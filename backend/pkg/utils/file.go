package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const FilesDir = "uploads/"

func DeleteFile(branch, filePath string) error {
	fullPath := filepath.Join(FilesDir, branch, filePath)

	if _, err := os.Stat(fullPath); err == nil {
		if err := os.Remove(fullPath); err != nil {
			return fmt.Errorf("failed to delete file %s: %w", fullPath, err)
		}
	} else if os.IsNotExist(err) {
		return nil
	} else {
		return fmt.Errorf("failed to check file %s: %w", fullPath, err)
	}

	return nil
}

func SaveFile(file io.Reader, branch, filePath string) error {
	fullPath := filepath.Join(FilesDir, branch)
	if err := os.MkdirAll(fullPath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", fullPath, err)
	}

	fullFilePath := filepath.Join(fullPath, filePath)
	outFile, err := os.Create(fullFilePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", fullFilePath, err)
	}
	defer outFile.Close()

	if _, err := io.Copy(outFile, file); err != nil {
		return fmt.Errorf("failed to copy file content to %s: %w", fullFilePath, err)
	}

	return nil
}
