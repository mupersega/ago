package vector

import (
	"fmt"
	"math"
)

// Vec2 represents a 2D vector with X and Y coordinates.
type Vec2 struct {
	X float64 // X component of the vector
	Y float64 // Y component of the vector
}

// Add adds two vectors and returns the resulting vector.
func (v Vec2) Add(v2 Vec2) Vec2 {
	return Vec2{v.X + v2.X, v.Y + v2.Y}
}

// Sub subtracts v2 from v and returns the resulting vector.
func (v Vec2) Sub(v2 Vec2) Vec2 {
	return Vec2{v.X - v2.X, v.Y - v2.Y}
}

// Mul multiplies the vector by a scalar and returns the resulting vector.
func (v Vec2) Mul(s float64) Vec2 {
	return Vec2{v.X * s, v.Y * s}
}

// Div divides the vector by a scalar and returns the resulting vector.
func (v Vec2) Div(s float64) Vec2 {
	return Vec2{v.X / s, v.Y / s}
}

// Mag calculates the magnitude (length) of the vector. It returns the square of the magnitude.
func (v Vec2) Mag() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Normalize returns a unit vector (a vector with magnitude 1) in the same direction.
func (v Vec2) Normalize() Vec2 {
	return v.Div(v.Mag())
}

// Dot calculates the dot product between two vectors. This measures how much one vector extends in the direction of another.
func (v Vec2) Dot(v2 Vec2) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

// Cross calculates the 2D cross product (which is a scalar in 2D). It can be used to determine the relative orientation of two vectors.
func (v Vec2) Cross(v2 Vec2) float64 {
	return v.X*v2.Y - v.Y*v2.X
}

// Angle returns the cosine of the angle between two vectors.
func (v Vec2) Angle(v2 Vec2) float64 {
	return v.Dot(v2) / (v.Mag() * v2.Mag())
}

// Rotate rotates the vector by the specified angle (in radians) and returns the resulting vector.
func (v Vec2) Rotate(angle float64) Vec2 {
	return Vec2{
		v.X*math.Cos(angle) - v.Y*math.Sin(angle),
		v.X*math.Sin(angle) + v.Y*math.Cos(angle),
	}
}

// Project projects vector v onto vector v2 and returns the resulting vector.
func (v Vec2) Project(v2 Vec2) Vec2 {
	return v2.Mul(v.Dot(v2) / v2.Mag())
}

// Reflect reflects vector v around vector v2 and returns the resulting vector. This is often used for calculating bounce directions.
func (v Vec2) Reflect(v2 Vec2) Vec2 {
	return v.Sub(v.Project(v2).Mul(2))
}

// Lerp linearly interpolates between vector v and vector v2 by a factor t (0 <= t <= 1).
func (v Vec2) Lerp(v2 Vec2, t float64) Vec2 {
	return v.Add(v2.Sub(v).Mul(t))
}

// Dist calculates the Euclidean distance between two vectors.
func (v Vec2) Dist(v2 Vec2) float64 {
	return v.Sub(v2).Mag()
}

// DistSq calculates the squared Euclidean distance between two vectors (more efficient when you don't need the actual distance).
func (v Vec2) DistSq(v2 Vec2) float64 {
	d := v.Sub(v2)
	return d.X*d.X + d.Y*d.Y
}

// AngleBetween returns the angle (in radians) between two vectors.
func (v Vec2) AngleBetween(v2 Vec2) float64 {
	return math.Acos(v.Dot(v2) / (v.Mag() * v2.Mag()))
}

// String formats the vector as a string for easy printing.
func (v Vec2) String() string {
	return fmt.Sprintf("Vec2{X: %v, Y: %v}", v.X, v.Y)
}

// Equals checks if two vectors are equal by comparing their X and Y components.
func (v Vec2) Equals(v2 Vec2) bool {
	return v.X == v2.X && v.Y == v2.Y
}

// Clone creates and returns a copy of the vector.
func (v Vec2) Clone() Vec2 {
	return Vec2{v.X, v.Y}
}
