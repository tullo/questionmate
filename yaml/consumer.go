package yaml

import "gopkg.in/yaml.v2"

type Consumer struct {
}

func (c Consumer) Consume(i interface{}, o interface{}) error {
	b := i.([]byte)
	return yaml.Unmarshal(b, o)
}
