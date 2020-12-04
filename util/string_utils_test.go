package util

import (
	"reflect"
	"testing"
)

func TestGetFileSuffix(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"0", args{"https://i.loli.net/2020/07/28/q74soVAQbl5NyK6.jpg"}, "q74soVAQbl5NyK6.jpg"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFileSuffix(tt.args.fileName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFileSuffix() = %v, want %v", got, tt.want)
			}
		})
	}
}
