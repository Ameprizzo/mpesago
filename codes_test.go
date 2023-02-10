package mpesa

import (
	"sync"
	"testing"
)

func TestResponseCode(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test success",
			args: args{
				code: "INS-0",
			},
			want: "Request processed successfully",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ResponseCode(tt.args.code); got != tt.want {
				t.Errorf("ResponseCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponseCodeConc(t *testing.T) {
	codes := []string{
		"INS-0",
		"INS-1",
		"INS-6",
		"INS-9",
		"INS-13",
		"INS-20",
		"INS-21",
		"INS-28",
		"INS-996",
		"INS-997",
		"INS-998",
	}

	wg := &sync.WaitGroup{}
	//run each test in own goroutine
	for _, code := range codes {
		wg.Add(1)
		k := code
		v := responseCodes[k]
		go func(code, value string) {

			if got := ResponseCode(code); got != value {
				t.Errorf("ResponseCode() = %v, want %v", got, value)
			}

			t.Logf("ResponseCode() = %v\n", value)
			wg.Done()
		}(k, v)
	}
	wg.Wait()
}
