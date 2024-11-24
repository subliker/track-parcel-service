package state

type Menu struct {
	Position MenuPosition
}

type MenuPosition uint

const (
	MenuPositionBase MenuPosition = iota
)
