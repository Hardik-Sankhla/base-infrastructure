package fakes

import (
	"context"
)

type FakeDownloader struct {
	SimulateError error
	Downloaded    map[string]string
}

func NewFakeDownloader() *FakeDownloader {
	return &FakeDownloader{
		Downloaded: make(map[string]string),
	}
}

func (d *FakeDownloader) Download(ctx context.Context, url string, dest string, expectedChecksum string) error {
	if d.SimulateError != nil {
		return d.SimulateError
	}
	d.Downloaded[url] = dest
	
	return nil
}
