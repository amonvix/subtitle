package srt

import (
	"fmt"
	"os"
	"strings"
	"subtitle/internal/domain"
	"time"
)

type Writer struct{}

func (w *Writer) Write(path string, segments []domain.Segment) error {
	var b strings.Builder
	for i, seg := range segments {
		idx := seg.Index
		if idx == 0 {
			idx = i + 1
		}
		b.WriteString(fmt.Sprintf("%d\n", idx))
		b.WriteString(fmt.Sprintf("%s --> %s\n", formatSRTTime(seg.Start), formatSRTTime(seg.End)))
		b.WriteString(strings.TrimSpace(seg.Text) + "\n\n")
	}
	return os.WriteFile(path, []byte(b.String()), 0o644)
}

func formatSRTTime(d time.Duration) string {
	hours := int(d / time.Hour)
	d -= time.Duration(hours) * time.Hour
	minutes := int(d / time.Minute)
	d -= time.Duration(minutes) * time.Minute
	seconds := int(d / time.Second)
	d -= time.Duration(seconds) * time.Second
	millis := int(d / time.Millisecond)

	return fmt.Sprintf("%02d:%02d:%02d,%03d", hours, minutes, seconds, millis)
}
