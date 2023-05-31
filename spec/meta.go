package spec

type Meta struct {
	User        string        `yaml:"user,omitempty" default:"tidb"`
	TidbVersion string        `yaml:"tidb_version,omitempty" default:"0"`
	Topology    Specification `yaml:"topology,omitempty"`
}
