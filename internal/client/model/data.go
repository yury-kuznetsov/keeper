package model

import "github.com/google/uuid"

const (
	// CategoryCredentials Логины и пароли
	CategoryCredentials = iota
	// CategoryText Текстовые заметки
	CategoryText
	// CategoryBinary Бинарные данные
	CategoryBinary
	// CategoryCard Банковские карты
	CategoryCard
)

// Data базовая структура для хранения локальных записей
type Data struct {
	ID       uuid.UUID
	Category int
	Data     []byte
	Version  int
}

// DataVersion структура для сверки версий с сервером
type DataVersion struct {
	ID            uuid.UUID
	VersionLocal  int
	VersionRemote int
}

// DataVersionRemote структура для описания версии записи на сервере
type DataVersionRemote struct {
	ID      uuid.UUID
	Version int
}

// DataCredentials структура для добавления логинов и паролей
type DataCredentials struct {
	Login    string
	Password string
}

// DataText структура для хранения текстовых заметок
type DataText struct {
	Text string
}

// DataBinary структура для хранения бинарных данных
type DataBinary struct {
	Binary []byte
}

// DataCard структура для хранения банковских карт
type DataCard struct {
	Number string
	Owner  string
	CVV    uint
}
