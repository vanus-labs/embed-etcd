package embedetcd

type Config struct {
	Name       string
	DataDir    string
	ClientAddr string
	PeerAddr   string
	Clusters   []string
}
