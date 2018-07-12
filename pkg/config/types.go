package config

// PackageConfig holds the global package information.
type PackageConfig struct {
	Name    string
	Authors []string
	Version string
	License string
}

// DependencyConfig holds information about a dependency.
type DependencyConfig struct {
	Version string
	Source  string
}

// Config is a struct represention of the Wombats.yaml file that each project
// must contain.
type Config struct {
	Package      PackageConfig
	Dependencies []DependencyConfig
}
