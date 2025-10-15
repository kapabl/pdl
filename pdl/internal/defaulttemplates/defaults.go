package defaulttemplates

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

const embeddedRoot = "assets/default"

//go:embed assets/default/**
var assets embed.FS

func MaterializeInto(root string) (string, error) {
	target := filepath.Join(root, ".pdl", "templates", "default")
	walkErr := fs.WalkDir(assets, embeddedRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, relErr := filepath.Rel(embeddedRoot, path)
		if relErr != nil {
			return relErr
		}
		if rel == "." {
			return os.MkdirAll(target, 0o755)
		}
		destination := filepath.Join(target, rel)
		if d.IsDir() {
			return os.MkdirAll(destination, 0o755)
		}
		content, readErr := fs.ReadFile(assets, path)
		if readErr != nil {
			return readErr
		}
		if strings.HasSuffix(destination, ".gotmpl") {
			destination = strings.TrimSuffix(destination, ".gotmpl")
		}
		if err := os.MkdirAll(filepath.Dir(destination), 0o755); err != nil {
			return err
		}
		return os.WriteFile(destination, content, 0o644)
	})
	if walkErr != nil {
		return "", walkErr
	}
	return target, nil
}
