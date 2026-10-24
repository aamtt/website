package main

import "math"

type Project struct {
	Title       string
	Description string
	Year        int
	URL         string
	Color       string
	A, B, C     int
	Phase       float64
}

var projects = []Project{
	{
		Title:       "pk",
		Description: "take a p(ee)k inside a live output",
		Year:        2026,
		URL:         "https://github.com/aamtt/pk",
		Color:       "#3F6B6B",
		A:           5,
		B:           5,
		C:           2,
		Phase:       0.785,
	},
	{
		Title:       "conway",
		Description: "conway's game of life from file content",
		Year:        2026,
		URL:         "https://github.com/aamtt/conway",
		Color:       "#5E4A72",
		A:           3,
		B:           2,
		C:           3,
		Phase:       0.9,
	},
	{
		Title:       "fossil",
		Description: "a transparent compressor",
		Year:        2026,
		URL:         "https://github.com/aamtt/fossil",
		Color:       "#8C2230",
		A:           2,
		B:           3,
		C:           5,
		Phase:       1.05,
	},
	{
		Title:       "jtp",
		Description: "a binary protocol for images",
		Year:        2026,
		URL:         "https://github.com/aamtt/jtp",
		Color:       "#3A5068",
		A:           3,
		B:           1,
		C:           2,
		Phase:       0,
	},
	{
		Title:       "oscilloscope",
		Description: "images converted to audio waveforms",
		Year:        2023,
		URL:         "https://github.com/aamtt/oscilloscope",
		Color:       "#4A6B4A",
		A:           1,
		B:           2,
		C:           3,
		Phase:       math.Pi / 2,
	},
	{
		Title:       "use listen along",
		Description: "a react hook for spotify sync",
		Year:        2021,
		URL:         "https://github.com/aamtt/use-listen-along",
		Color:       "#A8843F",
		A:           1,
		B:           1,
		C:           2,
		Phase:       0,
	},
}

var elsewhere = []Project{
	{
		Title:       "receipt",
		Description: "the site this one replaced",
		URL:         "/receipt",
		Color:       "#8a8279",
		A:           4,
		B:           3,
		C:           5,
		Phase:       0.35,
	},
	{
		Title:       "guestbook",
		Description: "sign it from your terminal",
		URL:         "/guestbook",
		Color:       "#8a8279",
		A:           2,
		B:           5,
		C:           3,
		Phase:       1.2,
	},
}
