package streamer

import (
	"fmt"

	"github.com/xfrr/goffmpeg/transcoder"
)

// Encoder is an interface for encoding video. Any type that wants to satisfy
// this interface must implement all its methods.
type Encoder interface {
	EncodeToMP4(v *Video, baseFileName string) error
}

// VideoEncoder is a type which satisfies the Encoder interface because it implements
// all the methods specified in Encoder.
type VideoEncoder struct{}

// EncodeToMP4 takes a Video object and a base file name, and encodes to MP4 format.
func (ve *VideoEncoder) EncodeToMP4(v *Video, baseFileName string) error {
	// Create a transcoder
	trans := new(transcoder.Transcoder)

	// Build the output path
	outputPath := fmt.Sprintf("%s/%s.mp4", v.OutputDir, baseFileName)

	// Initialize the transcoder
	err := trans.Initialize(v.InputFile, outputPath)
	if err != nil {
		return err
	}

	// Set codec
	trans.MediaFile().SetVideoCodec("libx264")

	// Start the transcoding process
	done := trans.Run(false)
	err = <-done
	if err != nil {
		return err
	}

	return nil
}
