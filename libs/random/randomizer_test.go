package random

import "testing"

func TestRandomString(t *testing.T) {
	type args struct {
		str    []byte
		offset int
		seed   uint64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "successTestRandom",
			args: args{
				str:    make([]byte, 27),
				offset: 7,
				seed:   0,
			},
		},
	}

	const fixLen int = 10000
	randStrings := make(map[string]struct{}, fixLen)
	prefix := []byte("FARAWAY")
	for _, tt := range tests {
		for i := 0; i < fixLen; i++ {
			copy(tt.args.str[:tt.args.offset], prefix)
			RandomString(tt.args.str, tt.args.offset, tt.args.seed)
			tt.args.seed++
			rndStr := string(tt.args.str)
			randStrings[rndStr] = struct{}{}
		}
		t.Run(tt.name, func(t *testing.T) {
			if len(randStrings) != fixLen {
				t.Errorf("RandomNumber() = %v, want %v", len(randStrings), fixLen)
			}
		})
	}

}
