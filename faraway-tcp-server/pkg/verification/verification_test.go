package verification

import "testing"

func TestVerifyPOW(t *testing.T) {
	type args struct {
		challenge []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "successTestHash",
			args: args{
				challenge: []byte("FARAWAYAPAxk0SCTXMmspn0Hclw"),
			},
			want: true,
		}, {
			name: "failTestHash",
			args: args{
				challenge: []byte("FARAWAYAPAxk0SCTXMmspn0Hdclw"),
			},
			want: false,
		}, {
			name: "failTestHash",
			args: args{
				challenge: []byte("asdfsdfdsfdsfsd"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VerifyPOW(tt.args.challenge); got != tt.want {
				t.Errorf("VerifyPOW() = %v, want %v", got, tt.want)
			}
		})
	}
}
