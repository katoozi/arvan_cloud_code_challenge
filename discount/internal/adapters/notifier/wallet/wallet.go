package wallet

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/katoozi/avran_cloud_code_challenge_discount_service/internal/ports"
)

type wallet struct {
	url            string
	requestTimeout time.Duration
}

// New is the wallet notifier factory method
func New(url string, requestTimeout time.Duration) ports.Notifier {
	return &wallet{
		url:            strings.TrimSuffix(url, "/"),
		requestTimeout: requestTimeout,
	}
}

func (walletNotifier *wallet) Transaction(ctx context.Context, cellphone string, amount uint64, description string, timestamp time.Time) error {
	// TODO: implement an skd for wallet service and used it.

	var client = http.Client{Timeout: walletNotifier.requestTimeout}

	request, err := http.NewRequestWithContext(ctx, "GET", walletNotifier.url+transactionEndpoint, nil)
	if err != nil {
		return err
	}

	queryString := request.URL.Query()
	queryString.Set("cellphone", cellphone)
	queryString.Set("amount", strconv.FormatUint(amount, 10))
	queryString.Set("timestamp", timestamp.Format(time.RFC3339))
	queryString.Set("type", "1")
	queryString.Set("description", description)

	request.URL.RawQuery = queryString.Encode()

	response, err := client.Do(request)
	if err != nil {
		return err
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("got error from wallet service. body: %s, status_code: %d", string(responseBody), response.StatusCode)
	}
	return nil
}
