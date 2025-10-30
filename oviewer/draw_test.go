package oviewer

import (
	"context"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gdamore/tcell/v2"
)

func getContents(t *testing.T, root *Root, y int, width int) string {
	t.Helper()
	var buf strings.Builder
	for x := range width {
		r, _, _, _ := root.Screen.GetContent(x, y)
		buf.WriteRune(r)
	}
	return buf.String()
}

func TestRoot_draw(t *testing.T) {
	tcellNewScreen = fakeScreen
	defer func() {
		tcellNewScreen = tcell.NewScreen
	}()
	type fields struct {
		fileNames   []string
		header      int
		wrapMode    bool
		lineNumMode bool
	}
	type args struct {
		ctx context.Context
	}
	type want struct {
		bottomLN int
		bottomLX int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{
			name: "testDraw1",
			fields: fields{
				fileNames:   []string{filepath.Join(testdata, "test.txt")},
				header:      0,
				wrapMode:    false,
				lineNumMode: true,
			},
			args: args{
				ctx: context.Background(),
			},
			want: want{
				bottomLN: 24,
				bottomLX: 0,
			},
		},
		{
			name: "testDraw2",
			fields: fields{
				fileNames: []string{filepath.Join(testdata, "test.txt")},
				header:    0,
				wrapMode:  true,
			},
			args: args{
				ctx: context.Background(),
			},
			want: want{
				bottomLN: 24,
				bottomLX: 0,
			},
		},
		{
			name: "testDraw-header",
			fields: fields{
				fileNames:   []string{filepath.Join(testdata, "test.txt")},
				header:      1,
				wrapMode:    false,
				lineNumMode: true,
			},
			args: args{
				ctx: context.Background(),
			},
			want: want{
				bottomLN: 24,
				bottomLX: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := rootFileReadHelper(t, tt.fields.fileNames...)
			root.Doc.Header = tt.fields.header
			root.Doc.WrapMode = tt.fields.wrapMode
			root.Doc.LineNumMode = tt.fields.lineNumMode
			root.ViewSync(tt.args.ctx)
			root.prepareScreen()
			root.draw(tt.args.ctx)
			if root.Doc.bottomLN != tt.want.bottomLN {
				t.Errorf("Root.draw() bottomLN = %v, want %v", root.Doc.bottomLN, tt.want.bottomLN)
			}
			if root.Doc.bottomLX != tt.want.bottomLX {
				t.Errorf("Root.draw() bottomLX = %v, want %v", root.Doc.bottomLX, tt.want.bottomLX)
			}
		})
	}
}

func TestRoot_draw2(t *testing.T) {
	tcellNewScreen = fakeScreen
	defer func() {
		tcellNewScreen = tcell.NewScreen
	}()
	type fields struct {
		fileNames        []string
		header           int
		sectionDelimiter string
		hideOtherSection bool
		wrapMode         bool
		lineNumMode      bool
		topLN            int
	}
	type args struct {
		ctx context.Context
	}
	type want struct {
		bottomLN            int
		bottomLX            int
		sectionHeaderHeight int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{
			name: "testDraw-section-header",
			fields: fields{
				fileNames:        []string{filepath.Join(testdata, "section-header.txt")},
				header:           3,
				sectionDelimiter: "^#",
				hideOtherSection: false,
				wrapMode:         false,
				lineNumMode:      true,
				topLN:            0,
			},
			args: args{
				ctx: context.Background(),
			},
			want: want{
				bottomLN:            24,
				bottomLX:            0,
				sectionHeaderHeight: 1,
			},
		},
		{
			name: "testDraw-section-header2",
			fields: fields{
				fileNames:        []string{filepath.Join(testdata, "section-header.txt")},
				header:           3,
				sectionDelimiter: "^#",
				hideOtherSection: false,
				wrapMode:         true,
				lineNumMode:      true,
				topLN:            0,
			},
			args: args{
				ctx: context.Background(),
			},
			want: want{
				bottomLN:            20,
				bottomLX:            0,
				sectionHeaderHeight: 1,
			},
		},
		{
			name: "testDraw-section-header2",
			fields: fields{
				fileNames:        []string{filepath.Join(testdata, "section-header.txt")},
				header:           3,
				sectionDelimiter: "^#",
				hideOtherSection: false,
				wrapMode:         true,
				lineNumMode:      true,
				topLN:            10,
			},
			args: args{
				ctx: context.Background(),
			},
			want: want{
				bottomLN:            32,
				bottomLX:            0,
				sectionHeaderHeight: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := rootFileReadHelper(t, tt.fields.fileNames...)
			root.Doc.Header = tt.fields.header
			if tt.fields.sectionDelimiter != "" {
				root.Doc.setSectionDelimiter(tt.fields.sectionDelimiter)
				root.Doc.SectionHeaderNum = 1
				root.Doc.SectionHeader = true
				root.Doc.HideOtherSection = tt.fields.hideOtherSection
			}
			root.Doc.topLN = tt.fields.topLN
			root.Doc.WrapMode = tt.fields.wrapMode
			root.Doc.LineNumMode = tt.fields.lineNumMode
			root.ViewSync(tt.args.ctx)
			root.prepareScreen()
			root.draw(tt.args.ctx)
			if root.Doc.bottomLN != tt.want.bottomLN {
				t.Errorf("Root.draw() bottomLN = %v, want %v", root.Doc.bottomLN, tt.want.bottomLN)
			}
			if root.Doc.bottomLX != tt.want.bottomLX {
				t.Errorf("Root.draw() bottomLX = %v, want %v", root.Doc.bottomLX, tt.want.bottomLX)
			}
			if got := root.Doc.sectionHeaderHeight; got != tt.want.sectionHeaderHeight {
				t.Errorf("Root.draw() sectionHeaderHeight = %v, want %v", got, tt.want.sectionHeaderHeight)
			}
		})
	}
}

func TestRoot_section2(t *testing.T) {
	tcellNewScreen = fakeScreen
	defer func() {
		tcellNewScreen = tcell.NewScreen
	}()
	type fields struct {
		fileNames        []string
		header           int
		sectionDelimiter string
		sectionHeaderNum int
		hideOtherSection bool
		wrapMode         bool
		lineNumMode      bool
		topLN            int
	}
	type args struct {
		ctx context.Context
	}
	type want struct {
		bottomLN            int
		bottomLX            int
		sectionHeaderHeight int
		wantY               int
		wantStr             string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{
			name: "testDraw-section2-1",
			fields: fields{
				fileNames:        []string{filepath.Join(testdata, "section2.txt")},
				header:           0,
				sectionDelimiter: "^-",
				sectionHeaderNum: 3,
				hideOtherSection: false,
				wrapMode:         true,
				lineNumMode:      true,
				topLN:            0,
			},
			args: args{
				ctx: context.Background(),
			},
			want: want{
				bottomLN:            24,
				bottomLX:            0,
				sectionHeaderHeight: 0,
				wantY:               0,
				wantStr:             " 1 test1  ",
			},
		},
		{
			name: "testDraw-section2-2",
			fields: fields{
				fileNames:        []string{filepath.Join(testdata, "section2.txt")},
				header:           0,
				sectionDelimiter: "^-",
				sectionHeaderNum: 1,
				hideOtherSection: false,
				wrapMode:         true,
				lineNumMode:      true,
				topLN:            4,
			},
			args: args{
				ctx: context.Background(),
			},
			want: want{
				bottomLN:            28,
				bottomLX:            0,
				sectionHeaderHeight: 1,
				wantY:               8,
				wantStr:             "13        ",
			},
		},
		{
			name: "testDraw-section2-3",
			fields: fields{
				fileNames:        []string{filepath.Join(testdata, "section2.txt")},
				header:           0,
				sectionDelimiter: "^-",
				sectionHeaderNum: 1,
				hideOtherSection: true,
				wrapMode:         true,
				lineNumMode:      true,
				topLN:            3,
			},
			args: args{
				ctx: context.Background(),
			},
			want: want{
				bottomLN:            27,
				bottomLX:            0,
				sectionHeaderHeight: 1,
				wantY:               10,
				wantStr:             "          ",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := rootFileReadHelper(t, tt.fields.fileNames...)
			root.Doc.Header = tt.fields.header
			if tt.fields.sectionDelimiter != "" {
				root.Doc.setSectionDelimiter(tt.fields.sectionDelimiter)
				root.Doc.SectionHeaderNum = tt.fields.sectionHeaderNum
				root.Doc.SectionHeader = true
				root.Doc.HideOtherSection = tt.fields.hideOtherSection
			}
			root.Doc.topLN = tt.fields.topLN
			root.Doc.WrapMode = tt.fields.wrapMode
			root.Doc.LineNumMode = tt.fields.lineNumMode
			root.Doc.HideOtherSection = tt.fields.hideOtherSection
			root.ViewSync(tt.args.ctx)
			root.prepareScreen()
			root.draw(tt.args.ctx)
			if root.Doc.bottomLN != tt.want.bottomLN {
				t.Errorf("Root.draw() bottomLN = %v, want %v", root.Doc.bottomLN, tt.want.bottomLN)
			}
			if root.Doc.bottomLX != tt.want.bottomLX {
				t.Errorf("Root.draw() bottomLX = %v, want %v", root.Doc.bottomLX, tt.want.bottomLX)
			}
			if got := root.Doc.sectionHeaderHeight; got != tt.want.sectionHeaderHeight {
				t.Errorf("Root.draw() sectionHeaderHeight = %v, want %v", got, tt.want.sectionHeaderHeight)
			}
			str := getContents(t, root, tt.want.wantY, 10)
			if got := str; got != tt.want.wantStr {
				t.Errorf("Root.draw() = [%v], want [%v]", got, tt.want.wantStr)
			}
		})
	}
}

func TestRoot_sectionLineHighlight(t *testing.T) {
	tcellNewScreen = fakeScreen
	defer func() {
		tcellNewScreen = tcell.NewScreen
	}()
	type fields struct {
		fileNames        []string
		sectionDelimiter string
		sectionHeaderNum int
	}
	type args struct {
		y int
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantStr   rune
		wantStyle bool
	}{
		{
			name: "test-sectionLineHighlight1",
			fields: fields{
				fileNames:        []string{filepath.Join(testdata, "section.txt")},
				sectionDelimiter: "^#",
				sectionHeaderNum: 1,
			},
			args: args{
				y: 2,
			},
			wantStr:   '#',
			wantStyle: true,
		},
		{
			name: "test-sectionLineHighlight2",
			fields: fields{
				fileNames:        []string{filepath.Join(testdata, "section.txt")},
				sectionDelimiter: "^#",
				sectionHeaderNum: 1,
			},
			args: args{
				y: 3,
			},
			wantStr:   '2',
			wantStyle: false,
		},
		{
			name: "test-sectionLineHighlightInvalid",
			fields: fields{
				fileNames:        []string{filepath.Join(testdata, "section.txt")},
				sectionDelimiter: "^#",
				sectionHeaderNum: 2,
			},
			args: args{
				y: 20,
			},
			wantStr:   '~',
			wantStyle: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := rootFileReadHelper(t, tt.fields.fileNames...)
			root.Doc.SectionHeader = true
			root.setSectionDelimiter(tt.fields.sectionDelimiter)
			root.Doc.SectionHeaderNum = tt.fields.sectionHeaderNum
			root.prepareScreen()
			ctx := context.Background()
			root.ViewSync(ctx)
			root.prepareDraw(ctx)
			root.draw(ctx)
			line, ok := root.scr.lines[tt.args.y]
			if !ok {
				t.Fatalf("line is not found %d", tt.args.y)
			}
			root.sectionLineHighlight(tt.args.y, line)
			p, _, style, _ := root.Screen.GetContent(0, tt.args.y)
			if p != tt.wantStr {
				t.Errorf("Root.sectionLineHighlight() = %v, want %v", p, tt.wantStr)
			}
			if tt.wantStyle {
				if style != applyStyle(tcell.StyleDefault, root.Doc.Style.SectionLine) {
					t.Errorf("Root.sectionLineHighlight() = %v, want %v", style, root.Doc.Style.SectionLine)
				}
			} else {
				if style == applyStyle(tcell.StyleDefault, root.Doc.Style.SectionLine) {
					t.Errorf("Root.sectionLineHighlight() = %v, want %v", style, tcell.StyleDefault)
				}
			}
		})
	}
}

func TestRoot_drawVerticalHeader(t *testing.T) {
	tcellNewScreen = fakeScreen
	defer func() {
		tcellNewScreen = tcell.NewScreen
	}()
	type fields struct {
		verticalHeader int
		headerColumn   int
	}
	type args struct {
		y     int
		lineC LineC
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "testDrawVerticalHeader",
			fields: fields{
				verticalHeader: 3,
				headerColumn:   0,
			},
			args: args{
				y: 0,
				lineC: LineC{
					lc: []content{
						{mainc: 'A', style: tcell.StyleDefault},
						{mainc: 'B', style: tcell.StyleDefault},
						{mainc: 'C', style: tcell.StyleDefault},
					},
					valid: true,
				},
			},
			want: "ABC",
		},
		{
			name: "testDrawHeaderColumn",
			fields: fields{
				verticalHeader: 0,
				headerColumn:   2,
			},
			args: args{
				y: 0,
				lineC: LineC{
					lc: []content{
						{mainc: 'A', style: tcell.StyleDefault},
						{mainc: 'B', style: tcell.StyleDefault},
						{mainc: 'C', style: tcell.StyleDefault},
					},
					columnRanges: []columnRange{
						{start: 0, end: 1},
						{start: 2, end: 2},
					},
					valid: true,
				},
			},
			want: "AB",
		},
		{
			name: "testDrawVerticalHeaderNone",
			fields: fields{
				verticalHeader: 0,
				headerColumn:   0,
			},
			args: args{
				y: 0,
				lineC: LineC{
					lc: []content{
						{mainc: 'A', style: tcell.StyleDefault},
						{mainc: 'B', style: tcell.StyleDefault},
						{mainc: 'C', style: tcell.StyleDefault},
					},
					valid: true,
				},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := rootHelper(t)
			root.Doc.WrapMode = false
			root.Doc.VerticalHeader = tt.fields.verticalHeader
			root.Doc.HeaderColumn = tt.fields.headerColumn
			root.prepareScreen()
			root.drawVerticalHeader(tt.args.y, 0, tt.args.lineC)
			got := getContents(t, root, tt.args.y, len(tt.want))
			if got != tt.want {
				t.Errorf("Root.drawVerticalHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoot_calculateVerticalHeader(t *testing.T) {
	type fields struct {
		verticalHeader int
		headerColumn   int
	}
	type args struct {
		lineC LineC
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "testCalculateVerticalHeader",
			fields: fields{
				verticalHeader: 3,
				headerColumn:   0,
			},
			args: args{
				lineC: LineC{},
			},
			want: 3,
		},
		{
			name: "testCalculateHeaderColumn",
			fields: fields{
				verticalHeader: 0,
				headerColumn:   2,
			},
			args: args{
				lineC: LineC{
					columnRanges: []columnRange{
						{start: 0, end: 1},
						{start: 2, end: 2},
					},
				},
			},
			want: 3,
		},
		{
			name: "testCalculateVerticalHeaderNone",
			fields: fields{
				verticalHeader: 0,
				headerColumn:   0,
			},
			args: args{
				lineC: LineC{},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := rootHelper(t)
			root.Doc.VerticalHeader = tt.fields.verticalHeader
			root.Doc.HeaderColumn = tt.fields.headerColumn
			if got := root.Doc.widthVerticalHeader(tt.args.lineC); got != tt.want {
				t.Errorf("Root.calculateVerticalHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNeedsDisplaySync(t *testing.T) {
	tests := []struct {
		name     string
		mainc    rune
		combc    []rune
		expected bool
	}{
		// No combining characters, not emoji
		{"ASCII letter", 'a', nil, false},
		{"ASCII digit", '5', nil, false},
		{"ASCII space", ' ', nil, false},

		// CJK characters (wide, but not ambiguous)
		{"CJK character", 'あ', nil, false},

		// Latin-1 characters (some are ambiguous!)
		{"Latin à (ambiguous)", 'à', nil, true},
		{"Latin á (ambiguous)", 'á', nil, true},
		{"Latin â (not ambiguous)", 'â', nil, false},
		{"Latin é (ambiguous)", 'é', nil, true},
		{"Latin ë (not ambiguous)", 'ë', nil, false},

		// Traditional ambiguous width characters
		{"Section sign", '§', nil, true},
		{"Degree sign", '°', nil, true},
		{"Plus-minus", '±', nil, true},

		// Box drawing characters
		{"Box horizontal", '─', nil, true},
		{"Box vertical", '│', nil, true},

		// Mathematical symbols (ambiguous)
		{"Mathematical symbol ∀", '∀', nil, true},
		{"Mathematical symbol ∃", '∃', nil, true},

		// Arrow symbols (ambiguous)
		{"Arrow →", '→', nil, true},
		{"Arrow ←", '←', nil, true},
		{"Arrow ↑", '↑', nil, true},
		{"Arrow ↓", '↓', nil, true},

		// With combining characters (always true)
		{"ASCII with combining", 'a', []rune{0x0300}, true},
		{"Emoji with ZWJ", '👨', []rune{0x200D}, true},
		{"Emoji with variation selector", '⚕', []rune{0xFE0F}, true},
		{"Emoji with skin tone", '👋', []rune{0x1F3FB}, true},

		// Emoji range test (no combining characters)
		{"Regional indicator", '🇺', nil, true}, // U+1F1FA
		{"Another regional", '🇲', nil, true},   // U+1F1F2
		{"Fire emoji", '🔥', nil, true},         // U+1F525
		{"Rocket emoji", '🚀', nil, true},       // U+1F680
		{"Grinning face", '😀', nil, true},      // U+1F600
		{"Thinking face", '🤔', nil, true},      // U+1F914
		{"Rainbow", '🌈', nil, true},            // U+1F308

		// Outside emoji range
		{"Before emoji range", 0x1F1DF, nil, false},
		{"After emoji range", 0x1FA00, nil, false},

		// Boundary value test
		{"First regional", 0x1F1E0, nil, true},
		{"Last in range", 0x1F9FF, nil, true},
		{"Just after range", 0x1FA00, nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := needsDisplaySync(tt.mainc, tt.combc)
			if result != tt.expected {
				t.Errorf("needsDisplaySync(%c / U+%04X, %v) = %v, want %v",
					tt.mainc, tt.mainc, tt.combc, result, tt.expected)
			}
		})
	}
}

func TestNeedsDisplaySyncComplexSequences(t *testing.T) {
	tests := []struct {
		name     string
		mainc    rune
		combc    []rune
		expected bool
	}{
		// ZWJ sequences
		{"ZWJ sequence start", '👨', []rune{0x200D, '⚕', 0xFE0F}, true},
		{"With multiple combining", '👋', []rune{0x1F3FB, 0x200D, '♂', 0xFE0F}, true},

		// Variation selectors
		{"Variation selector 15", '⚕', []rune{0xFE0E}, true},
		{"Variation selector 16", '⚕', []rune{0xFE0F}, true},

		// Skin tone modifiers
		{"Light skin tone", '👋', []rune{0x1F3FB}, true},
		{"Medium-light skin", '👋', []rune{0x1F3FC}, true},
		{"Medium skin", '👋', []rune{0x1F3FD}, true},
		{"Medium-dark skin", '👋', []rune{0x1F3FE}, true},
		{"Dark skin tone", '👋', []rune{0x1F3FF}, true},

		// Multiple combining characters
		{"Multiple combiners", 'a', []rune{0x0300, 0x0301}, true},
		{"Empty combining", 'a', []rune{}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := needsDisplaySync(tt.mainc, tt.combc)
			if result != tt.expected {
				t.Errorf("needsDisplaySync(%c / U+%04X, %v) = %v, want %v",
					tt.mainc, tt.mainc, tt.combc, result, tt.expected)
			}
		})
	}
}

func TestNeedsDisplaySyncPerformance(t *testing.T) {
	// Performance test for ASCII characters (most frequent case)
	asciiChars := []rune{
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		' ', '!', '@', '#', '$', '%', '^', '&', '*', '(',
	}

	for _, r := range asciiChars {
		if needsDisplaySync(r, nil) {
			t.Errorf("ASCII character %c should not need display sync", r)
		}
	}
}
