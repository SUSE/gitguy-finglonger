package config

// Config structure for the github project
// TriagedColumn id in github
// BlockedColumn id in github
// InProgressColumn id in github
// ProgressColumn id, this is used only for PRs, in github
// DoneColumn id in github
// BacklogReleaseSquadColumn id in github
type Config struct {
	TriagedColumnID             int `yaml:"triagedColumnID,omitempty"`
	BlockedColumnID             int `yaml:"blockedColumnID,omitempty"`
	InProgressColumnID          int `yaml:"inProgressColumnID,omitempty"`
	ProgressColumnID            int `yaml:"progressColumnID,omitempty"`
	DoneColumnID                int `yaml:"doneColumnID,omitempty"`
	BacklogReleaseSquadColumnID int `yaml:"backlogReleaseSquadColumnID,omitempty"`
	Server                      struct {
		ServerAddr   string `yaml:"address,omitempty"`
		ReadTimeout  int    `yaml:"readTimeout,omitempty"`
		WriteTimeout int    `yaml:"writeTimeout,omitempty"`
	} `yaml:"server,omitempty"`
	Github struct {
		APIURL string `yaml:"apiURL"`
		Token  string `yaml:"token"`
		Secret string `yaml:"secret"`
	} `yaml:"github,omitempty"`
}
