package config

const (
	ProjectKey        = "go-react-starter"
	TestTableCreation = false
)

type CustomData struct {
	Port int `json:"port"`
}

func InitialCustomData() *CustomData {
	return &CustomData{
		Port: 4999,
	}
}
