package verification

import (
	"testing"
)

func TestVerifyPOW(t *testing.T) {
	verification, err := NewPOWPOWVerifier(3)
	if err != nil {
		t.Errorf("test failed - %s", err)
	}

	type args struct {
		challenge []byte
		solution  []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "successTestHash",
			args: args{
				challenge: []byte("FARAWAY"),
				solution:  []byte("FARAWAYAPAxk0SCTXMmspn0Hclw"),
			},
			want: true,
		}, {
			name: "failTestHash",
			args: args{
				challenge: []byte("FARAWAY"),
				solution:  []byte("FARAWAYAPAxk0SCTXMmspn0Hdclw"),
			},
			want: false,
		}, {
			name: "failTestHash",
			args: args{
				challenge: []byte("asd"),
				solution:  []byte("asdfsdfdsfdsfsd"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := verification.Verify(tt.args.challenge, tt.args.solution); got != tt.want {
				t.Errorf("VerifyPOW() = %v, want %v", got, tt.want)
			}
		})
	}
}
