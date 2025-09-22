package slurm

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// InterLinkConfig holds the whole configuration
type SlurmConfig struct {
	VKConfigPath              string   `yaml:"VKConfigPath"`
	Sbatchpath                string   `yaml:"SbatchPath"`
	Scancelpath               string   `yaml:"ScancelPath"`
	Squeuepath                string   `yaml:"SqueuePath"`
	Sinfopath                 string   `yaml:"SinfoPath"`
	Sidecarport               string   `yaml:"SidecarPort"`
	Socket                    string   `yaml:"Socket"`
	ExportPodData             bool     `yaml:"ExportPodData"`
	Commandprefix             string   `yaml:"CommandPrefix"`
	ImagePrefix               string   `yaml:"ImagePrefix"`
	DataRootFolder            string   `yaml:"DataRootFolder"`
	Namespace                 string   `yaml:"Namespace"`
	Tsocks                    bool     `yaml:"Tsocks"`
	Tsockspath                string   `yaml:"TsocksPath"`
	Tsockslogin               string   `yaml:"TsocksLoginNode"`
	BashPath                  string   `yaml:"BashPath"`
	VerboseLogging            bool     `yaml:"VerboseLogging"`
	ErrorsOnlyLogging         bool     `yaml:"ErrorsOnlyLogging"`
	SingularityDefaultOptions []string `yaml:"SingularityDefaultOptions"`
	SingularityPrefix         string   `yaml:"SingularityPrefix"`
	SingularityPath           string   `yaml:"SingularityPath"`
	set                       bool
	EnrootDefaultOptions      []string `yaml:"EnrootDefaultOptions" default:"[\"--rw\"]"`
	EnrootPrefix              string   `yaml:"EnrootPrefix"`
	EnrootPath                string   `yaml:"EnrootPath"`
	ContainerRuntime          string   `yaml:"ContainerRuntime" default:"singularity"` // "singularity" or "enroot"
}

type CreateStruct struct {
	PodUID string `json:"PodUID"`
	PodJID string `json:"PodJID"`
}

type ContainerCommand struct {
	containerName    string
	isInitContainer  bool
	runtimeCommand   []string
	containerCommand []string
	containerArgs    []string
	containerImage   string
}

type Runtime interface {
	prepareCommand(config SlurmConfig, container v1.Container, metadata metav1.ObjectMeta)
}

type SingularityRuntime struct {
	ContainerCommand
}

type EnrootRuntime struct {
	ContainerCommand
}
