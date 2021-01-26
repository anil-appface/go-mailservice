package model

// Config struct for webapp config
type Config struct {
	Server struct {
		// Host is the local machine IP Address to bind the HTTP Server to
		Host string `yaml:"host"`

		// Port is the local machine TCP Port to bind the HTTP Server to
		Port string `yaml:"port"`
	} `yaml:"server"`

	MailJet struct {
		PrivateKey string `yaml:"privateKey"`
		PublicKey  string `yaml:"publicKey"`
	} `yaml:"mailJet"`
}
