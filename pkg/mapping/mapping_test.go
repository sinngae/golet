package gomapping

import (
	"fmt"
	"testing"
)

func TestMapToStruct(t *testing.T) {
	type args struct {
		src  map[string]interface{}
		dest interface{}
		tag  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case 0",
			args: args{
				src: map[string]interface{}{
					"fa": 1,
					"fb": "this is test",
				},
				dest: &struct {
					FieldA int    `test:"field:fa,required"`
					FieldB string `test:"field:fb,omit"`
				}{
					FieldA: 0,
					FieldB: "",
				},
				tag: "test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MapToStruct(tt.args.src, tt.args.dest, tt.args.tag); (err != nil) != tt.wantErr {
				t.Errorf("MapToStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
			fmt.Println(tt.args.dest)
		})
	}
}

func TestSliceToStruct(t *testing.T) {
	type args struct {
		src  []interface{}
		dest interface{}
		tag  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case 0",
			args: args{
				src: []interface{}{
					"this is a test",
					23,
				},
				dest: &struct {
					FieldA int    `test:"idx:1,required"`
					FieldB string `test:"idx:0,omit"`
				}{
					FieldA: 0,
					FieldB: "",
				},
				tag: "test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SliceToStruct(tt.args.src, tt.args.dest, tt.args.tag); (err != nil) != tt.wantErr {
				t.Errorf("SliceToStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
			fmt.Println(tt.args.dest)
		})
	}
}
