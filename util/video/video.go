package video

import (
	"os"
	"time"

	"github.com/alfg/mp4"
)

func GetVideoDuration(filePath string) (time.Duration, error) {
	// Open the video file
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}
	// Decode the MP4 file
	mp4file, err := mp4.OpenFromReader(file, info.Size())
	if err != nil {
		return 0, err
	}

	// Get the duration from the MP4 file
	duration := mp4file.Moov.Mvhd.Duration

	return time.Duration(duration) * 1_000_000, nil
}
