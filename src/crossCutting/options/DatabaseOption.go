package options

import "time"

type DatabaseOption struct {
	Port            int           `json:"port"`
	Url             string        `json:"url"`
	Username        string        `json:"username"`
	Password        string        `json:"password"`
	Database        string        `json:"database"`
	ConnMaxLifetime time.Duration `json:"set-conn-maxlifetime-seconds"`
}
