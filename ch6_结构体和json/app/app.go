package app

type App struct {
	AppName    string
	AppVersion string `json:"app_version"`
	appAuthor  string "pleuvoir"
	DefaultD   string "default"
}
