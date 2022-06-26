package factory

import (
	adapter "github.com/nakoding-community/golang-boilerplate/internal/adapter"
	"github.com/nakoding-community/golang-boilerplate/internal/database"
	"github.com/nakoding-community/golang-boilerplate/internal/usecase"

	"gorm.io/gorm"
)

type (
	Factory struct {
		ConnectionGorm *gorm.DB

		Adapter struct {
			OutBound *adapter.OutBound
			InBound  *adapter.InBound
		}

		Usecase *usecase.Usecase
	}
)

func NewFactory() *Factory {
	f := &Factory{}
	f.setupDb()
	// f.SetupModelPsqlGorm()
	f.setupAdapterOutBound()
	f.setupUseCase()
	f.setupAdapterInBound()

	return f
}

func (f *Factory) setupDb() {
	database.Init()

	conn := "goboil_hex_db"
	db, err := database.Connection[gorm.DB](conn)
	if err != nil {
		panic("Failed setup db, connection is undefined")
	}
	f.ConnectionGorm = db
}

func (f *Factory) setupAdapterOutBound() {
	f.Adapter.OutBound = adapter.NewOutBound(f.ConnectionGorm)
}

func (f *Factory) setupUseCase() {
	f.Usecase = usecase.New(
		f.Adapter.OutBound.Orm.User,
	)
}

func (f *Factory) setupAdapterInBound() {
	f.Adapter.InBound = adapter.NewInBound(
		f.Usecase.UsecaseUser,
	)
}
