package main

import "testing"

func TestThresholds_GetDutyCycleFromTemperature(t1 *testing.T) {
	ps := []Threshold{
		{
			Temperature: 50,
			DutyCycle:   50,
		},
		{
			Temperature: 70,
			DutyCycle:   70,
		},
		{
			Temperature: 80,
			DutyCycle:   80,
		},
		{
			Temperature: 90,
			DutyCycle:   100,
		},
	}
	type fields struct {
		points []Threshold
	}
	type args struct {
		temp int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "low",
			fields: fields{
				points: ps,
			},
			args: args{
				temp: 49,
			},
			want: 50,
		},
		{
			name: "middle",
			fields: fields{
				points: ps,
			},
			args: args{
				temp: 71,
			},
			want: 70,
		},
		{
			name: "high",
			fields: fields{
				points: ps,
			},
			args: args{
				temp: 95,
			},
			want: 100,
		},
		{
			name: "highest",
			fields: fields{
				points: ps,
			},
			args: args{
				temp: 101,
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Thresholds{
				points: tt.fields.points,
			}
			if got := t.GetDutyCycleFromTemperature(tt.args.temp); got != tt.want {
				t1.Errorf("GetDutyCycleFromTemperature() = %v, want %v", got, tt.want)
			}
		})
	}
}
