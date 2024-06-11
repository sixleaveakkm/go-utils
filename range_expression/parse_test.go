package range_expression

import (
	"reflect"
	"testing"
)

func TestParseToList(t *testing.T) {
	type args struct {
		exp string
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			args: args{
				exp: "1-3,5,7-9",
			},
			want:    []int{1, 2, 3, 5, 7, 8, 9},
			wantErr: false,
		}, {
			args: args{
				exp: "1-3,5,7-9,",
			},
			want:    []int{1, 2, 3, 5, 7, 8, 9},
			wantErr: false,
		}, {
			args: args{
				exp: "1-5,2,3-9",
			},
			want:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			wantErr: false,
		}, {
			args: args{
				exp: "1",
			},
			want:    []int{1},
			wantErr: false,
		}, {
			args: args{
				exp: "1,3,9",
			},
			want:    []int{1, 3, 9},
			wantErr: false,
		}, {
			args: args{
				exp: "0-6",
			},
			want:    []int{0, 1, 2, 3, 4, 5, 6},
			wantErr: false,
		}, {
			args: args{
				exp: "6-0,1-4",
			},
			wantErr: true,
		}, {
			args: args{
				exp: "a,b,c",
			},
			wantErr: true,
		}, {
			args: args{
				exp: "1~5",
			},
			wantErr: true,
		}, {
			args: args{
				exp: "1-f",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.args.exp, func(t *testing.T) {
			got, err := Parse(tt.args.exp)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
