package service

import "testing"

func TestReplace(t *testing.T) {
	var picMapping = make(map[string]map[string]string)
	picMapping["/Users/chenjie/Downloads/movie/demo.md"] = map[string]string{
		"https://tva1.sinaimg.cn/large/008eGmZEly1gnea76kznyj30mq0h2ta5.jpg": "newjpg1",
		"https://tva1.sinaimg.cn/large/008eGmZEly1gnea7db8kuj30lk0jqq8a.jpg": "newjpg2",
	}

	type args struct {
		picMapping *map[string]map[string]string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{picMapping: &picMapping},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Replace(tt.args.picMapping)
		})
	}
}
