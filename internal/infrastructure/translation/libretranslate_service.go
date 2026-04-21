package translation

import (
	"fmt"
	"os/exec"
)

type LibreTranslateService struct {
	ScriptPath string
	Endpoint   string
}

func NewLibreTranslateService() *LibreTranslateService{
	return &LibreTranslateService{}
}

func (l *LibreTranslateService) TranslateSRT(inputSRT, outputSRT, sourceLang, targetLang string) error {
	if l.ScriptPath == "" {
		l.ScriptPath = "scripts/translate_srt.py"
	}
	if l.Endpoint == "" {
		l.Endpoint = "https://libretranslate.com/translate"
	}

	cmd := exec.Command(
		"python3",
		l.ScriptPath,
		"--input", inputSRT,
		"--output", outputSRT,
		"--source", sourceLang,
		"--target", targetLang,
		"--endpoint", l.Endpoint,
	)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("translation script failed: %w - %s", err, string(output))
	}
	return nil
}
