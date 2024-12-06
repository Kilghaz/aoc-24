package math

type Vector2i struct {
	X, Y int
}

func AddVector2i(a, b Vector2i) Vector2i {
	return Vector2i{a.X + b.X, a.Y + b.Y}
}
