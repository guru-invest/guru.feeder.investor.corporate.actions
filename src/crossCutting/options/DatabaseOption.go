package options

type DatabaseOption struct {
	Port     string `json:"port"`
	Url      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}
