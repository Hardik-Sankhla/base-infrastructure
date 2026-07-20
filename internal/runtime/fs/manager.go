package fs

import (
	"fmt"
	"os"
	"path/filepath"
)

// Manager defines the interface for filesystem operations
type Manager interface {
	ConfigDir() string
	CacheDir() string
	DataDir() string
	TempDir() string
	PluginDir() string
	AtomicWrite(path string, data []byte) error
}

type DefaultManager struct {
	baseDir string
}

func NewManager(baseDir string) *DefaultManager {
	return &DefaultManager{baseDir: baseDir}
}

func (m *DefaultManager) ConfigDir() string {
	return filepath.Join(m.baseDir, "config")
}

func (m *DefaultManager) CacheDir() string {
	return filepath.Join(m.baseDir, "cache")
}

func (m *DefaultManager) DataDir() string {
	return filepath.Join(m.baseDir, "data")
}

func (m *DefaultManager) TempDir() string {
	return filepath.Join(m.baseDir, "tmp")
}

func (m *DefaultManager) PluginDir() string {
	return filepath.Join(m.baseDir, "plugins")
}

// AtomicWrite writes data to a temporary file and renames it to target path
func (m *DefaultManager) AtomicWrite(path string, data []byte) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	tmpFile, err := os.CreateTemp(dir, "atomic-*")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	tmpPath := tmpFile.Name()

	defer func() {
		tmpFile.Close()
		os.Remove(tmpPath)
	}()

	if _, err := tmpFile.Write(data); err != nil {
		return fmt.Errorf("failed to write data: %w", err)
	}

	if err := tmpFile.Sync(); err != nil {
		return fmt.Errorf("failed to sync data: %w", err)
	}

	if err := tmpFile.Close(); err != nil {
		return fmt.Errorf("failed to close temp file: %w", err)
	}

	return os.Rename(tmpPath, path)
}
