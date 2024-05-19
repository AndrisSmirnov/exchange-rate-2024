package rate_service

type RateService struct {
	repository Repository
}

func NewRateService(
	repository Repository,
) *RateService {
	return &RateService{
		repository: repository,
	}
}
