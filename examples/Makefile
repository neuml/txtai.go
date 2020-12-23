.PHONY: build

# Setup environment
export BUILD_DIR := ./build

# Create build directory
build:
	mkdir -p ${BUILD_DIR}

# Embeddings test
embeddings: build
	go build -o ${BUILD_DIR} embeddings.go
	${BUILD_DIR}/embeddings

# Extractor test
extractor: build
	go build -o ${BUILD_DIR} extractor.go
	${BUILD_DIR}/extractor

# Labels test
labels: build
	go build -o ${BUILD_DIR} labels.go
	${BUILD_DIR}/labels
