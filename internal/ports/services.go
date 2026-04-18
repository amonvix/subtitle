package ports

import "subtitle/internal/domain"

// VideoProcessor handles FFmpeg operations.
type VideoProcessor interface {
	ExtractAudio(videoPath, audioPath string) error
	BurnSubtitles(videoPath, srtPath, outputPath string) error
}

// Transcriber converts audio into subtitle segments.
type Transcriber interface {
	Transcribe(audioPath string) ([]domain.Segment, error)
}

// Translator transforms an SRT file into a target language.
type Translator interface {
	TranslateSRT(inputSRT, outputSRT, sourceLang, targetLang string) error
}

// SRTWriter persists subtitle segments in SRT format.
type SRTWriter interface {
	Write(path string, segments []domain.Segment) error
}
