package registry

import "errors"

var integrations map[string][]Assertion = map[string][]Assertion{}

type Assertion interface {
	Process()
}

func AddIntegrations(key string, i ...Assertion)  {
	arr, ok := integrations[key]
	if !ok {
		arr = []Assertion{}
	}
	arr = append(arr, i...)

	integrations[key] = arr

}

func GetIntegrations(key string) ([]Assertion, error)  {
	arr, ok := integrations[key]
	if !ok {
		return nil, errors.New("Assertion not found")
	}

	return arr, nil
}


