package math

type Vector2i struct {
	X, Y int
}

var (
	Up        = Vector2i{X: 0, Y: 1}
	Down      = Vector2i{X: 0, Y: -1}
	Left      = Vector2i{X: -1, Y: 0}
	Right     = Vector2i{X: 1, Y: 0}
	UpRight   = Vector2i{X: 1, Y: 1}
	DownRight = Vector2i{X: 1, Y: -1}
	UpLeft    = Vector2i{X: -1, Y: 1}
	DownLeft  = Vector2i{X: -1, Y: -1}
)

var OrthogonalDirections = []Vector2i{
	Up, Down, Left, Right,
}

var DiagonalDirections = []Vector2i{
	UpRight, DownRight, UpLeft, DownLeft,
}

func AddVector2i(a, b Vector2i) Vector2i {
	return Vector2i{a.X + b.X, a.Y + b.Y}
}

func SubVector2i(a, b Vector2i) Vector2i {
	return Vector2i{a.X - b.X, a.Y - b.Y}
}

func MultiplyScalar(vector Vector2i, scalar int) Vector2i {
	return Vector2i{vector.X * scalar, vector.Y * scalar}
}
