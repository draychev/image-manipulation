package annotate

import (
	"bytes"
	"os"
	"testing"
)

func TestDrawString(t *testing.T) {
	inputFile := "test-input.png"
	outputFile := "test-output.png"
	annotationText := "hello-this-is-a-test-annotation"
	err := Annotate(inputFile, outputFile, annotationText)
	if err != nil {
		t.Errorf("Annotate(%s, %s, %s) = %v; wanted no errors", inputFile, outputFile, annotationText, err)
	}

	gotFile, err := os.ReadFile(outputFile)
	if err != nil {
		t.Errorf("ReadFile(%s) = %+v; needed success", outputFile, err)
	}

	expectedOutputFile := "expected-output.png"
	expectedFile, err := os.ReadFile(expectedOutputFile)
	if !bytes.Equal(gotFile, expectedFile) {
		t.Errorf("output file %s bytes not same as expected %s", outputFile, expectedOutputFile)
	}
}
