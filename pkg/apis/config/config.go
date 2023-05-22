package config

type Config struct {
	Debug             bool
	Port              int
	KubeConfig        string
	NameSpace         string
	ClusterInfos      string
	LocalClusterInfos string
	BasicAuthUser     string
	BasicAuthPassword string
}
