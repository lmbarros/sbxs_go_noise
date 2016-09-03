package noise

// Noiser1D is something that can generate 1D noise.
type Noiser1D interface {
	Noise1D(x float64) float64
}

// Noiser2D is something that can generate 2D noise.
type Noiser2D interface {
	Noise2D(x, y float64) float64
}

// Noiser3D is something that can generate 3D noise.
type Noiser3D interface {
	Noise3D(x, y, z float64) float64
}

// Noiser4D is something that can generate 4D noise.
type Noiser4D interface {
	Noise4D(x, y, z, w float64) float64
}
