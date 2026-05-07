package processor

type Token struct {
	Word  string
	Focus int
	Width int
	Pause PauseLength
}

type PauseLength int

const (
	Normal PauseLength = iota
	Medium
	Long
)

var PauseLenName = map[PauseLength]string{
	Normal: "normal",
	Medium: "medium",
	Long:   "long",
}

var PauseLenMult = map[PauseLength]float64{
	Normal: 1.0,
	Medium: 1.5,
	Long:   2.0,
}
