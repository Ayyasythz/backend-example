package server

import (
	"sagara_backend_test/config"
	"sagara_backend_test/internal/interfaces/dao"
	"sagara_backend_test/internal/usecases"
	"sagara_backend_test/internal/usecases/wardrobe"
	"sagara_backend_test/lib/database/sql"
)

type container struct {
	Cfg        config.MainConfig
	WardrobeUc usecases.WardrobeUseCases
}

type options struct {
	Cfg *config.MainConfig
	DB  *sql.Store
}

func newContainer(opts *options) *container {
	wardrobeRepo := dao.NewWardrobeRepository(&dao.OptsWardrobeRepository{DB: opts.DB})

	wardrobeUc := wardrobe.New(&wardrobe.Opts{
		WardrobeRepo: wardrobeRepo,
		TxMgr:        nil,
	})

	return &container{
		Cfg:        *opts.Cfg,
		WardrobeUc: wardrobeUc,
	}
}
