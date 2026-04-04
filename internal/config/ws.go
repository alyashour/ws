package config

import "path/filepath"

type Ws struct {
	ConfPath            string
	DefaultTaskFileName string
	RemoteKeyPath       string
}

func (ws *Ws) GetDataPath() string {
	return filepath.Join(ws.ConfPath, "data")
}

func (ws *Ws) GetTaskdataPath() string {
	return filepath.Join(ws.GetDataPath(), "tasks")
}

func (ws *Ws) GetDefaultTaskFilePath() string {
	return filepath.Join(ws.GetTaskdataPath(), ws.DefaultTaskFileName)
}
