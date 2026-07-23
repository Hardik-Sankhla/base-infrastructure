package state

import (
	"fmt"
	"io"
	"os"

	"github.com/base-infrastructure/platform/internal/domain/models"
	"gopkg.in/yaml.v3"
)

// Parser handles the loading and unmarshaling of StateManifests
type Parser struct{}

// NewParser creates a new state manifest parser
func NewParser() *Parser {
	return &Parser{}
}

// Parse parses a byte slice into a StateManifest
func (p *Parser) Parse(data []byte) (*models.StateManifest, error) {
	var manifest models.StateManifest
	if err := yaml.Unmarshal(data, &manifest); err != nil {
		return nil, fmt.Errorf("failed to parse yaml: %w", err)
	}
	return &manifest, nil
}

// Load reads a state manifest from a file path
func (p *Parser) Load(filepath string) (*models.StateManifest, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open state file: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read state file: %w", err)
	}

	return p.Parse(data)
}
