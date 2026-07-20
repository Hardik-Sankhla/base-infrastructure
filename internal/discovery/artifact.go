package discovery

// DiscoveryArtifact represents a strongly typed output from a discovery stage.
// All specific discovery domain models (e.g., Hardware, OSInfo) must
// implement this interface.
type DiscoveryArtifact interface {
	ArtifactType() string
}
