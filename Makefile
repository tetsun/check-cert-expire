NAME := check-cert-expire
VERSION := 0.1.0

.PHONY: \
	all \
	release_dir \
	linux64 \
	linux386 \
	darwin64 \
	darwin386 \
	windows64 \
	windows386 \
	source \
	clean

all: linux64 linux386 darwin64 darwin386 windows64 windows386

release_dir:
	mkdir -p ./release

linux64: release_dir
	$(eval GOOS := linux)
	$(eval GOARCH := amd64)
	$(eval TARGET := ${NAME}_${VERSION}_${GOOS}_${GOARCH})
	mkdir -p ./build/${TARGET}
	GOOS=${GOOS} GOARCH=${GOARCH} go build -o ./build/${TARGET}/${NAME}
	cd ./build && zip -r ${TARGET}.zip ${TARGET}
	mv ./build/${TARGET}.zip ./release/

linux386: release_dir
	$(eval GOOS := linux)
	$(eval GOARCH := 386)
	$(eval TARGET := ${NAME}_${VERSION}_${GOOS}_${GOARCH})
	mkdir -p ./build/${TARGET}
	GOOS=${GOOS} GOARCH=${GOARCH} go build -o ./build/${TARGET}/${NAME}
	cd ./build && zip -r ${TARGET}.zip ${TARGET}
	mv ./build/${TARGET}.zip ./release/

darwin64: release_dir
	$(eval GOOS := darwin)
	$(eval GOARCH := amd64)
	$(eval TARGET := ${NAME}_${VERSION}_${GOOS}_${GOARCH})
	mkdir -p ./build/${TARGET}
	GOOS=${GOOS} GOARCH=${GOARCH} go build -o ./build/${TARGET}/${NAME}
	cd ./build && zip -r ${TARGET}.zip ${TARGET}
	mv ./build/${TARGET}.zip ./release/

darwin386: release_dir
	$(eval GOOS := darwin)
	$(eval GOARCH := 386)
	$(eval TARGET := ${NAME}_${VERSION}_${GOOS}_${GOARCH})
	mkdir -p ./build/${TARGET}
	GOOS=${GOOS} GOARCH=${GOARCH} go build -o ./build/${TARGET}/${NAME}
	cd ./build && zip -r ${TARGET}.zip ${TARGET}
	mv ./build/${TARGET}.zip ./release/

windows64: release_dir
	$(eval GOOS := windows)
	$(eval GOARCH := amd64)
	$(eval TARGET := ${NAME}_${VERSION}_${GOOS}_${GOARCH})
	mkdir -p ./build/${TARGET}
	GOOS=${GOOS} GOARCH=${GOARCH} go build -o ./build/${TARGET}/${NAME}.exe
	cd ./build && zip -r ${TARGET}.zip ${TARGET}
	mv ./build/${TARGET}.zip ./release/

windows386: release_dir
	$(eval GOOS := windows)
	$(eval GOARCH := 386)
	$(eval TARGET := ${NAME}_${VERSION}_${GOOS}_${GOARCH})
	mkdir -p ./build/${TARGET}
	GOOS=${GOOS} GOARCH=${GOARCH} go build -o ./build/${TARGET}/${NAME}.exe
	cd ./build && zip -r ${TARGET}.zip ${TARGET}
	mv ./build/${TARGET}.zip ./release/

clean:
	go clean
	rm -rf ./build ./release
