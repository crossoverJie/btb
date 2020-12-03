package io

import (
	"reflect"
	"testing"
)

func TestReadFileAllContext(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"1", args{"/Users/chenjie/Documents/Hexo/source/_posts/volatile.md"}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadFileAllContext(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFileAllContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("ReadFileAllContext() got = %v, want %v", *got, tt.want)
			}
		})
	}
}

func TestReadFileLine(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name string
		args args
	}{
		{"0", args{"/Users/chenjie/Documents/Hexo/source/_posts/volatile.md"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReadFileLine(tt.args.filePath)
		})
	}
}
