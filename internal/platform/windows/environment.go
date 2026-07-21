package windows

import (
	"context"
	"os"
	"os/user"
	"strings"

	"github.com/base-infrastructure/platform/internal/domain/models"
)

type EnvironmentProvider struct{}

func NewEnvironmentProvider() *EnvironmentProvider {
	return &EnvironmentProvider{}
}

func (p *EnvironmentProvider) GetEnvironmentInfo(ctx context.Context) (models.EnvironmentInfo, error) {
	var info models.EnvironmentInfo

	// Terminal check
	if stat, err := os.Stdout.Stat(); err == nil {
		info.IsTerminal = (stat.Mode() & os.ModeCharDevice) != 0
	}

	// Root/Admin check (approximate without syscalls)
	if currentUser, err := user.Current(); err == nil {
		// A common hack in Go on windows is checking if we have access to open the PhysicalDrive0, 
		// but checking a basic heuristic is often enough for simple contexts.
		// We'll leave IsRoot=false by default unless we implement a dedicated syscall.
		_ = currentUser
	}

	// Container check
	// DOTNET_RUNNING_IN_CONTAINER is common for Windows containers
	if os.Getenv("DOTNET_RUNNING_IN_CONTAINER") == "true" {
		info.IsContainer = true
		info.ContainerRuntime = "docker" // Defaulting for Windows
	}

	// CI Check
	if os.Getenv("CI") != "" {
		info.IsCI = true
		if os.Getenv("GITHUB_ACTIONS") == "true" {
			info.CIProvider = "github"
		} else if strings.EqualFold(os.Getenv("GITLAB_CI"), "true") {
			info.CIProvider = "gitlab"
		} else {
			info.CIProvider = "unknown"
		}
	}

	return info, nil
}
