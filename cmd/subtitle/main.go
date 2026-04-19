package main

import (
	"flag"
	"log"

	"subtitle/internal/infrastructure/srt"
	"subtitle/internal/infrastructure/transcription"
	"subtitle/internal/infrastructure/translation"
	"subtitle/internal/infrastructure/video"
	"subtitle/internal/usecase"
)

func main() {
	input := flag.String("input", "", "input video file")
	output := flag.String("output", "output.mp4", "output video file")
	source := flag.String("source", "pt", "source language")
	target := flag.String("target", "en", "target language")

	flag.Parse()

	if *input == "" {
		log.Fatal("input file is required")
	}

	videoService := video.NewFFmpegService()
	transcriptionService := transcription.NewWhisperService()
	translationService := translation.NewLibreTranslateService()
	srtWriter := srt.NewWriter()

	pipeline := usecase.NewPipeline(
		videoService,
		transcriptionService,
		translationService,
		srtWriter,
	)

	err := pipeline.Execute(*input, *output, *source, *target)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Processing completed successfully")
}
