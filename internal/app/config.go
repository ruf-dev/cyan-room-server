// Code generated by RedSock CLI. DO NOT EDIT.

package app

import (
	"context"

	"go.redsock.ru/rerrors"
	"go.redsock.ru/toolbox/closer"

	"go.redsock.ru/ruf/cyan-room/internal/config"
)

func (a *App) InitConfig() (err error) {
	a.Ctx, a.Stop = context.WithCancel(context.Background())
	closer.Add(func() error { a.Stop(); return nil })

	a.Cfg, err = config.Load()
	if err != nil {
		return rerrors.Wrap(err, "error reading config")
	}

	return nil
}