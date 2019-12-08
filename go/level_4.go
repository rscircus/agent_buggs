// Objective: Set the cameras to "Idle"
package main

const foundIntruder bool = true

func main() {
	camera := online()
	status := "Idle"
	if foundIntruder == true {
		status = startRecording(camera)
	}

	// Something suspicious happened with the status code
	// so let's start recording
	if status != "Idle" && status != "Recording" {
		status = "Recording"
	}
	println("Status:", status)
}

// TODO: Change this
func online() RecordingDevice {
	return FakeCam{name: "Perimeter Camera"}
}

type FakeCam struct {
	name string
}

func (f FakeCam) record() string {
	return "Idle"
}

type RecordingDevice interface {
	record() string
}

// OK, the idea was to implement FakeCam using the RecordingDevice interface
// and then overwrite the record() function -- which then has FakeCam as
// receiver -- which is called by the interface's startRecording function

type Camera struct {
	name string
}

func startRecording(device RecordingDevice) string {
	return device.record()
}

func (c Camera) record() string {
	if foundIntruder {
		return "Recording"
	} else {
		return "Idle"
	}
}
