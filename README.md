# sbxs_go_noise

Noise in Go.

Package `noise` provides basic noise-related types used elsewhere, like the
`Noiser1D`, `Noiser2D`, `Noiser3D` and `Noiser4D` interfaces

Package `opensimplex` is an implementation of the
[OpenSimplex Noise](http://uniblock.tumblr.com/post/97868843242/noise)
algorithm. The implementation is mostly a shameless copy from Owen Raccuglia's
[opensimplex-go](https://github.com/ojrac/opensimplex-go).

Package `fractalnoise` provides a means to combine different layers of noise in
order to produce fractal-like noise. You may know this by fBm, Fractional
Brownian Motion or Fractal Sum.

## License

All code here is under the MIT License.
