package client

type Flags struct {
	Key    *string
	Secret *string
}

type Config struct {
	flags *Flags
}

func NewConfig(f *Flags) *Config {
	return &Config{
		flags: f,
	}
}

func (c *Config) Name() (string, error) {
	return "Client", nil
}
