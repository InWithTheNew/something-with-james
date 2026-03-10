package animals
// Struct to define what a `Dog` is.
type Dog struct {
	Breeds []string
}

// Function to define what `Dog.GetBreeds` would do.
func (d *Dog) GetBreeds() []string {
	return d.Breeds
}
