package project

// ResourceAccess ...
type ResourceAccess interface {
	GenerateProjectStructure(p *Project) (err error)
}
