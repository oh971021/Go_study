package main

import "fmt"

type TapeInterface interface {
	Play(string)
	Stop()
}

type TapePlayer struct {
	Batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("playing", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("stopped")
}

type TapeRecorder struct {
	Microphones int
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("Recording", song)
}

func (t TapeRecorder) Record() {
	fmt.Println("Recording")
}

func (t TapeRecorder) Stop() {
	fmt.Println("Stopped!!")
}

func playList(device TapeInterface, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	recorder, ok := device.(TapeRecorder)
	if ok {
		recorder.Record()
	}
	device.Stop()
}

func main() {
	player := TapePlayer{}
	record := TapeRecorder{}
	mixtape := []string{"first", "second", "third"}
	playList(player, mixtape)
	playList(record, mixtape)
}
