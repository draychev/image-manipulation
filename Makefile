#!make


.PHONY: test-annotate
test-annotate:
	go run ./main.go --annotation "hello world" --inputfile ./tests/pnggrad8rgb.png --outputfile ./output.png
