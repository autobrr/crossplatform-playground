//go:build windows || darwin

package application

import "context"

func (a *Application) Start(ctx context.Context) error {
	a.startTray(ctx)
	return nil
}
