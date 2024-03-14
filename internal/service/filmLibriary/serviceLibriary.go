package filmlibriary

import (
	"github.com/Alphonnse/filmLibriary/internal/repository"
)

type serviceLibriary struct {
	libriaryRepository repository.Respository
}

func NewServiceLibriary(repo repository.Respository) *serviceLibriary {
	return &serviceLibriary{
		libriaryRepository: repo,
	}
}
