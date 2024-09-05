package unpack

import "testing"

func TestGetStr(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    `"a4bc2d5e" => "aaaabccddddde"`,
			args:    args{str: "a4bc2d5e"},
			want:    "aaaabccddddde",
			wantErr: false,
		}, {
			name:    `"abcd" => "abcd"`,
			args:    args{str: "abcd"},
			want:    "abcd",
			wantErr: false,
		}, {
			name:    `"45" => ""`,
			args:    args{str: "45"},
			want:    "",
			wantErr: true,
		}, {
			name:    `"" => ""`,
			args:    args{str: ""},
			want:    "",
			wantErr: false,
		}, {
			name:    `"qwe\4\5" => "qwe45"`,
			args:    args{str: ""},
			want:    "",
			wantErr: false,
		}, {
			name:    `"qwe\45" => "qwe44444"`,
			args:    args{str: `qwe\45`},
			want:    "qwe44444",
			wantErr: false,
		}, {
			name:    `"qwe\\5" => "qwe\\\\\"`,
			args:    args{str: `qwe\\5`},
			want:    `qwe\\\\\`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetStr(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetStr() got = %v, want %v", got, tt.want)
			}
		})
	}
}
