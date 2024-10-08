package oviewer

import (
	"reflect"
	"testing"

	"github.com/gdamore/tcell/v2"
)

func Test_escapeSequence_convert(t *testing.T) {
	type fields struct {
		state int
	}
	type args struct {
		st *parseState
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      bool
		wantState int
	}{
		{
			name: "test-escapeSequence",
			fields: fields{
				state: ansiText,
			},
			args: args{
				st: &parseState{
					mainc: 0x1b,
				},
			},
			want:      true,
			wantState: ansiEscape,
		},
		{
			name: "test-SubString",
			fields: fields{
				state: ansiEscape,
			},
			args: args{
				st: &parseState{
					mainc: 'P',
				},
			},
			want:      true,
			wantState: ansiSubstring,
		},
		{
			name: "test-SubString2",
			fields: fields{
				state: ansiSubstring,
			},
			args: args{
				st: &parseState{
					mainc: 0x1b,
				},
			},
			want:      true,
			wantState: ansiControlSequence,
		},
		{
			name: "test-OtherSequence",
			fields: fields{
				state: ansiEscape,
			},
			args: args{
				st: &parseState{
					mainc: '(',
				},
			},
			want:      true,
			wantState: otherSequence,
		},
		{
			name: "test-Other",
			fields: fields{
				state: ansiEscape,
			},
			args: args{
				st: &parseState{
					mainc: '@',
				},
			},
			want:      false,
			wantState: ansiText,
		},
		{
			name: "test-OtherSequence2",
			fields: fields{
				state: otherSequence,
			},
			args: args{
				st: &parseState{
					mainc: 'a',
				},
			},
			want:      true,
			wantState: ansiEscape,
		},
		{
			name: "test-ControlSequence",
			fields: fields{
				state: ansiControlSequence,
			},
			args: args{
				st: &parseState{
					mainc: 'm',
				},
			},
			want:      true,
			wantState: ansiText,
		},
		{
			name: "test-ControlSequence2",
			fields: fields{
				state: ansiControlSequence,
			},
			args: args{
				st: &parseState{
					mainc: 'A',
				},
			},
			want:      true,
			wantState: ansiText,
		},
		{
			name: "test-SysSequence",
			fields: fields{
				state: systemSequence,
			},
			args: args{
				st: &parseState{
					mainc: 0x07,
				},
			},
			want:      true,
			wantState: ansiText,
		},
		{
			name: "test-OscHyperLink",
			fields: fields{
				state: oscHyperLink,
			},
			args: args{
				st: &parseState{
					mainc: 'a',
				},
			},
			want:      false,
			wantState: ansiText,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			es := newESConverter()
			es.state = tt.fields.state
			if got := es.convert(tt.args.st); got != tt.want {
				t.Errorf("escapeSequence.convert() = %v, want %v", got, tt.want)
			}
			if es.state != tt.wantState {
				t.Errorf("escapeSequence.convert() = %v, want %v", es.state, tt.wantState)
			}
		})
	}
}

func Test_csToStyle(t *testing.T) {
	t.Parallel()
	type args struct {
		style        tcell.Style
		csiParameter string
	}
	tests := []struct {
		name string
		args args
		want tcell.Style
	}{
		{
			name: "color8bit",
			args: args{
				style:        tcell.StyleDefault,
				csiParameter: "38;5;1",
			},
			want: tcell.StyleDefault.Foreground(tcell.ColorMaroon),
		},
		{
			name: "color8bit2",
			args: args{
				style:        tcell.StyleDefault,
				csiParameter: "38;5;21",
			},
			want: tcell.StyleDefault.Foreground(tcell.GetColor("#0000ff")),
		},
		{
			name: "attributes",
			args: args{
				style:        tcell.StyleDefault,
				csiParameter: "2;3;4;5;6;7;8;9",
			},
			want: tcell.StyleDefault.Dim(true).Italic(true).Underline(true).Blink(true).Reverse(true).StrikeThrough(true),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := csToStyle(tt.args.style, tt.args.csiParameter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("csToStyle() = %v, want %v", got, tt.want)
				gfg, gbg, gattr := got.Decompose()
				wfg, wbg, wattr := tt.want.Decompose()
				t.Errorf("csToStyle() = %x,%x,%v, want %x,%x,%v", gfg.Hex(), gbg.Hex(), gattr, wfg.Hex(), wbg.Hex(), wattr)
			}
		})
	}
}

func Test_parseCSI(t *testing.T) {
	type args struct {
		params string
	}
	tests := []struct {
		name string
		args args
		want OVStyle
	}{
		{
			name: "test-attributes",
			args: args{
				params: "2;3;4;5;6;7;8;9",
			},
			want: OVStyle{
				Dim:           true,
				Italic:        true,
				Underline:     true,
				Blink:         true,
				Reverse:       true,
				StrikeThrough: true,
			},
		},
		{
			name: "test-attributesErr",
			args: args{
				params: "38;38;38",
			},
			want: OVStyle{
				Dim:           false,
				Italic:        false,
				Underline:     false,
				Blink:         false,
				Reverse:       false,
				StrikeThrough: false,
			},
		},
		{
			name: "test-attributesNone",
			args: args{
				params: "28",
			},
			want: OVStyle{
				Dim:           false,
				Italic:        false,
				Underline:     false,
				Blink:         false,
				Reverse:       false,
				StrikeThrough: false,
			},
		},
		{
			name: "test-Default",
			args: args{
				params: "49",
			},
			want: OVStyle{
				Background:    "default",
				Dim:           false,
				Italic:        false,
				Underline:     false,
				Blink:         false,
				Reverse:       false,
				StrikeThrough: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseCSI(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseCSI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_colorName(t *testing.T) {
	type args struct {
		colorNumber int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test-ColorName1",
			args: args{
				colorNumber: 1,
			},
			want: "maroon",
		},
		{
			name: "test-ColorName249",
			args: args{
				colorNumber: 249,
			},
			want: "#bcbcbc",
		},
		{
			name: "test-ColorNameNotImplemented",
			args: args{
				colorNumber: 999,
			},
			want: "",
		},
		{
			name: "test-ColorNameMinus",
			args: args{
				colorNumber: -1,
			},
			want: "black",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := colorName(tt.args.colorNumber); got != tt.want {
				t.Errorf("colorName() = %v, want %v", got, tt.want)
			}
		})
	}
}
