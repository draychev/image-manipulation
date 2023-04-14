package main

import (
	"flag"

	"github.com/draychev/image-manipulation/pkg/annotate"
	"github.com/draychev/image-manipulation/pkg/logger"
)

var log = logger.NewPretty("image-manipulation/main")

func main() {
	annotationText := flag.String("annotation", "", "annotation to add to an image")
	inputFile := flag.String("inputfile", "", "input file name to manipulate")
	outputFile := flag.String("outputfile", "", "output file name to save the final result to")
	flag.Parse()

	// If the annotation flag was provided, print the annotation
	if *annotationText != "" && *inputFile != "" && *outputFile != "" {
		log.Info().Msgf("Will annotate image %s with %s into %s", *inputFile, *annotationText, *outputFile)
	} else {
		log.Fatal().Msg("Pass required parameters")
	}

	if err := annotate.Annotate(*inputFile, *outputFile, *annotationText); err != nil {
		log.Fatal().Msgf("Could not annotate image %s", *inputFile)
		return
	}
	log.Info().Msgf("Annotated image %sand saved it into %s", *inputFile, *outputFile)
}
