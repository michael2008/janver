[![GoDoc](https://godoc.org/github.com/djhaskin987/janver?status.svg)](https://godoc.org/github.com/djhaskin987/janver)

Provides many different popular version comparison algorithms as string
comparisons. See the docs for more information.

*Another* version comparison library, you ask? Well, yes.

Why? Because there are lots of version comparison schemes out there, far more
than just SemVer. This library supports several:

  - Debian package version comparison
  - Maven (`org.apache.maven.artifact.versioning.ComparableVersion`) version
    comparison
  - Naive (old RPM) version comparison
  - Python (PEP 440) version comparison
  - RPM version comparison
  - Ruby (`Gem::Version`) version comparison
  - SemVer 2.0 version comparison

It supports so many by normalizing the version number according what
scheme is being used, then feeds the newly transformed version numbers to the
venerable debian version comparison algorithm for comparison.

## Jan?

[Johannes Vermeer (JanVer)](https://en.wikipedia.org/wiki/Johannes_Vermeer) is the inspiration for this library's name.

## Usage

To import:

```golang
import (
	"gitlab.com/djhaskin987/janver/pkg/vercmp"
)
```

To use:

```golang
vercmp.MavenCompare(avers, bvers)
```

You'll get a negative, positive, or zero value depending on whether
`avers < bvers`, `avers > bvers`, or `avers = bvers`, respectively. :)

If this does not work as advertised, please log a bug :)

## License

Copyright © 2017 Daniel Jay Haskin et. al., see the AUTHORS.md file.

Distributed under the MIT License, see the LICENSE file.
