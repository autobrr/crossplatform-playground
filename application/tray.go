package application

import (
	"context"
	"io"
	"log/slog"
	"os"

	"github.com/autobrr/crossplatform-playground/assets"

	"github.com/energye/systray"
)

var menu = make(map[string]*systray.MenuItem)

func (a *Application) startTray(ctx context.Context) {
	systray.Run(func() {
		defer os.Exit(0)
		b, _ := assets.Assets.Open(trayIcon)
		file, _ := io.ReadAll(b)
		systray.SetIcon(file)
		systray.SetTemplateIcon(file, file)
		systray.SetTooltip("Autobrr")
		systray.SetTitle("Autobrr")
		systray.SetOnRClick(a.showMenu)
		a.setupTrayMenu(ctx)
		if err := a.start(ctx); err != nil {
			slog.Error("Error starting application: %s", err)
			os.Exit(1)
		}
	}, func() {
		// stop
		slog.Info("Stopping application")
		a.Stop()

		os.Exit(0)
	})
}

func (a *Application) stopTray() {
	return
}

func (a *Application) showMenu(menu systray.IMenu) {
	if err := menu.ShowMenu(); err != nil {
		panic(err)
	}
}

func (a *Application) setupTrayMenu(ctx context.Context) {

	menu["gui"] = systray.AddMenuItem("Open WebUI", "open the web page for this Notifiarr client")
	menu["gui"].Click(func() { a.openGUI(ctx) })

	//menu["exit"] = systray.AddMenuItem("Quit", "exit "+c.Flags.Name())
	menu["exit"] = systray.AddMenuItem("Quit", "exit "+"app")
	menu["exit"].Click(func() {
		//logs.Log.Printf("Need help? %s\n=====> Exiting! User Requested", mnd.HelpLink)
		systray.Quit() // this kills the app
	})
}

//func (a *Application) start(ctx context.Context) error {
//	signal.Notify(a.sigkill, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
//	signal.Notify(a.sighup, syscall.SIGHUP)
//
//	slog.Info("Starting application")
//
//	//var err error
//	for {
//		select {
//		case sigc := <-a.sigkill:
//			slog.Info("Signal %v received", sigc)
//			return a.stop(ctx)
//		case sigc := <-a.sighup:
//			slog.Info("reload Signal %v received", sigc)
//
//		}
//
//	}
//}

func (a *Application) stop(ctx context.Context) error {
	return nil
}

func (a *Application) openGUI(ctx context.Context) {
	return
}
