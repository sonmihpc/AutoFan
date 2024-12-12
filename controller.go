package main

import (
	"log"
	"sync"
	"time"
)

type Mode int

const (
	MaxMode  Mode = 0
	MeanMode Mode = 1
)

type Controller struct {
	interval   int
	mode       Mode
	pwdIds     []int
	executor   Executor
	devices    GpuDevicesInterface
	thresholds Thresholds

	lastDutyCycle int
}

func (c *Controller) Execute(dutyCycle int) {
	var wg sync.WaitGroup
	for _, pwdId := range c.pwdIds {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			if err := c.executor.Execute(id, dutyCycle); err != nil {
				log.Printf("Unable to set fan: id=%v duty-cycle=%v", id, dutyCycle)
			}
		}(pwdId)
	}
	wg.Wait()
}

func (c *Controller) Run() {
	ticker := time.NewTicker(time.Second * time.Duration(c.interval))
	for {
		select {
		case <-ticker.C:
			var temp int
			if c.mode == MaxMode {
				temp = c.devices.GetMaxTemperature()
			} else {
				temp = c.devices.GetMeanTemperature()
			}
			dutyCycle := c.thresholds.GetDutyCycleFromTemperature(temp)
			if dutyCycle != c.lastDutyCycle {
				c.Execute(dutyCycle)
				log.Printf("Temperature=%v Celsius and setup duty-cycle=%v%%", temp, dutyCycle)
				c.lastDutyCycle = dutyCycle
			} else {
				log.Printf("DutyCycle did not change")
			}
		default:
			time.Sleep(time.Second * time.Duration(c.interval))
		}
	}
}

func NewController(interval int, mode Mode, pwdIds []int, executor Executor, devices GpuDevicesInterface, thresholds Thresholds) *Controller {
	return &Controller{
		interval:      interval,
		mode:          mode,
		pwdIds:        pwdIds,
		executor:      executor,
		devices:       devices,
		thresholds:    thresholds,
		lastDutyCycle: 0,
	}
}

func NewControllerFromConfig(c Config) *Controller {
	var mode Mode
	if c.Mode == "max" {
		mode = MaxMode
	} else {
		mode = MeanMode
	}
	var executor Executor
	if c.IpmiDebug {
		executor = &FakeExecutor{}
	} else {
		executor = &IPMIExecutor{}
	}
	var devices GpuDevicesInterface
	if c.GpuDebug {
		devices = &FakeGpuDevices{}
	} else {
		devices = &GpuDevices{}
	}
	thresholds := Thresholds{
		points: c.Thresholds,
	}

	return NewController(c.Interval, mode, c.PwdIds, executor, devices, thresholds)
}
