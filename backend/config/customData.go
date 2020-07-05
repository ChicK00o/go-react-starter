package config

type CustomData struct {
	Port        int    `json:"port"`
}

func InitialCustomData() *CustomData {
	return &CustomData{
		Port:        4999,
	}
}
