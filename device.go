package main

import (
	"github.com/NVIDIA/go-nvml/pkg/nvml"
	"log"
	"math/rand/v2"
)

type GpuDevice struct {
	Id               int
	Name             string
	UtilizationRates int
	Temperature      int
}

type GpuDevicesInterface interface {
	updateDevices()
	GetMaxTemperature() int
	GetMeanTemperature() int
}

type GpuDevices struct {
	devices []GpuDevice
}

func (g *GpuDevices) updateDevices() {
	ret := nvml.Init()
	if ret != nvml.SUCCESS {
		log.Printf("Unable to initial NVML: %v", nvml.ErrorString(ret))
		return
	}
	defer func() {
		ret := nvml.Shutdown()
		if ret != nvml.SUCCESS {
			log.Printf("Unable to shutdown NVML: %v", nvml.ErrorString(ret))
		}
	}()

	count, ret := nvml.DeviceGetCount()
	if ret != nvml.SUCCESS {
		log.Printf("Unable to get deivce count: %v", nvml.ErrorString(ret))
		return
	}

	devices := make([]GpuDevice, 0)

	for i := 0; i < count; i++ {
		device, ret := nvml.DeviceGetHandleByIndex(i)
		if ret != nvml.SUCCESS {
			log.Printf("Unable to get device at index: %d: %v", i, nvml.ErrorString(ret))
			continue
		}

		name, ret := device.GetName()
		if ret != nvml.SUCCESS {
			log.Printf("Unable to get device name: %v", nvml.ErrorString(ret))
			continue
		}

		utilization, ret := device.GetUtilizationRates()
		if ret != nvml.SUCCESS {
			log.Printf("Unable to get device utilization: %v", nvml.ErrorString(ret))
			continue
		}

		temp, ret := device.GetTemperature(nvml.TEMPERATURE_GPU)
		if ret != nvml.SUCCESS {
			log.Printf("Unable to get device temperatur: %v", nvml.ErrorString(ret))
			continue
		}

		devices = append(devices, GpuDevice{
			Id:               i,
			Name:             name,
			UtilizationRates: int(utilization.Gpu),
			Temperature:      int(temp),
		})
	}
	g.devices = devices
}

func (g *GpuDevices) GetMaxTemperature() int {
	g.updateDevices()
	if len(g.devices) == 0 {
		log.Println("Unable to get device temperature: the device count == 0")
		return 0
	}

	maxTemp := 0
	for _, device := range g.devices {
		if device.Temperature > maxTemp {
			maxTemp = device.Temperature
		}
	}
	return maxTemp
}

func (g *GpuDevices) GetMeanTemperature() int {
	g.updateDevices()
	if len(g.devices) == 0 {
		log.Println("Unable to get device temperature: the device count == 0")
		return 0
	}

	total := 0
	for _, device := range g.devices {
		total += device.Temperature
	}
	return total / len(g.devices)
}

type FakeGpuDevices struct{}

func (f *FakeGpuDevices) updateDevices() {
	log.Printf("Fake GPU Devices Update ")
}

func (f *FakeGpuDevices) GetMaxTemperature() int {
	return rand.IntN(101)
}

func (f *FakeGpuDevices) GetMeanTemperature() int {
	return rand.IntN(101)
}
