# txtai: AI-powered search engine for Go

[![Version](https://img.shields.io/github/release/neuml/txtai.go.svg?style=flat&color=success)](https://github.com/neuml/txtai.go/releases)
[![GitHub Release Date](https://img.shields.io/github/release-date/neuml/txtai.go.svg?style=flat&color=blue)](https://github.com/neuml/txtai.go/releases)
[![GitHub issues](https://img.shields.io/github/issues/neuml/txtai.go.svg?style=flat&color=success)](https://github.com/neuml/txtai.go/issues)
[![GitHub last commit](https://img.shields.io/github/last-commit/neuml/txtai.go.svg?style=flat&color=blue)](https://github.com/neuml/txtai.go)

[txtai](https://github.com/neuml/txtai) builds an AI-powered index over sections of text. txtai supports building text indices to perform similarity searches and create extractive question-answering based systems. txtai also has functionality for zero-shot classification.

This repository contains Go bindings for the txtai API. Full txtai functionality is supported.

## Installation
txtai.go can be installed as follows:

```bash
go get -u github.com/neuml/txtai.go
```

Alternatively, adding the following import within a module will also download txtai.go

```golang
import "github.com/neuml/txtai.go"
```

## Examples
The examples directory has a series of examples that give an overview of txtai. See the list of examples below.

| Example     |      Description      |
|:----------|:-------------|
| [Introducing txtai](https://github.com/neuml/txtai.go/blob/master/examples/embeddings.go) | Overview of the functionality provided by txtai |
| [Extractive QA with txtai](https://github.com/neuml/txtai.go/blob/master/examples/extractor.go) | Extractive question-answering with txtai |
| [Labeling with zero-shot classification](https://github.com/neuml/txtai.go/blob/master/examples/labels.go) | Labeling with zero-shot classification |

txtai.go connects to a txtai api instance. See [this link](https://github.com/neuml/txtai#api) for details on how to start a new api instance.

Once an api instance is running, do the following to run the examples.

```
git clone https://github.com/neuml/txtai.go
cd txtai.go/examples
make embeddings|extractor|labels
```
