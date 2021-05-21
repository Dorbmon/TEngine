package TEngine

type HitPositon struct {
	X, Y int32
}

const (
	ButtonLeft = iota
	ButtonMid
	ButtonRight
	ButtonX1
	ButtonX2
)
const (
	MouseRelease = iota
	MousePress
)

type Hitable interface {
	OnHit(Position HitPositon, HitButton int, State int) bool
}
