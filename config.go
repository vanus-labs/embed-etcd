package embedetcd

type Config struct {
	Name                string   `yaml:"name"`
	DataDir             string   `yaml:"data_dir"`
	ListenClientAddr    string   `yaml:"listen_client_addr"`
	ListenPeerAddr      string   `yaml:"listen_peer_addr"`
	AdvertiseClientAddr string   `yaml:"advertise_client_addr"`
	AdvertisePeerAddr   string   `yaml:"advertise_peer_addr"`
	Clusters            []string `yaml:"clusters"`
}
