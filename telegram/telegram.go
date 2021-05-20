package telegram

const (
	ErrorMSG            = "!!! [E]:\n"
	NewSignalOrderMSG   = "!! [N]:\n"
	NewTargetReachedMSG = "! [T]:\n"
	ProfitMSG           = "!!! [P]:\n"
	LossMSG             = "!!! [L]:\n"
)

// Inform myself of errors & transactions so I'm aware of errors (& profits) on
// the go in order to interfere immediately if need be
func SendMessage(message string) error {
	// TODO
	return nil
}
