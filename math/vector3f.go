package math

import "math"

type Numeric interface {
	~int | ~uint8 | ~uint16 | ~uint32 | ~float32 | ~float64
}

type Vector3[T Numeric] struct {
	X, Y, Z T
}

func AddVector3[T Numeric](a, b Vector3[T]) Vector3[T] {
	return Vector3[T]{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

func SubVector3[T Numeric](a, b Vector3[T]) Vector3[T] {
	return Vector3[T]{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

func MultiplyScalarVector3[T Numeric](vector Vector3[T], scalar T) Vector3[T] {
	return Vector3[T]{vector.X * scalar, vector.Y * scalar, vector.Z * scalar}
}

func MagnitudeVector3[T Numeric](vector Vector3[T]) T {
	return T(math.Sqrt(math.Pow(float64(vector.X), 2) + math.Pow(float64(vector.Y), 2) + math.Pow(float64(vector.Z), 2)))
}
