package fractalnoise

import (
	"github.com/lmbarros/sbxs_go_noise"
)

// Params contains additional parameters passed to the fractal noise generator
// constructors.
//
// If any of these values is zero, a sensible default value is used instead:
// Layers (4), Frequency (1.0), Lacunarity (2.0), Amplitude (1.0), Gain (0.5).
type Params struct {
	Layers     int
	Frequency  float64
	Lacunarity float64
	Amplitude  float64
	Gain       float64
}

// generator1d is a generator of 1D fractal noise.
type generator1d struct {
	noiser noise.Noiser1D
	params Params
}

// generator2d is a generator of 2D fractal noise.
type generator2d struct {
	noiser noise.Noiser2D
	params Params
}

// generator3d is a generator of 3D fractal noise.
type generator3d struct {
	noiser noise.Noiser3D
	params Params
}

// generator4d is a generator of 4D fractal noise.
type generator4d struct {
	noiser noise.Noiser4D
	params Params
}

// New1D instantiates a one-dimensional fractal noise generator.
func New1D(noiser noise.Noiser1D, params Params) noise.Noiser1D {
	return &generator1d{
		noiser: noiser,
		params: newParams(params),
	}
}

// New2D instantiates a two-dimensional fractal noise generator.
func New2D(noiser noise.Noiser2D, params Params) noise.Noiser2D {
	return &generator2d{
		noiser: noiser,
		params: newParams(params),
	}
}

// New3D instantiates a three-dimensional fractal noise generator.
func New3D(noiser noise.Noiser3D, params Params) noise.Noiser3D {
	return &generator3d{
		noiser: noiser,
		params: newParams(params),
	}
}

// New4D instantiates a four-dimensional fractal noise generator.
func New4D(noiser noise.Noiser4D, params Params) noise.Noiser4D {
	return &generator4d{
		noiser: noiser,
		params: newParams(params),
	}
}

// Noise1D generates 1D fractal noise sampled at a given coordinate.
func (g *generator1d) Noise1D(x float64) float64 {
	sum := 0.0
	freq := g.params.Frequency
	amp := g.params.Amplitude

	for i := 0; i < g.params.Layers; i++ {
		d := float64(i) * 500.0 // to avoid "radial artifacts" around zero
		sum += g.noiser.Noise1D((x+d)*freq) * amp
		freq *= g.params.Lacunarity
		amp *= g.params.Gain
	}

	return sum
}

// Noise2D generates 2D fractal noise sampled at given coordinates.
func (g *generator2d) Noise2D(x, y float64) float64 {
	sum := 0.0
	freq := g.params.Frequency
	amp := g.params.Amplitude

	for i := 0; i < g.params.Layers; i++ {
		d := float64(i) * 500.0 // to avoid "radial artifacts" around zero
		sum += g.noiser.Noise2D((x+d)*freq, (y+d)*freq) * amp
		freq *= g.params.Lacunarity
		amp *= g.params.Gain
	}

	return sum
}

// Noise3D generates 3D fractal noise sampled at given coordinates.
func (g *generator3d) Noise3D(x, y, z float64) float64 {
	sum := 0.0
	freq := g.params.Frequency
	amp := g.params.Amplitude

	for i := 0; i < g.params.Layers; i++ {
		d := float64(i) * 500.0 // to avoid "radial artifacts" around zero
		sum += g.noiser.Noise3D((x+d)*freq, (y+d)*freq, (z+d)*freq) * amp
		freq *= g.params.Lacunarity
		amp *= g.params.Gain
	}

	return sum
}

// Noise4D generates 5D fractal noise sampled at given coordinates.
func (g *generator4d) Noise4D(x, y, z, w float64) float64 {
	sum := 0.0
	freq := g.params.Frequency
	amp := g.params.Amplitude

	for i := 0; i < g.params.Layers; i++ {
		d := float64(i) * 500.0 // to avoid "radial artifacts" around zero
		sum += g.noiser.Noise4D((x+d)*freq, (y+d)*freq, (z+d)*freq, (w+d)*freq) * amp
		freq *= g.params.Lacunarity
		amp *= g.params.Gain
	}

	return sum
}

// newParams returns a Params structure initialized either with the values in
// the Params passed as parameter, or with the default values.
func newParams(params Params) Params {
	result := Params{
		Layers:     4,
		Frequency:  1.0,
		Lacunarity: 2.0,
		Amplitude:  1.0,
		Gain:       0.5,
	}

	if params.Layers != 0 {
		result.Layers = params.Layers
	}

	if params.Frequency != 0.0 {
		result.Frequency = params.Frequency
	}

	if params.Lacunarity != 0.0 {
		result.Lacunarity = params.Lacunarity
	}

	if params.Amplitude != 0.0 {
		result.Amplitude = params.Amplitude
	}

	if params.Gain != 0.0 {
		result.Gain = params.Gain
	}

	return result
}
