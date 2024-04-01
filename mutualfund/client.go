package mutualfund

const (
	// MutualFundLatestBaseURL is https://api.mfapi.in/mf/%s/latest
	MutualFundLatestBaseURL = "https://api.mfapi.in/mf/%s/latest"
	// MutualFundHistoryBaseURL is https://api.mfapi.in/mf/%s
	MutualFundHistoryBaseURL = "https://api.mfapi.in/mf/%s"
)

type Handler struct {
	// log *slog.Logger
}

// func NewHandler(l *slog.Logger) *Handler {
func NewHandler() *Handler {
	return &Handler{
		// log: l,
	}
}
