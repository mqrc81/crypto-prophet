package bot

import (
	"fmt"

	"github.com/adshao/go-binance/v2"
)

func NewBuyOrder(signal Signal) (Signal, error) {
	response, err := client.NewCreateOrderService().
		Symbol(signal.CryptoBase).Side(binance.SideTypeBuy).Type(binance.OrderTypeStopLoss).Quantity(str(signal.Amount)).
		StopPrice(str(signal.Increments[0])).NewOrderRespType(binance.NewOrderRespTypeRESULT).
		Do(ctx)

	if response == nil || response.ClientOrderID == "" {
		return signal, err
	}

	signal.ID = response.ClientOrderID
	return signal, nil
}

func str(float float64) string {
	return fmt.Sprintf("%f", float)
}
