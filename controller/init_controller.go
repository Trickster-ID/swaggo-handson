package controller

import "context"

func InitHttpUserController(ctx context.Context) UserControllerInterface {
	return InitUserController(ctx)
}
