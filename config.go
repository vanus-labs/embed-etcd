package embedetcd

type Config struct {
	Name                string            `yaml:"name"`
	DataDir             string            `yaml:"data_dir"`
	ListenClientAddr    string            `yaml:"listen_client_addr"`
	ListenPeerAddr      string            `yaml:"listen_peer_addr"`
	AdvertiseClientAddr string            `yaml:"advertise_client_addr"`
	AdvertisePeerAddr   string            `yaml:"advertise_peer_addr"`
	Clusters            map[string]string `yaml:"clusters"`
	TLSConfig           TLSConfig         `yaml:"tls_config"`
}

type TLSConfig struct {
	CertFile       string `yaml:"cert_file"`
	KeyFile        string `yaml:"key_file"`
	ClientCertFile string `yaml:"client_cert_file"`
	ClientKeyFile  string `yaml:"client_key_file"`
	TrustedCAFile  string `yaml:"trusted_ca_file"`
}
