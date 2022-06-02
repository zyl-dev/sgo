package uuid

import "testing"

func TestNewUlidSeq(t *testing.T) {
	type args struct {
		a int64
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test1"},
		{name: "test2"},
		{name: "test3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUlidSeq()
			t.Logf("NewUlidSeq() = %v", got)
		})
	}
}

func TestNewXidSeq(t *testing.T) {
	type args struct {
		a int64
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test1"},
		{name: "test2"},
		{name: "test3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewXidSeq()
			t.Logf("NewXidSeq() = %v", got)
		})
	}
}

func TestGenerateNewUuid(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{a: 25}, want: 25},
		{name: "test2", args: args{a: 25}, want: 25},
		{name: "test3", args: args{a: 30}, want: 30},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateNewUuid(tt.args.a)
			t.Logf("GenerateNewUuid() = %v", got)
			if len(got) != tt.want {
				t.Errorf("length of GenerateNewUuid() = %v, want %v", len(got), tt.want)
			}
		})
	}
}

func TestNewCode(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{a: 4}, want: 4},
		{name: "test2", args: args{a: 6}, want: 6},
		{name: "test3", args: args{a: 10}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCode(tt.args.a)
			t.Logf("NewCode() = %v", got)
			if len(got) != tt.want {
				t.Errorf("length of NewCode() = %v, want %v", len(got), tt.want)
			}
		})
	}
}
