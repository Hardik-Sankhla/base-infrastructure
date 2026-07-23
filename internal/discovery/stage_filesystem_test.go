package discovery

import (
	"context"
	"log/slog"
	"testing"

	"github.com/base-infrastructure/platform/internal/core"
	"github.com/base-infrastructure/platform/internal/domain/models"
	"github.com/base-infrastructure/platform/internal/platform/mock"
)

func TestFilesystemStage_Success(t *testing.T) {
	p := mock.NewPlatform()
	p.MockFilesystem.Info = models.FilesystemInfo{
		Mounts: []models.MountPoint{
			{Device: "/dev/sda1", MountPath: "/", FSType: "ext4", Options: "rw", IsReadOnly: false},
		},
		RootCapacity: models.FilesystemCapacity{
			TotalBytes: 1000,
			UsedBytes:  500,
			FreeBytes:  500,
		},
		HomeDir: "/home/test",
	}

	stage := &FilesystemStage{}
	if stage.Name() != "filesystem" {
		t.Errorf("expected name 'filesystem', got %s", stage.Name())
	}

	ctx, cancel := context.WithTimeout(context.Background(), stage.Timeout())
	defer cancel()

	dctx := core.NewContext(slog.Default(), nil, nil, nil, p)

	if err := stage.Initialize(dctx); err != nil {
		t.Fatalf("Initialize failed: %v", err)
	}

	artifact, err := stage.Run(ctx, dctx)
	if err != nil {
		t.Fatalf("Run failed: %v", err)
	}

	if err := stage.Validate(artifact); err != nil {
		t.Fatalf("Validate failed: %v", err)
	}

	fsInfo, ok := artifact.(models.FilesystemInfo)
	if !ok {
		t.Fatalf("Expected FilesystemInfo artifact")
	}

	if len(fsInfo.Mounts) != 1 || fsInfo.Mounts[0].FSType != "ext4" {
		t.Errorf("Expected 1 mount of type ext4, got %+v", fsInfo.Mounts)
	}
	if fsInfo.HomeDir != "/home/test" {
		t.Errorf("Expected HomeDir /home/test, got %s", fsInfo.HomeDir)
	}
}
