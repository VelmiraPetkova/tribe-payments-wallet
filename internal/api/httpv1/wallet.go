package httpv1

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/sumup-oss/go-pkgs/logger"

	"tribe-payments-wallet-golang-interview-assignment/internal/wallet"
)

type WalletRequestJson struct {
	Name string `json:"name"`
}

func (w *WalletRequestJson) Bind(r *http.Request) error {
	if w.Name == "" {
		return errors.New("missing required field name.")
	}
	return nil
}

const AuthenticatedUserParameter = "X-Auth-UserId"

var walletStore = wallet.NewWalletRepository()

func NewCreateWalletHandler(log logger.StructuredLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		log.Info("CreateWalletHandler")

		var walletRequest WalletRequestJson
		err := render.Bind(r, &walletRequest)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		userId := r.Header.Get(AuthenticatedUserParameter)
		if userId == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		newWallet := wallet.CreateWallet(userId, walletRequest.Name)

		err = walletStore.SaveFakeWalet(newWallet)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		render.Status(r, http.StatusCreated)
		render.Render(w, r, newWallet)

	}
}

func NewGetWalletHandler(log logger.StructuredLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		log.Info("NewGetWalletHandler")

		walletId := chi.URLParam(r, "walletID")
		if walletId == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		userId := r.Header.Get(AuthenticatedUserParameter)
		if userId == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		wallet, err := walletStore.GetWallet(walletId, userId)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		render.Status(r, http.StatusOK)
		render.Render(w, r, wallet)
	}
}

type DepositRequestJson struct {
	DepositValue float64 `json:"depositValue"`
}

func (w *DepositRequestJson) Bind(r *http.Request) error {
	if w.DepositValue <= 0.0 {
		return errors.New("invalid deposit value")
	}
	return nil
}

func NewDepositToWalletHandler(log logger.StructuredLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		log.Info("NewDepositToWalletHandler")

		walletId := chi.URLParam(r, "walletID")
		if walletId == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		userId := r.Header.Get(AuthenticatedUserParameter)
		if userId == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		wallet, err := walletStore.GetWallet(walletId, userId)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var depositRequest DepositRequestJson
		err = render.Bind(r, &depositRequest)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		wallet.DepositMoney(depositRequest.DepositValue)

		render.Status(r, http.StatusOK)
		render.Render(w, r, wallet)
	}
}

type WithDrawRequestJson struct {
	WithDrawValue float64 `json:"withDrawValue"`
}

func (w *WithDrawRequestJson) Bind(r *http.Request) error {
	if w.WithDrawValue <= 0.0 {
		return errors.New("invalid Withdraw value")
	}
	return nil
}

func NewWithDrawFromWalletHandler(log logger.StructuredLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		log.Info("NewWithDrawFromWalletHandler")

		walletId := chi.URLParam(r, "walletID")
		if walletId == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		userId := r.Header.Get(AuthenticatedUserParameter)
		if userId == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		wallet, err := walletStore.GetWallet(walletId, userId)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var withDrawRequest WithDrawRequestJson
		err = render.Bind(r, &withDrawRequest)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = wallet.WithdrawMoney(withDrawRequest.WithDrawValue)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		render.Status(r, http.StatusOK)
		render.Render(w, r, wallet)
	}
}
