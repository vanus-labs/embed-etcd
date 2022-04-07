package embed_etcd

type Config struct {
	Name        string
	DataDir     string
	ClientAddrs string
	PeerAddrs   string
	//Security   SecurityConfig `toml:"security" json:"security"`
}

// SecurityConfig indicates the security configuration for pd server
type SecurityConfig struct {
	//grpcutil.TLSConfig
	// RedactInfoLog indicates that whether enabling redact log
	RedactInfoLog bool `toml:"redact-info-log" json:"redact-info-log"`
	//Encryption    encryption.Config `toml:"encryption" json:"encryption"`
}
