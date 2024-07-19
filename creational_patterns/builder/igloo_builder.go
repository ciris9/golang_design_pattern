package main

type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (i *IglooBuilder) setWindowType() {
	i.windowType = "snow window"
}

func (i *IglooBuilder) setDoorType() {
	i.doorType = "snow door"
}

func (i *IglooBuilder) setNumFloor() {
	i.floor = 1
}

func (i *IglooBuilder) getHouse() House {
	return House{
		windowType: i.windowType,
		doorType:   i.doorType,
		floor:      i.floor,
	}
}
