.PHONY: build run

default: build in/tt_api.json in/tt_api.map run

run:
	./run --build --api-file in/tt_api.json

build: .built snowcrash/.built

.built:
	docker build -t blueprint-parser .
	touch .built

snowcrash/.built:
	docker build -t snowcrash snowcrash
	touch .built

in/tt_api.%: snowcrash/.built
	(cd snowcrash && ./run -f json -o tt_api.json -s tt_api.map tt_api.apib)
	mv snowcrash/tt_api.json in
	mv snowcrash/tt_api.map in
