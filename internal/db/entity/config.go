package entity

type Config struct {
	Id    int
	Value string
}

func (c Config) GetId() int {
	return c.Id
}

func (c Config) Table() string {
	return "configs"
}

func (c Config) ToJson() map[string]interface{} {
	return map[string]interface{}{
		"value": c.Value,
	}
}
