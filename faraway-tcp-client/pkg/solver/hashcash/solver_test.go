package hashcash

import (
	"reflect"
	"testing"
)

func TestMine(t *testing.T) {
	type args struct {
		challenge  []byte
		difficulty int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "successPOWTestMine",
			args: args{
				challenge:  []byte("FARAWAY"),
				difficulty: 3,
			},
			want: []byte("FARAWAYAPAxk0SCTXMmspn0Hclw"),
		},
	}

	for _, tt := range tests {
		solver := NewSolver(tt.args.difficulty)
		t.Run(tt.name, func(t *testing.T) {
			if got := solver.Solve(tt.args.challenge); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHash(t *testing.T) {
	type args struct {
		str        []byte
		complexity int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "successTestHash",
			args: args{
				str:        []byte("FARAWAYAPAxk0SCTXMmspn0Hclw"),
				complexity: 3,
			},
			want: true,
		}, {
			name: "failTestHash",
			args: args{
				str:        []byte("FARAWAYAPAxk0SCTXMfspn0Hclw"),
				complexity: 3,
			},
			want: false,
		}, {
			name: "failTestHash",
			args: args{
				str:        []byte("asdfsdfdsfdsfsd"),
				complexity: 3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hash(tt.args.str, tt.args.complexity); got != tt.want {
				t.Errorf("Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}
