package oviewer

import (
	"path/filepath"
	"testing"

	"github.com/gdamore/tcell/v2"
)

func TestRoot_leftStatus(t *testing.T) {
	type fields struct {
		eventer   Eventer
		caption   string
		inputMode InputMode
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "test normal",
			fields: fields{
				caption:   "test",
				eventer:   normal(),
				inputMode: Normal,
			},
			want: "test:",
		},
		{
			name: "test search prompt",
			fields: fields{
				caption:   "test",
				eventer:   newSearchEvent(blankCandidate(), forward),
				inputMode: Search,
			},
			want: "/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := rootHelper(t)
			root.Doc.Caption = tt.fields.caption
			root.input.Event = tt.fields.eventer
			got, _ := root.leftStatus()
			gotStr := got.String()
			if gotStr != tt.want {
				t.Errorf("Root.leftStatus() got = %v, want %v", gotStr, tt.want)
			}
		})
	}
}

func TestRoot_leftStatus2(t *testing.T) {
	tcellNewScreen = fakeScreen
	defer func() {
		tcellNewScreen = tcell.NewScreen
	}()
	type fields struct {
		fileNames []string
		eventer   Eventer
		inputMode InputMode
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "test normal",
			fields: fields{
				fileNames: []string{filepath.Join(testdata, "test.txt"), filepath.Join(testdata, "section2.txt")},
				eventer:   normal(),
				inputMode: Normal,
			},
			want: "[0]" + filepath.Join(testdata, "test.txt") + ":",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := rootFileReadHelper(t, tt.fields.fileNames...)
			root.input.Event = tt.fields.eventer
			got, _ := root.leftStatus()
			gotStr := got.String()
			if gotStr != tt.want {
				t.Errorf("Root.leftStatus() got = %v, want %v", gotStr, tt.want)
			}
		})
	}
}

func TestRoot_statusDisplay(t *testing.T) {
	type fields struct {
		FollowSection bool
		FollowAll     bool
		FollowMode    bool
		FollowName    bool
		WatchMode     bool
		pauseFollow   bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "test watch",
			fields: fields{
				WatchMode: true,
			},
			want: "(Watch)",
		},
		{
			name: "test follow section",
			fields: fields{
				FollowSection: true,
			},
			want: "(Follow Section)",
		},
		{
			name: "test follow all",
			fields: fields{
				FollowAll: true,
			},
			want: "(Follow All)",
		},
		{
			name: "test follow name",
			fields: fields{
				FollowMode: true,
				FollowName: true,
			},
			want: "(Follow Name)",
		},
		{
			name: "test follow mode",
			fields: fields{
				FollowMode: true,
			},
			want: "(Follow Mode)",
		},
		{
			name: "pause follow",
			fields: fields{
				FollowMode:  true,
				pauseFollow: true,
			},
			want: "||(Follow Mode)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := rootHelper(t)
			root.Doc.WatchMode = tt.fields.WatchMode
			root.Doc.FollowSection = tt.fields.FollowSection
			root.FollowAll = tt.fields.FollowAll
			root.Doc.FollowName = tt.fields.FollowName
			root.Doc.FollowMode = tt.fields.FollowMode
			root.Doc.pauseFollow = tt.fields.pauseFollow
			if got := root.statusDisplay(); got != tt.want {
				t.Errorf("Root.statusDisplay() = %v, want %v", got, tt.want)
			}
		})
	}
}
