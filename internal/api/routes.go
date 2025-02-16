package api

import (
	"tribe-payments-wallet-golang-interview-assignment/internal/api/httpv1"

	"github.com/go-chi/chi/v5"
	"github.com/sumup-oss/go-pkgs/logger"
)

func RegisterRoutes(
	mux *chi.Mux,
	log logger.StructuredLogger,
) {
	mux.Get("/live", Health)

	mux.Route("/v1", func(r chi.Router) {
		r.Post("/wallet", httpv1.NewCreateWalletHandler(log))
		r.Route("/wallet/{walletID}", func(r chi.Router) {
			r.Get("/", httpv1.NewGetWalletHandler(log))
			r.Post("/deposit", httpv1.NewDepositToWalletHandler(log))
			r.Post("/withdraw", httpv1.NewWithDrawFromWalletHandler(log))
		})
	})
}
