package runtime

import (
	"context"

	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Downloader interface {
	Download(ctx context.Context, url, dest, expectedChecksum string) error
}

type DefaultDownloader struct {
	client *http.Client
}

func NewDownloader() *DefaultDownloader {
	return &DefaultDownloader{
		client: &http.Client{},
	}
}

func (d *DefaultDownloader) Download(ctx context.Context, url, dest, expectedChecksum string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := d.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	file, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer file.Close()

	hasher := sha256.New()
	writer := io.MultiWriter(file, hasher)

	if _, err := io.Copy(writer, resp.Body); err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	if expectedChecksum != "" {
		actualChecksum := fmt.Sprintf("%x", hasher.Sum(nil))
		if actualChecksum != expectedChecksum {
			return fmt.Errorf("checksum mismatch: expected %s, got %s", expectedChecksum, actualChecksum)
		}
	}

	return nil
}
