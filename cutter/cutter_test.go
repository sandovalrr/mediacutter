package cutter_test

import (
	"testing"
	"time"

	"github.com/sandovalrr/mediacutter/cutter"
	"github.com/sandovalrr/mediacutter/models"
)

var samples = time.Duration(4)
var audioPath = "../assets/piano2.wav"
var videoPath = "../assets/video.avi"
var chunkPath = "../assets/chunks"

func TestNewAudioCutter(t *testing.T) {
	cutter := cutter.NewAudioCutter(models.CutterOption{
		Name:      audioPath,
		Samples:   samples,
		ChunkPath: chunkPath,
	})

	if cutter == nil {
		t.Error("Expected cutter were defined")
	}

	if cutter.Repo == nil {
		t.Error("Expected cutter repo were defined")
	}
}

func TestNewVideoCutter(t *testing.T) {
	cutter := cutter.NewVideoCutter(models.CutterOption{
		Name:      videoPath,
		Samples:   samples,
		ChunkPath: chunkPath,
	})

	if cutter == nil {
		t.Error("Expected cutter were defined")
	}

	if cutter.Repo == nil {
		t.Error("Expected cutter repo were defined")
	}
}

func TestAudioLen(t *testing.T) {
	cutter := cutter.NewAudioCutter(models.CutterOption{
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

func TestVideoLen(t *testing.T) {
	cutter := cutter.NewVideoCutter(models.CutterOption{
		Name:      videoPath,
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

func TestAudioSplit(t *testing.T) {
	cutter := cutter.NewAudioCutter(models.CutterOption{
		Name:      audioPath,
		Samples:   samples,
		ChunkPath: chunkPath,
	})

	err := cutter.Split()
	if err != nil {
		t.Errorf("Not errors expected, but found %v", err)
	}
}
func TestVideoSplit(t *testing.T) {
	cutter := cutter.NewVideoCutter(models.CutterOption{
		Name:      videoPath,
		Samples:   samples,
		ChunkPath: chunkPath,
	})

	err := cutter.Split()
	if err != nil {
		t.Errorf("Not errors expected, but found %v", err)
	}
}
