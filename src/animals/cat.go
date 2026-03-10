package animals

type Cat struct {
	Breeds []string
}

func (c *Cat) GetBreeds() []string {
	return c.Breeds
}