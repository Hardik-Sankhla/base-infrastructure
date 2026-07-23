package runtime

import (
	"fmt"
	"os"
	"path/filepath"
)

// FSManager defines the interface for filesystem operations
type FSManager interface {
	ConfigDir() string
	CacheDir() string
	DataDir() string
	TempDir() string
	PluginDir() string
	AtomicWrite(path string, data []byte) error
}

type DefaultFSManager struct {
	baseDir string
}

func NewFSManager(baseDir string) *DefaultFSManager {
	return &DefaultFSManager{baseDir: baseDir}
}

func (m *DefaultFSManager) ConfigDir() string {
	return filepath.Join(m.baseDir, "config")
}

func (m *DefaultFSManager) CacheDir() string {
	return filepath.Join(m.baseDir, "cache")
}

func (m *DefaultFSManager) DataDir() string {
	return filepath.Join(m.baseDir, "data")
}

func (m *DefaultFSManager) TempDir() string {
	return filepath.Join(m.baseDir, "tmp")
}

func (m *DefaultFSManager) PluginDir() string {
	return filepath.Join(m.baseDir, "plugins")
}

// AtomicWrite writes data to a temporary file and renames it to target path
func (m *DefaultFSManager) AtomicWrite(path string, data []byte) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
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
