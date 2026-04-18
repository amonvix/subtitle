package usecase

import (
	"fmt"
	"os"
	"path/filepath"
	"subtitle/internal/ports"
)

// PipelineConfig holds all user-provided runtime options.
type PipelineConfig struct {
	InputVideo    string
	OutputVideo   string
	WorkingDir    string
	SourceLang    string
	TargetLang    string
	KeepArtifacts bool
}

// VideoPipeline orchestrates transcription + translation + subtitle burning.
type VideoPipeline struct {
	Video      ports.VideoProcessor
	Transcribe ports.Transcriber
	Translate  ports.Translator
	SRT        ports.SRTWriter
}

func (p *VideoPipeline) Run(cfg PipelineConfig) error {
	if cfg.WorkingDir == "" {
		cfg.WorkingDir = filepath.Join(os.TempDir(), "subtitle-work")
	}
	if err := os.MkdirAll(cfg.WorkingDir, 0o755); err != nil {
		return fmt.Errorf("create working dir: %w", err)
	}

	audioPath := filepath.Join(cfg.WorkingDir, "audio.wav")
	originalSRT := filepath.Join(cfg.WorkingDir, "original.srt")
	translatedSRT := filepath.Join(cfg.WorkingDir, "translated.srt")

	if err := p.Video.ExtractAudio(cfg.InputVideo, audioPath); err != nil {
		return fmt.Errorf("extract audio: %w", err)
	}

	segments, err := p.Transcribe.Transcribe(audioPath)
	if err != nil {
		return fmt.Errorf("transcribe audio: %w", err)
	}

	if err := p.SRT.Write(originalSRT, segments); err != nil {
		return fmt.Errorf("write original srt: %w", err)
	}

	if err := p.Translate.TranslateSRT(originalSRT, translatedSRT, cfg.SourceLang, cfg.TargetLang); err != nil {
		return fmt.Errorf("translate srt: %w", err)
	}

	if err := p.Video.BurnSubtitles(cfg.InputVideo, translatedSRT, cfg.OutputVideo); err != nil {
		return fmt.Errorf("burn subtitles: %w", err)
	}

	if !cfg.KeepArtifacts {
		_ = os.Remove(audioPath)
		_ = os.Remove(originalSRT)
		_ = os.Remove(translatedSRT)
	}

	return nil
}
