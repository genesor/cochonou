# cochonou

[![GoDoc](https://godoc.org/github.com/genesor/cochonou?status.svg)](https://godoc.org/github.com/genesor/cochonou) [![Build Status](https://travis-ci.org/genesor/cochonou.svg?branch=master)](https://travis-ci.org/genesor/cochonou) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) [![Quay.io](https://quay.io/repository/genesor/cochonou/status)](https://quay.io/repository/genesor/cochonou)

Sub-domain generator for image listing written in Go

# Installation

## User

```bash
# Pull the latest version from quay.io
docker pull quay.io/genesor/cochonou
# Create & start the cochonou container from the pulled image
docker run -d -p 9494:9494 --name cochonou quay.io/genesor/cochonou
```

Your cochonou will be then accessible via [http://localhost:9494](http://localhost:9494)

## Contributor

You need to have [glide](https://github.com/Masterminds/glide) installed

If you don't have the github.com in your GOPATH/src
```bash
cd $GOPATH/src
mkdir github.com
```

Then follow this

```bash
# Create the genesor directory inside your GOPATH
cd $GOPATH/src/github.com
mkdir genesor
cd genesor
git clone git@github.com:genesor/cochonou.git
cd cochonou
# Install dependencies (glide etc.)
make install
# Launch the project with
make run
```

