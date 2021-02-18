package service

import (
	"fmt"
	"testing"
)

func TestSMSUpload_upload(t *testing.T) {
	type fields struct {
		url           string
		authorization string
		ua            string
	}
	type args struct {
		path string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantUploadPath string
		wantErr        bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			fields: fields{
				url:           "https://sm.ms/api/v2/upload",
				authorization: "mMRzVXAST9OY5kRhLQTdzPmgVHoFPqeY",
			},
			args: args{
				path: "/Users/chenjie/Downloads/movie/image/demo.md---008eGmZEly1gmmkz0fp9mj301c01c742.jpg",
			},
			wantUploadPath: "",
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SMSUpload{
				Url:           tt.fields.url,
				Authorization: tt.fields.authorization,
			}
			gotUploadPath, err := s.upload(tt.args.path)
			fmt.Sprintln(gotUploadPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("upload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotUploadPath != tt.wantUploadPath {
				t.Errorf("upload() gotUploadPath = %v, want %v", gotUploadPath, tt.wantUploadPath)
			}
		})
	}
}
