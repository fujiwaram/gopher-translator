# Gopher Translator
[![Build Status](https://travis-ci.org/fujiwaram/gopher-translator.svg?branch=master)](https://travis-ci.org/golang/lint)

gopher translator for [goplay.space](https://goplay.space/)

## Installation

```
$ go get -u github.com/fujiwaram/gopher-translator
```

Install the docker and docker-compose, if you don't install yet.

```
$ make build-all
```

## Usage

```
$ make up
```
```
$ translate.sh TEXT

e.g. 
$ translate.sh ABC
```

Order for gopher is created as file in ./output directory.
