package nets

import "testing"

func TestGetHttpUrl(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 string
	}{
		{
			name:  "complete url with protocol and item",
			args:  args{url: "http://www.example.com"},
			want:  true,
			want1: "http://www.example.com",
		},
		{
			name:  "without protocol 1",
			args:  args{url: "www.example.com"},
			want:  false,
			want1: "http://www.example.com",
		},
		{
			name:  "without protocol 2",
			args:  args{url: "://www.example.com"},
			want:  false,
			want1: "http://www.example.com",
		},
		{
			name:  "error input1",
			args:  args{url: ""},
			want:  false,
			want1: "",
		},
		{
			name:  "error input2",
			args:  args{url: "error://error://"},
			want:  false,
			want1: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetHttpUrl(tt.args.url)
			if got != tt.want {
				t.Errorf("GetHttpUrl() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetHttpUrl() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
