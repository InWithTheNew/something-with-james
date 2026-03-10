package animals

type Cow struct {
	Breeds []string
}

func (c *Cow) GetBreeds() []string {
	return c.Breeds
}
