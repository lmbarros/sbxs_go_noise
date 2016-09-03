// Package opensimplex provides an implementation of the OpenSimplex Noise
// algorithm in 2, 3 and 4 dimensions.
//
// OpenSimplex Noise generates "visually axis-decorrelated coherent noise",
// similar to Perlin's Simplex noise, but unencumbered by patents.
//
// The algorithm was created by Kurt Spencer (see his posts at
// http://uniblock.tumblr.com/post/97868843242/noise and
// http://uniblock.tumblr.com/post/99279694832/2d-and-4d-noise-too)
//
// This implementation is no more than a slight adaptation of the public domain
// Go code originally written by Owen Raccuglia (available at
// https://github.com/ojrac/opensimplex-go). This is not, however, a drop-in
// replacement of Owen's code.
package opensimplex
