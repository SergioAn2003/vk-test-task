package global

import "errors"

var (
	ErrDBUnvailable = errors.New("база данных недоступна")

	// ErrInternalError внутряя ошибка
	ErrInternalError = errors.New("произошла внутреняя ошибка, пожалуйста попробуйте выполнить действие позже")

	// ErrNoData данные не найдены"
	ErrNoData = errors.New("данные не найдены")

	ErrMessageWithExceededButtonTypes = errors.New("сообщение имеет кнопки с разными типами")

	ErrUnknownButtonType = errors.New("неопознанный тип кнопки")

	// ErrMessageWithInlineButtonsDontHaveContext сообщение с inline кнопками не имеет контекста
	ErrMessageWithInlineButtonsDontHaveContext = errors.New("сообщение с inline кнопками или не имеет контекста")

	// ErrUniqueConstraintViolated нарушено ограничение уникальности
	ErrUniqueConstraintViolated = errors.New("нарушено ограничение уникальности")

	ErrConnectionRefused = errors.New("connection refused")

	ErrForbiddenUser = errors.New("forbidden user")
)
