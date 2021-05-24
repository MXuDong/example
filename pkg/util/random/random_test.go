package random

import "testing"

func TestAnalyzePercentage(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "common right1",
			args:    args{input: "123"},
			want:    "123",
			wantErr: false,
		},
		{
			name:    "common right2",
			args:    args{input: "1%123"},
			want:    "123",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AnalyzePercentage(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("AnalyzePercentage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AnalyzePercentage() got = %v, want %v", got, tt.want)
			}
		})
	}
}
