package usecase

// UseCase controller -> UseCase -> repo -> entity
type UseCase struct {
	//mqtt    MQTTApiRepo
}

func New() *UseCase {
	return &UseCase{
		//mqtt: m,
	}
}
