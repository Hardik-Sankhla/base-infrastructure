package fakes

import (
	"context"
)

type FakeDownloader struct {
	Downloaded map[string]string // url -> dest
	SimulateError   error
	ExpectedHash    string
}

func NewFakeDownloader() *FakeDownloader {
	return &FakeDownloader{
		Downloaded: make(map[string]string),
	}
}

func (d *FakeDownloader) Download(ctx context.Context, url, dest, expectedChecksum string) error {
	if d.SimulateError != nil {
		return d.SimulateError
	}
	d.Downloaded[url] = dest

	return nil
}
