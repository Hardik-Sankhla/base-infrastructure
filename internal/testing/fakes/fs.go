package fakes

import "sync"

// FakeFS implements fs.Manager for testing
type FakeFS struct {
	Files map[string][]byte
	mu    sync.RWMutex
}

func NewFakeFS() *FakeFS {
	return &FakeFS{
		Files: make(map[string][]byte),
	}
}

func (f *FakeFS) ConfigDir() string { return "/fake/config" }
func (f *FakeFS) CacheDir() string  { return "/fake/cache" }
func (f *FakeFS) DataDir() string   { return "/fake/data" }
func (f *FakeFS) TempDir() string   { return "/fake/tmp" }
func (f *FakeFS) PluginDir() string { return "/fake/plugins" }

func (f *FakeFS) AtomicWrite(path string, data []byte) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	cpy := make([]byte, len(data))
	copy(cpy, data)
	f.Files[path] = cpy
	return nil
}
