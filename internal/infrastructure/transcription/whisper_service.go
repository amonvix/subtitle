package transcription

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"subtitle/internal/domain"
	"time"
)

type whisperSegment struct {
	Start float64 `json:"start"`
	End   float64 `json:"end"`
	Text  string  `json:"text"`
}

type WhisperService struct {
	ScriptPath string
	Model      string
}

func NewWhisperService() *WhisperService {
	return &WhisperService{}
}

func (w *WhisperService) Transcribe(audioPath string) ([]domain.Segment, error) {
	if w.ScriptPath == "" {
		w.ScriptPath = "scripts/transcribe.py"
	}
	if w.Model == "" {
		w.Model = "base"
	}

	jsonPath := audioPath + ".json"
	cmd := exec.Command("python3", w.ScriptPath, "--audio", audioPath, "--output", jsonPath, "--model", w.Model)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("whisper script failed: %w - %s", err, string(output))
	}

	raw, err := os.ReadFile(jsonPath)
	if err != nil {
		return nil, fmt.Errorf("read whisper output: %w", err)
	}

	var parsed []whisperSegment
	if err := json.Unmarshal(raw, &parsed); err != nil {
		return nil, fmt.Errorf("parse whisper json: %w", err)
	}

	segments := make([]domain.Segment, 0, len(parsed))
	for i, seg := range parsed {
		segments = append(segments, domain.Segment{
			Index: i + 1,
			Start: time.Duration(seg.Start * float64(time.Second)),
			End:   time.Duration(seg.End * float64(time.Second)),
			Text:  seg.Text,
		})
	}

	_ = os.Remove(jsonPath)
	return segments, nil
}
