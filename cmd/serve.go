package cmd

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"slices"
	"syscall"
	"time"

	"github.com/spf13/cobra"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/driver"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/interfaces/html"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/interfaces/restapi"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/interfaces/worker"
)

type httpHandlerMerger struct {
	api  http.Handler
	html http.Handler
}

func (m *httpHandlerMerger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if slices.Contains(html.PathList(), r.URL.Path) {
		m.html.ServeHTTP(w, r)
	} else {
		m.api.ServeHTTP(w, r)
	}
}

func newServeCommand() *cobra.Command {
	var r driver.ServerRegistry
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Start the server",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			var err error
			r, err = driver.NewServerRegistry(cmd.Context())
			if err != nil {
				return err
			}
			return setup(cmd.Context(), r)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, stop := signal.NotifyContext(cmd.Context(), os.Interrupt, syscall.SIGTERM)
			defer stop()

			state, err := r.ChargeService().RefreshChargeState(ctx, r.AppConfig().TeslaVIN)
			if err != nil {
				return err
			}
			r.Logger().Info(
				"vehicle state",
				slog.String("ChargePortLatch", state.ChargePortLatch),
				slog.Int("BatteryLevel", state.BatteryLevel),
				slog.Bool("ChargeEnableRequest", state.ChargeEnableRequest),
				slog.Int("ChargeAmps", state.ChargeAmps),
			)

			w := worker.New(r)
			go func() {
				r.Logger().Info("starting surplus metrics collector worker")
				if err := w.RunSurplusPowerCollector(ctx); err != nil {
					r.Logger().Error(err.Error())
				}
			}()
			go func() {
				r.Logger().Info("starting plug-in detection worker")
				if err := w.RunPlugInWatcher(ctx); err != nil {
					r.Logger().Error(err.Error())
				}
			}()

			srv := http.Server{
				Addr: fmt.Sprintf(":%d", r.ServerConfig().Port),
				Handler: &httpHandlerMerger{
					api:  restapi.NewHandler(r),
					html: html.NewHandler(r),
				},
				ReadTimeout:  5 * time.Second,
				WriteTimeout: 10 * time.Second,
				IdleTimeout:  120 * time.Second,
			}
			go func() {
				r.Logger().Info(fmt.Sprintf("starting server 0.0.0.0:%d", r.ServerConfig().Port))
				if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					r.Logger().Error(err.Error())
				} else {
					r.Logger().Info("shut down server successfully")
				}
			}()

			<-ctx.Done()
			r.Logger().Info("shutting down server")

			shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()
			if err := srv.Shutdown(shutdownCtx); err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}
