package config

type Config struct {
	Id    int    `json:"-"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (c Config) GetId() int {
	return c.Id
}

func (c Config) Table() string {
	return "configs"
}

func (c Config) ToJson() map[string]interface{} {
	return map[string]interface{}{
		"name":  c.Name,
		"value": c.Value,
	}
}
