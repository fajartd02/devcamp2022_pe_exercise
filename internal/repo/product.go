package product

type rng interface {
	Int63() int64
}

type Repo struct {
	rng rng
}

func New(rng rng) (*Repo, error) {
	return &Repo{
		rng: rng,
	}, nil
}
