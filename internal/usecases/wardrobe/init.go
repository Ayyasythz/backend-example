package wardrobe

import (
	"sagara_backend_test/internal/domain/repository"
	"sagara_backend_test/internal/usecases"
	"sagara_backend_test/lib/txmanager"
)

type Module struct {
	wardrobeRepo repository.WardrobeRepository
	txMgr        txmanager.TxManager
}

type Opts struct {
	WardrobeRepo repository.WardrobeRepository
	TxMgr        txmanager.TxManager
}

func New(opts *Opts) usecases.WardrobeUseCases {
	return &Module{
		wardrobeRepo: opts.WardrobeRepo,
		txMgr:        opts.TxMgr,
	}
}
