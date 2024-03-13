package filmlibriary 

import (
	"github.com/Alphonnse/filmLibriary/internal/service"
)


type ImplementationLibriary struct {
	libriaryService service.ServiceLibriaryShape
}

func NewImplementationLibriary(libriaryService service.ServiceLibriaryShape) *ImplementationLibriary {
	return &ImplementationLibriary{
		libriaryService: libriaryService,
	}
}
