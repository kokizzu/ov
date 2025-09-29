package oviewer

import (
	"reflect"
	"regexp"
	"testing"
)

func Test_max(t *testing.T) {
	t.Parallel()
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{a: 1, b: 0},
			want: 1,
		},
		{
			name: "test2",
			args: args{a: 1, b: 2},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := max(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_min(t *testing.T) {
	t.Parallel()
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{a: 1, b: 0},
			want: 0,
		},
		{
			name: "test2",
			args: args{a: 1, b: 2},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := min(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeStr(t *testing.T) {
	t.Parallel()
	type args struct {
		list []string
		s    string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test1",
			args: args{
				list: []string{"a", "b", "c"},
				s:    "c",
			},
			want: []string{"a", "b"},
		},
		{
			name: "testZero",
			args: args{
				list: []string{},
				s:    "c",
			},
			want: []string{},
		},
		{
			name: "noRemove",
			args: args{
				list: []string{"a", "b", "c"},
				s:    "",
			},
			want: []string{"a", "b", "c"},
		},
		{
			name: "doubleRemove",
			args: args{
				list: []string{"a", "b", "b", "c"},
				s:    "b",
			},
			want: []string{"a", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := remove(tt.args.list, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeInt(t *testing.T) {
	t.Parallel()
	type args struct {
		list []int
		c    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				list: []int{1, 2, 3},
				c:    3,
			},
			want: []int{1, 2},
		},
		{
			name: "testZero",
			args: args{
				list: []int{},
				c:    3,
			},
			want: []int{},
		},
		{
			name: "noRemove",
			args: args{
				list: []int{1, 2, 3},
				c:    4,
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := remove(tt.args.list, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toAddTop(t *testing.T) {
	t.Parallel()
	type args struct {
		list []string
		s    string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test1",
			args: args{
				list: []string{"a", "b", "c"},
				s:    "f",
			},
			want: []string{"f", "a", "b", "c"},
		},
		{
			name: "test2",
			args: args{
				list: []string{},
				s:    "f",
			},
			want: []string{"f"},
		},
		{
			name: "testNoAdd",
			args: args{
				list: []string{"a", "b", "c"},
				s:    "",
			},
			want: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := toAddTop(tt.args.list, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toTop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toAddLast(t *testing.T) {
	t.Parallel()
	type args struct {
		list []string
		s    string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test1",
			args: args{
				list: []string{
					"a",
				},
				s: "b",
			},
			want: []string{
				"a",
				"b",
			},
		},
		{
			name: "test2",
			args: args{
				list: []string{
					"a",
					"b",
				},
				s: "a",
			},
			want: []string{
				"a",
				"b",
			},
		},
		{
			name: "testNoAdd",
			args: args{
				list: []string{
					"a",
					"b",
				},
				s: "",
			},
			want: []string{
				"a",
				"b",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := toAddLast(tt.args.list, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toAddLast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toLast(t *testing.T) {
	t.Parallel()
	type args struct {
		list []string
		s    string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test1",
			args: args{
				list: []string{"a", "b", "c"},
				s:    "a",
			},
			want: []string{"b", "c", "a"},
		},
		{
			name: "testAdd",
			args: args{
				list: []string{"a", "b", "c"},
				s:    "x",
			},
			want: []string{"a", "b", "c", "x"},
		},
		{
			name: "testNoAdd",
			args: args{
				list: []string{"a", "b", "c"},
				s:    "",
			},
			want: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := toLast(tt.args.list, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toLast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_remove(t *testing.T) {
	t.Parallel()
	type args struct {
		list []string
		s    string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test1",
			args: args{
				list: []string{
					"a",
				},
				s: "b",
			},
			want: []string{
				"a",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := remove(tt.args.list, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_allIndex(t *testing.T) {
	t.Parallel()
	type args struct {
		s      string
		substr string
		reg    *regexp.Regexp
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{
				s:      "a,b,c",
				substr: ",",
				reg:    nil,
			},
			want: [][]int{
				{1, 2},
				{3, 4},
			},
		},
		{
			name: "test2",
			args: args{
				s:      "a|b|c",
				substr: "|",
				reg:    nil,
			},
			want: [][]int{
				{1, 2},
				{3, 4},
			},
		},
		{
			name: "testTab",
			args: args{
				s:      "a	b	c",
				substr: "	",
				reg:    nil,
			},
			want: [][]int{
				{1, 2},
				{3, 4},
			},
		},
		{
			name: "testUnicode",
			args: args{
				s:      "a│b│c",
				substr: "│",
				reg:    nil,
			},
			want: [][]int{
				{1, 4},
				{5, 8},
			},
		},
		{
			name: "testRegex",
			args: args{
				s:      "a  b c",
				substr: `/\s+/`,
				reg:    regexp.MustCompile(`\s+`),
			},
			want: [][]int{
				{1, 3},
				{4, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := allIndex(tt.args.s, tt.args.substr, tt.args.reg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_allStringIndex(t *testing.T) {
	t.Parallel()
	type args struct {
		s      string
		substr string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{
				s:      "a,b,c",
				substr: ",",
			},
			want: [][]int{
				{1, 2},
				{3, 4},
			},
		},
		{
			name: "testNone",
			args: args{
				s:      "a,b,c",
				substr: "@",
			},
			want: nil,
		},
		{
			name: "testNoSubstr",
			args: args{
				s:      "a,b,c",
				substr: "",
			},
			want: nil,
		},
		{
			name: "testDoubleQuote",
			args: args{
				s:      `a,"b,c",d`,
				substr: ",",
			},
			want: [][]int{
				{1, 2},
				{7, 8},
			},
		},
		{
			name: "testDoubleQuote2",
			args: args{
				s:      `a,"060  ",d`,
				substr: ",",
			},
			want: [][]int{
				{1, 2},
				{9, 10},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := allStringIndex(tt.args.s, tt.args.substr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_abs(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "positive number",
			n:    5,
			want: 5,
		},
		{
			name: "negative number",
			n:    -7,
			want: 7,
		},
		{
			name: "zero",
			n:    0,
			want: 0,
		},
		{
			name: "large positive",
			n:    123456,
			want: 123456,
		},
		{
			name: "large negative",
			n:    -987654,
			want: 987654,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := abs(tt.n)
			if got != tt.want {
				t.Errorf("abs() = %v, want %v", got, tt.want)
			}
		})
	}
}
