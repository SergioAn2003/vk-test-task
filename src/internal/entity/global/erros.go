package global

import "errors"

var (
	ErrDBUnvailable = errors.New("база данных недоступна")

	// ErrInternalError внутряя ошибка
	ErrInternalError = errors.New("произошла внутреняя ошибка, пожалуйста попробуйте выполнить действие позже")

	// ErrNoData данные не найдены"
	ErrNoData = errors.New("данные не найдены")

	// ErrUniqueConstraintViolated нарушено ограничение уникальности
	ErrUniqueConstraintViolated = errors.New("нарушено ограничение уникальности")

	// ErrParamsIncorect неверные параметры запросва
	ErrParamsIncorect = errors.New("неверные пармаметры запроса")
)
