package usecase

type UseCase string

const (
	General    UseCase = "GENERAL"
	ShortURL   UseCase = "SHORT_URL"
	GettingURL UseCase = "GETTING_URL"
	DetailsURL UseCase = "DETAILS_URL"
	ToggleURL  UseCase = "TOGGLE_URL"
	UpdateURL  UseCase = "UPDATE_URL"
)

func (u UseCase) String() string {
	return string(u)
}
