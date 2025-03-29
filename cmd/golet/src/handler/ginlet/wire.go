//go:build !wireinject
// +build !wireinject

package ginlet

import (
	"github.com/google/wire"
)

//go:generate wire
func InitService() Service {
	wire.Build(NewService)
	return nil
}
