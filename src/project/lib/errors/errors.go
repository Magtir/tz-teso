package errors

const (
	ERROR_NONE = 0

	ERROR_NOT_FOUND = 404
	ERROR_INTERNAL = 500
)

var AllErrors = map[int]string{
	ERROR_NOT_FOUND: "Объект не найден",
	ERROR_INTERNAL: "Внутренняя ошибка сервера",
}