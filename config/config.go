package config

import "gopkg.in/yaml.v3"

type Config struct {
	Database
	PublisherConfig
	SubscriberConfig
}

type file struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		Sslmode  string `yaml:"sslmode"`
	} `yaml:"database"`
	Nats struct {
		ClusterID string `yaml:"clusterID"`
		Publisher struct {
			ClientID string `yaml:"clientID"`
			Channel  string `yaml:"channel"`
		} `yaml:"publisher"`
		Subscriber struct {
			ClientID   string `yaml:"clientID"`
			Channel    string `yaml:"channel"`
			QueueGroup string `yaml:"queueGroup"`
		} `yaml:"subscriber"`
	} `yaml:"nats"`
}

func ParseConfig(fileBytes []byte) (*Config, error) {
	cf := file{}
	err := yaml.Unmarshal(fileBytes, &cf)

	if err != nil {
		return nil, err
	}

	return &Config{
		Database: Database{
			host:     cf.Database.Host,
			port:     cf.Database.Port,
			user:     cf.Database.User,
			password: cf.Database.Password,
			name:     cf.Database.Name,
			sslmode:  cf.Database.Sslmode,
		},
		PublisherConfig: PublisherConfig{
			ClusterID: cf.Nats.ClusterID,
			ClientID:  cf.Nats.Publisher.ClientID,
			Channel:   cf.Nats.Publisher.Channel,
		},
		SubscriberConfig: SubscriberConfig{
			ClusterID:  cf.Nats.ClusterID,
			ClientID:   cf.Nats.Subscriber.ClientID,
			Channel:    cf.Nats.Subscriber.Channel,
			QueueGroup: cf.Nats.Subscriber.QueueGroup,
		},
	}, nil
}
