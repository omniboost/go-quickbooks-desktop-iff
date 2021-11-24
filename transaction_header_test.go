package qdiif_test

import (
	"testing"

	"github.com/PumpkinSeed/iif"
	qdiif "github.com/omniboost/go-quickbooks-desktop-iif"
)

func TestTransactionHeader(t *testing.T) {
	data := []iif.DataLine{
		qdiif.TransactionHeader{
			TransactionID: 12,
		},
	}
	err := iif.Export(data, "tmp.iif")
	if err != nil {
		t.Error(err)
	}
}
