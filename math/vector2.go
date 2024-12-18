package math

type Vector2[T Numeric] struct {
	X, Y int
}

type Vector2i = Vector2[int]
type Vector2f = Vector2[float64]

var (
	Up        = Vector2[int]{X: 0, Y: 1}
	Down      = Vector2[int]{X: 0, Y: -1}
	Left      = Vector2[int]{X: -1, Y: 0}
	Right     = Vector2[int]{X: 1, Y: 0}
	UpRight   = Vector2[int]{X: 1, Y: 1}
	DownRight = Vector2[int]{X: 1, Y: -1}
	UpLeft    = Vector2[int]{X: -1, Y: 1}
	DownLeft  = Vector2[int]{X: -1, Y: -1}
)

var OrthogonalDirections = []Vector2[int]{
	Up, Down, Left, Right,
}

var DiagonalDirections = []Vector2[int]{
	UpRight, DownRight, UpLeft, DownLeft,
}

func AddVector2[T Numeric](a, b Vector2[T]) Vector2[T] {
	return Vector2[T]{a.X + b.X, a.Y + b.Y}
}

func SubVector2[T Numeric](a, b Vector2[T]) Vector2[T] {
	return Vector2[T]{a.X - b.X, a.Y - b.Y}
}

func MultiplyScalarVector2[T Numeric](vector Vector2[T], scalar int) Vector2[T] {
	return Vector2[T]{vector.X * scalar, vector.Y * scalar}
}
