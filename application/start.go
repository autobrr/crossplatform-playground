//go:build !windows && !darwin

package application

import "context"

func (a *Application) Start(ctx context.Context) error {
	//if err := a.start(ctx); err != nil {
	//	return err
	//}
	a.startTray(ctx)

	return nil
}
