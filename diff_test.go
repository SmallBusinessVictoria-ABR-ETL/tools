package sbv_abr_etl

import (
	"os"
	"testing"
)

func Test_diff(t *testing.T) {

	f1, _ := os.Open("/home/jason/Downloads/VIC180716_ABR_Agency_Data.txt")
	f2, _ := os.Open("/home/jason/Downloads/VIC180716_ABR_Agency_Data_new.txt")

	type args struct {
		one *os.File
		two *os.File
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test", args{f1, f2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Diff(tt.args.one, tt.args.two)
		})
	}
}
