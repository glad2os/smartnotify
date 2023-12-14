package config

type Config struct {
	UI       UIConfig `yaml:"ui"`
	Settings Settings `yaml:"settings"`
	Users    []User   `yaml:"users"`
}

type UIConfig struct {
	ShowUI bool `yaml:"showUI"`
}

type Settings struct {
	Timezone string `yaml:"timezone"`
}

type User struct {
	Name          string         `yaml:"name"`
	Contact       Contact        `yaml:"contact"`
	Notifications []Notification `yaml:"notifications"`
}

type Contact struct {
	Method  string `yaml:"method"`
	Address string `yaml:"address"`
}

type Notification struct {
	Title   string  `yaml:"title"`
	Message string  `yaml:"message"`
	Trigger Trigger `yaml:"trigger"`
}

type Trigger struct {
	Type     string `yaml:"type"`
	Schedule string `yaml:"schedule"`
}
