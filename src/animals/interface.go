package animals

// Interfacing is super important - it allows us to define a contract of what
// all Animals will do. This is a little down the line for us.
type Animal interface {
	GetBreeds() []string
}
