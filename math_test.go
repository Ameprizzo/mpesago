package mpesa

import "testing"

func TestTwoDecimalPlaces1(t *testing.T) {
	type args struct {
		req requestType
		f   float64
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "pushpay",
			args: args{
				req: pushPay,
				f:   12.3456789,
			},
			want: 12.35,
		},
		{
			name: "disburse",
			args: args{
				req: disburse,
				f:   12.3496789,
			},
			want: 12.35,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoDecimalPlaces(tt.args.req, tt.args.f); got != tt.want {
				t.Errorf("TwoDecimalPlaces() = %v, want %v", got, tt.want)
			}
		})
	}
}
