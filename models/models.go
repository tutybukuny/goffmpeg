package models

import (
	"fmt"

	"github.com/tutybukuny/goffmpeg/constant"
)

type Ffmpeg struct {
	FfmpegBinPath  string
	FfprobeBinPath string
}

type Metadata struct {
	Streams []Streams `json:"streams"`
	Format  Format    `json:"format"`
}

type Streams struct {
	Index              int
	ID                 string      `json:"id"`
	CodecName          string      `json:"codec_name"`
	CodecLongName      string      `json:"codec_long_name"`
	Profile            string      `json:"profile"`
	CodecType          string      `json:"codec_type"`
	CodecTimeBase      string      `json:"codec_time_base"`
	CodecTagString     string      `json:"codec_tag_string"`
	CodecTag           string      `json:"codec_tag"`
	Width              int         `json:"width"`
	Height             int         `json:"height"`
	CodedWidth         int         `json:"coded_width"`
	CodedHeight        int         `json:"coded_height"`
	HasBFrames         int         `json:"has_b_frames"`
	SampleAspectRatio  string      `json:"sample_aspect_ratio"`
	DisplayAspectRatio string      `json:"display_aspect_ratio"`
	PixFmt             string      `json:"pix_fmt"`
	Level              int         `json:"level"`
	ChromaLocation     string      `json:"chroma_location"`
	Refs               int         `json:"refs"`
	QuarterSample      string      `json:"quarter_sample"`
	DivxPacked         string      `json:"divx_packed"`
	RFrameRrate        string      `json:"r_frame_rate"`
	AvgFrameRate       string      `json:"avg_frame_rate"`
	TimeBase           string      `json:"time_base"`
	DurationTs         int         `json:"duration_ts"`
	Duration           string      `json:"duration"`
	Disposition        Disposition `json:"disposition"`
	BitRate            string      `json:"bit_rate"`
}

type Disposition struct {
	Default         int `json:"default"`
	Dub             int `json:"dub"`
	Original        int `json:"original"`
	Comment         int `json:"comment"`
	Lyrics          int `json:"lyrics"`
	Karaoke         int `json:"karaoke"`
	Forced          int `json:"forced"`
	HearingImpaired int `json:"hearing_impaired"`
	VisualImpaired  int `json:"visual_impaired"`
	CleanEffects    int `json:"clean_effects"`
}

type Format struct {
	Filename       string
	NbStreams      int    `json:"nb_streams"`
	NbPrograms     int    `json:"nb_programs"`
	FormatName     string `json:"format_name"`
	FormatLongName string `json:"format_long_name"`
	Duration       string `json:"duration"`
	Size           string `json:"size"`
	BitRate        string `json:"bit_rate"`
	ProbeScore     int    `json:"probe_score"`
	Tags           Tags   `json:"tags"`
}

type IProgress interface {
	ToString() string
	GetType() constant.ProgressType
}

type Progress struct {
	progressType constant.ProgressType
	Raw          string
}

func (p *Progress) ToString() string {
	return p.Raw
}

func (p *Progress) GetType() constant.ProgressType {
	return p.progressType
}

func NewProgress(line string) *Progress {
	return &Progress{progressType: constant.Raw, Raw: line}
}

type FrameProgress struct {
	progressType    constant.ProgressType
	FramesProcessed string
	CurrentTime     string
	CurrentBitrate  string
	Progress        float64
	Speed           string
	FPS             string
}

func NewFrameProgress() *FrameProgress {
	return &FrameProgress{progressType: constant.Frame}
}

func (p *FrameProgress) ToString() string {
	return "frame=" + p.FramesProcessed + " time=" + p.CurrentTime + " bitrate=" + p.CurrentBitrate +
		" progress=" + fmt.Sprintf("%f", p.Progress) + " speed=" + p.Speed
}

func (p *FrameProgress) GetType() constant.ProgressType {
	return p.progressType
}

type OpeningFileProgress struct {
	progressType constant.ProgressType
	FilePath     string
	WritingRate  string
	Bitrate      string
	Speed        string
}

func (p *OpeningFileProgress) ToString() string {
	str := "opening '" + p.FilePath + "' for writing"
	if p.WritingRate != "" {
		str += "rate=" + p.WritingRate
	}
	str += " speed" + p.Speed
	return str
}

func (p *OpeningFileProgress) GetType() constant.ProgressType {
	return p.progressType
}

func NewOpeningFileProgress() *OpeningFileProgress {
	return &OpeningFileProgress{progressType: constant.OpeningFile}
}

type Tags struct {
	Encoder string `json:"ENCODER"`
}
