# go-mlkit — brand assets

Official logos for the **go-mlkit** organization.

<img src="avatar/go-mlkit.png" width="160" alt="go-mlkit">

The mark is a decision tree: a root that splits in two, each branch splitting again, with a leaf at the end of every path. Trees and forests are among the classifiers the library will provide. It is written by `mark/main.go`, which computes the layout from a
few numbers; nothing is placed by hand. Its gradient (`#F0D060` to `#7A5C00`)
is a gold that sets go-mlkit apart from the violet of go-numerics, the emerald of go-ndarray and go-fft, the blue of go-plot, the magenta of go-bignum, the brick red of go-sparse, the orange of go-approx, the rose of go-dsp and the cyan of go-symbolic; its dark end reads
on white at 6.25:1 and its light end on a dark page at 12.78:1.

## Formats

- **SVG** — `svg/color`, `svg/white` (white, transparent), `svg/black` (black, transparent)
- **PNG** 16→1024 px, 3 variants — `png/<variant>/<size>/`
- **JPG** (color, white background) — `jpg/<size>/`
- **ICO** (Windows) — `ico/`
- **ICNS** (macOS) — `icns/`
- **GitHub avatar** 512 px — `avatar/`
- **Social preview** 1280×640 (repo banner) — `social/`

`check/kit.py` counts the whole matrix on every pull request, and the `mark`
job runs `mark/main.go` again and fails if `svg/color/go-mlkit.svg` differs.

---
*Generated from `svg/color/go-mlkit.svg` by the fleet's Go brand pipeline (brandgen) — rounded square badge, line glyph, color / white / black variants.*
