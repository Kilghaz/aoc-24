package math

type Vector2i struct {
	X, Y int
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
