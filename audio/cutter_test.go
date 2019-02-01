package audio_test

import (
	"testing"
	"time"

	"github.com/sandovalrr/mediacutter"
	"github.com/sandovalrr/mediacutter/audio"
)

var samples = time.Duration(4)
var audioPath = "../assets/piano2.wav"
var chunkPath = "../assets/chunks"

func TestNewAudioCutter(t *testing.T) {
	cutter := audio.NewAudioCutter(mediacutter.CutterOption{
		Name:      audioPath,
		Samples:   samples,
		ChunkPath: chunkPath,
	})

	if cutter == nil {
		t.Error("Expected cutter were defined")
	}
}

func TestLen(t *testing.T) {
	cutter := audio.NewAudioCutter(mediacutter.CutterOption{
		Name:      audioPath,
		Samples:   samples,
		ChunkPath: chunkPath,
	})

	len, err := cutter.Len()
	if err != nil {
		t.Errorf("Not errors expected, but found %v", err)
	}
	if len == time.Duration(0) {
		t.Error("Expected file length > 0")
	}
}

func TestSplit(t *testing.T) {
	cutter := audio.NewAudioCutter(mediacutter.CutterOption{
		Name:      audioPath,
		Samples:   samples,
		ChunkPath: chunkPath,
	})

	err := cutter.Split()
	if err != nil {
		t.Errorf("Not errors expected, but found %v", err)
	}
}
