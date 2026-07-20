package platform

// BasePlatform provides default stub implementations for all Platform interface methods.
// OS-specific platform structs can embed this to avoid having to implement every provider
// if they don't support it yet.
type BasePlatform struct{}

func (b *BasePlatform) Hardware() HardwareProvider       { return nil }
func (b *BasePlatform) OS() OSProvider                   { return nil }
func (b *BasePlatform) Filesystem() FilesystemProvider   { return nil }
func (b *BasePlatform) Network() NetworkProvider         { return nil }
func (b *BasePlatform) Environment() EnvironmentProvider { return nil }
func (b *BasePlatform) Software() SoftwareProvider       { return nil }
func (b *BasePlatform) Process() ProcessProvider         { return nil }
func (b *BasePlatform) Security() SecurityProvider       { return nil }
func (b *BasePlatform) User() UserProvider               { return nil }
func (b *BasePlatform) Service() ServiceProvider         { return nil }
