package project

// Repository ...
type Repository interface {
	GenerateProjectStructure(p *Project) (err error)
}
