package discount

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/katoozi/avran_cloud_code_challenge_discount_service/internal/discount/models"
	"github.com/katoozi/avran_cloud_code_challenge_discount_service/internal/ports"
)

// RedeemHandler will inject dependencies and return the redeem http handler
func RedeemHandler(persistantLayer ports.PersistantLayer, notifier ports.Notifier) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		voucherCode := req.URL.Query().Get("voucher-code")
		cellphone := req.URL.Query().Get("cellphone")

		discountPrice, err := persistantLayer.RedeemVoucher(cellphone, voucherCode)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte(err.Error()))
			return
		}

		if err := notifier.Transaction(req.Context(), cellphone, discountPrice, fmt.Sprintf("redeem the %s", voucherCode), time.Now()); err != nil {
			// TODO: undo the redemption
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte(err.Error()))
			return
		}

		res.WriteHeader(http.StatusOK)
		res.Write([]byte("operation is successfull"))
	}
}

// CreateVoucherHandler will inject dependencies and return the create voucher http handler
func CreateVoucherHandler(persistantLayer ports.PersistantLayer, notifier ports.Notifier) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		voucherCode := req.URL.Query().Get("voucher-code")

		// TODO: handle the parse error
		maxUseTime, _ := strconv.ParseUint(req.URL.Query().Get("max-use-time"), 10, 64)

		// TODO: handle the parse error
		discountPrice, _ := strconv.ParseUint(req.URL.Query().Get("discount-price"), 10, 64)

		err := persistantLayer.SaveVoucher(models.Voucher{
			Code:          voucherCode,
			DiscountPrice: discountPrice,
			MaxUseTime:    maxUseTime,
			CreatedAt:     time.Now(),
		})
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte(err.Error()))
			return
		}

		res.WriteHeader(http.StatusOK)
		res.Write([]byte("operation is successfull"))
	}
}

// VoucherRedemptionsHandler will inject dependencies and return the voucher redemption records http handler
func VoucherRedemptionsHandler(persistantLayer ports.PersistantLayer, notifier ports.Notifier) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		voucherCode := req.URL.Query().Get("voucher-code")

		voucherRedemptions, err := persistantLayer.VoucherRedemptionRecords(voucherCode)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte(err.Error()))
			return
		}

		json.NewEncoder(res).Encode(voucherRedemptions)
	}

}
