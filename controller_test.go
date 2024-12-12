package main

import (
	"testing"
)

func TestNewController(t *testing.T) {
	interval := 4
	pwdIds := []int{2, 3, 4, 5}
	executor := &FakeExecutor{}
	devices := &FakeGpuDevices{}
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
	thresholds := Thresholds{
		points: ps,
	}

	controller := NewController(interval, MaxMode, pwdIds, executor, devices, thresholds)
	controller.Run()
}
