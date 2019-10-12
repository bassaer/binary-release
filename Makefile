.PHONY: build run package clean

BASE = $(shell basename ${PWD})
VERSION = $(shell ./scripts/changelog.sh -v)
DIST = ${BASE}-${VERSION}

build:
	go build -o ${DIST} ./...

run:
	@./${DIST}

package:
	rm -rf ./release
	mkdir release
	tar zcvf release/${DIST}.tag.gz ${DIST}

clean:
	rm -f ./${BASE}-*
