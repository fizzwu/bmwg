package data

var data = []Kitten{
	Kitten{
		Id:     "1",
		Name:   "Sam",
		Weight: 12.1,
	},
	Kitten{
		Id:     "2",
		Name:   "Fat Freddy's Cat",
		Weight: 11.1,
	},
	Kitten{
		Id:     "3",
		Name:   "Garfield",
		Weight: 20.2,
	},
}

type MemoryStore struct {
}

func (s *MemoryStore) Search(name string) []Kitten {
	var kittens []Kitten
	for _, val := range data {
		if val.Name == name {
			kittens = append(kittens, val)
		}
	}
	return kittens
}
