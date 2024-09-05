package ntp

import (
	"testing"
	"time"
)

func TestGetTime(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Default test",
			args:    args{address: ""},
			want:    time.Now(),
			wantErr: false,
		}, {
			name:    "Incorrect address",
			args:    args{address: "1111"},
			want:    time.Time{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTime(tt.args.address)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want.Sub(got) > time.Second {
				t.Errorf("GetTime() got = %v, want %v", got, tt.want)
			}
		})
	}
}
