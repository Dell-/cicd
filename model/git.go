package model

type Repository struct {
	name   string
	url    string
	branch string
}

type Git struct {
	key          string
	repositories map[string]Repository
}

func (git *Git) clone() error {

	return nil
}

func (git *Git) branch(branch string) error {

	return nil
}
