package state

import (
	"errors"
	"fmt"

	"github.com/base-infrastructure/platform/internal/domain/models"
)

// Validator validates a StateManifest
type Validator struct{}

// NewValidator creates a new state manifest validator
func NewValidator() *Validator {
	return &Validator{}
}

// Validate checks a StateManifest for logical consistency
func (v *Validator) Validate(manifest *models.StateManifest) error {
	if manifest == nil {
		return errors.New("manifest is nil")
	}

	if manifest.Version == "" {
		return errors.New("manifest version is required")
	}

	if len(manifest.Capabilities) == 0 {
		return errors.New("manifest must contain at least one capability")
	}

	seenIDs := make(map[string]bool)
	for i, cap := range manifest.Capabilities {
		if cap.ID == "" {
			return fmt.Errorf("capability at index %d is missing an ID", i)
		}

		if cap.State != models.StateAvailable && cap.State != models.StateMissing {
			return fmt.Errorf("capability '%s' has invalid state: %s", cap.ID, cap.State)
		}

		if seenIDs[cap.ID] {
			return fmt.Errorf("duplicate capability ID found: %s", cap.ID)
		}
		seenIDs[cap.ID] = true
	}

	return nil
}
