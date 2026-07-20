package models

// ManifestSignature represents a cryptographic signature for a plugin manifest
type ManifestSignature struct {
	Algorithm string `json:"algorithm"`
	Hash      string `json:"hash"`
	Signature string `json:"signature"`
	Signer    string `json:"signer"`
}

// TrustedPublisher represents an entity allowed to sign plugins
type TrustedPublisher struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Fingerprint string `json:"fingerprint"`
}
