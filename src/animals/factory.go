package animals

// Creates animal based on what's passed into the function
func New(name string) Animal {
	switch name {
	case "cat":
		return &Cat{Breeds: []string{"Siamese", "Maine Coon"}}
	case "dog":
		return &Dog{Breeds: []string{}}
	case "cow":
		return &Cow{Breeds: []string{}}
	default:
		return nil
	}
}
