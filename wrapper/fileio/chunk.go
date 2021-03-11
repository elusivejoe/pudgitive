package fileio

type Chunk struct {
	Payload []byte
	HasNext bool
}
