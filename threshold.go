package main

const DefaultDutyCycle = 50

type Threshold struct {
	Temperature int `mapstructure:"temperature" yaml:"temperature"`
	DutyCycle   int `mapstructure:"duty-cycle" yaml:"duty-cycle"`
}

type Thresholds struct {
	points []Threshold
}

func (t *Thresholds) GetDutyCycleFromTemperature(temp int) int {
	length := len(t.points)
	if temp < t.points[0].Temperature {
		return t.points[0].DutyCycle
	}
	if temp >= t.points[length-1].Temperature {
		return t.points[length-1].DutyCycle
	}
	for i, point := range t.points {
		if temp >= point.Temperature && temp < t.points[i+1].Temperature {
			return point.DutyCycle
		}
	}
	return DefaultDutyCycle
}
