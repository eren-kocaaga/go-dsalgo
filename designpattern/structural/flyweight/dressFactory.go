package flyweight

import "fmt"

const (
	TerroristDressType        = "tDress"
	CounterTerroristDressType = "ctDress"
)

var (
	dressFactorySingleInstance = &DressFactory{
		dressMap: make(map[string]Dress),
	}
)

type DressFactory struct {
	dressMap map[string]Dress
}

func (d *DressFactory) getDressByType(dressType string) (Dress, error) {
	if d.dressMap[dressType] != nil {
		return d.dressMap[dressType], nil
	}

	switch dressType {
	case TerroristDressType:
		d.dressMap[dressType] = newTerroristDress()
	case CounterTerroristDressType:
		d.dressMap[dressType] = newCounterTerroristDress()
	default:
		return nil, fmt.Errorf("wrong dress type passed")
	}

	return d.dressMap[dressType], nil
}

func getSingleInstance() *DressFactory {
	return dressFactorySingleInstance
}
