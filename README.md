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

You'll need to update the values inside `dev.env` to work correctly with your domain.

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

## Supported domains providers

Cochonou supports thw following providers:

* OVH

### Setup OVH

To use cochonou with OVH you need to create an OVH API App and a Token allowing Cochonou to manage Subdomain.

Create your app here [https://eu.api.ovh.com/createToken/](https://eu.api.ovh.com/createToken/) and allow the following access:

* **GET** /domain/*
* **POST** /domain/*
* **PUT** /domain/*
* **DELETE** /domain/*

You can choose the validity you want, but you'll have to update the value inside the project when the Token will expire and you'll need to create a new one.
