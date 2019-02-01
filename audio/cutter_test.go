package audio_test

import (
	"testing"
	"time"

	"github.com/sandovalrr/mediacutter"
	"github.com/sandovalrr/mediacutter/audio"
)

var samples = time.Duration(15)
var fileName = "../assets/audio_test.mp3"

func TestNewAudioCutter(t *testing.T) {
	cutter := audio.NewAudioCutter(mediacutter.CutterOption{
		Name:    fileName,
		Samples: samples,
	})

	if cutter == nil {
		t.Error("Expected cutter were defined")
	}
}

func TestLen(t *testing.T) {
	cutter := audio.NewAudioCutter(mediacutter.CutterOption{
		Name:    fileName,
		Samples: samples,
	})

	len, err := cutter.Len()
	if err != nil {
		t.Errorf("Not expection errors, but found %v", err)
	}
	if len == time.Duration(0) {
		t.Error("Expected file length > 0")
	}
}
