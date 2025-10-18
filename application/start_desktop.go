//go:build windows || darwin

package application

func (a *Application) Start(ctx context.Context) error {
	a.startTray(ctx)
	return nil
}
