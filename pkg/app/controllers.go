package app

import (
	"exchange_rate/pkg/controllers"
	"exchange_rate/pkg/packages/errors"
)

func createControllers(useCases *UseCases, serverURL, basicValCode string) (*controllers.Controllers, *errors.Error) {
	return controllers.NewControllers(useCases, serverURL, basicValCode)
}
