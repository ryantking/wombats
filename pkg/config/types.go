package config

// PackageConfig holds the global package information.
type PackageConfig struct {
	Name       string   `toml:"name"`
	Authors    []string `toml:"authors"`
	Version    string   `toml:"version"`
	License    string   `toml:"license"`
	EntryPoint string   `toml:"entry_point"`
	PatsccArgs []string `toml:"patscc_args"`
	CLibs      []string `toml:"c_libs"`
	GccArgs    []string `toml:"gcc_args"`
}

// Config is a struct representing the Wombats.yaml file that each project
// must contain.
type Config struct {
	Package      PackageConfig     `toml:"package"`
	Dependencies map[string]string `toml:"dependencies"`
}
