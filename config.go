package embedetcd

type Config struct {
	Name       string   `yaml:"name"`
	DataDir    string   `yaml:"data_dir"`
	ClientAddr string   `yaml:"client_addr"`
	PeerAddr   string   `yaml:"peer_addr"`
	Clusters   []string `yaml:"clusters"`
}
