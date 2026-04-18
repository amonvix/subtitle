package video

import (
	"fmt"
	"os/exec"
)

type FFmpegService struct{}

func (f *FFmpegService) ExtractAudio(videoPath, audioPath string) error {
	cmd := exec.Command(
		"ffmpeg",
		"-y",
		"-i", videoPath,
		"-vn",
		"-acodec", "pcm_s16le",
		"-ar", "16000",
		"-ac", "1",
		audioPath,
	)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("ffmpeg extract error: %w - %s", err, string(output))
	}
	return nil
}

func (f *FFmpegService) BurnSubtitles(videoPath, srtPath, outputPath string) error {
	cmd := exec.Command(
		"ffmpeg",
		"-y",
		"-i", videoPath,
		"-vf", fmt.Sprintf("subtitles=%s", srtPath),
		"-c:a", "copy",
		outputPath,
	)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("ffmpeg burn subtitles error: %w - %s", err, string(output))
	}
	return nil
}
