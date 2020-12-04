package util

import "testing"

func TestMatchPicture(t *testing.T) {
	type args struct {
		context string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"1", args{"![](https://i.loli.net/2020/10/09/MbZHxvDFeyhmAtQ.jpg)"}, "https://i.loli.net/2020/10/09/MbZHxvDFeyhmAtQ.jpg", false},
		{"2", args{"123"}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MatchPicture(&tt.args.context)
			if (err != nil) != tt.wantErr {
				t.Errorf("MatchPicture() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if *got != tt.want {
				t.Errorf("MatchPicture() got = %v, want %v", *got, tt.want)
			}
		})
	}
}
