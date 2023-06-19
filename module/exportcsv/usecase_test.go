package exportcsv

import (
	"reflect"
	"testing"
)

func TestTransactionStringConverter(t *testing.T) {
	type args struct {
		transactions []Transaction
	}
	tests := []struct {
		name    string
		args    args
		want    [][]string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TransactionStringConverter(tt.args.transactions)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransactionStringConverter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransactionStringConverter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_ExportCSV(t *testing.T) {
	type args struct {
		req *ExportCSVRequest
	}
	tests := []struct {
		name    string
		u       UseCase
		args    args
		want    [][]string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.ExportCSV(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.ExportCSV() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseCase.ExportCSV() = %v, want %v", got, tt.want)
			}
		})
	}
}
