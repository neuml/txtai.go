<p align="center">
    <img src="https://raw.githubusercontent.com/neuml/txtai/master/logo.png"/>
</p>

<p align="center">
    <b>Go client for txtai</b>
</p>

<p align="center">
    <a href="https://github.com/neuml/txtai.go/releases">
        <img src="https://img.shields.io/github/release/neuml/txtai.go.svg?style=flat&color=success" alt="Version"/>
    </a>
    <a href="https://github.com/neuml/txtai.go/releases">
        <img src="https://img.shields.io/github/release-date/neuml/txtai.go.svg?style=flat&color=blue" alt="GitHub Release Date"/>
    </a>
    <a href="https://github.com/neuml/txtai.go/issues">
        <img src="https://img.shields.io/github/issues/neuml/txtai.go.svg?style=flat&color=success" alt="GitHub Issues"/>
    </a>
    <a href="https://github.com/neuml/txtai.go">
        <img src="https://img.shields.io/github/last-commit/neuml/txtai.go.svg?style=flat&color=blue" alt="GitHub Last Commit"/>
    </a>
</p>

[txtai](https://github.com/neuml/txtai) is an all-in-one embeddings database for semantic search, LLM orchestration and language model workflows.

This repository contains Go bindings for the txtai API.

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
| [Pipelines and workflows](https://github.com/neuml/txtai.go/blob/master/examples/pipelines.go) | Pipelines and workflows |

txtai.go connects to a txtai api instance. See [this link](https://neuml.github.io/txtai/api/) for details on how to start a new api instance.

Once an api instance is running, do the following to run the examples.

```
git clone https://github.com/neuml/txtai.go
cd txtai.go/examples
make embeddings|extractor|labels|pipelines
```
