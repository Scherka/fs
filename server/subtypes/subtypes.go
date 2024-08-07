package subtypes

// Sort - константы сортировки
type Sort int

// числовые константы
type intConst int

// Response - тело ответа
type Response struct {
	ErrorCode     int      `json:"error_code"`      //код ошибки 0-нет ошибки 1-есть ошибка
	ErrorMessage  string   `json:"error_message"`   //описание ошибки
	Data          []Record `json:"data"`            //структура каталога
	Root          string   `json:"root"`            //root, структура которого изучалась
	FullSize      int64    `json:"full_size"`       //полный размер обрабатываемой директории
	LoadingTime   float64  `json:"loading_time"`    //время работы сканера
	DateOfRequest string   `json:"date_of_request"` //дата запроса
	TimeOfRequest string   `json:"time_of_request"` //время запроса
}

// entitySruct - содержит имя, тип и размер папки/файла
type Record struct {
	Name          string `json:"name"` //Имя объекта
	EntityType    string `json:"type"` //Тип объекта
	Size          int64  `json:"-"`    //Размер объекта в байтах
	SizeFormatted string `json:"size"` //Форматированный размер объекта
}

// envParam - переменная окружения
type EnvParam struct {
	Port              string //значение порта
	Root              string //начальная директория
	DB_INSERTER_PATH  string //запрос на recive.php
	STAT_DISPLAY_PATH string //запрос на stat.php
}

var ConfigParam EnvParam

const (
	Asc  Sort = iota //флаг сортировки по возрастанию
	Desc Sort = iota //флаг сортировки по убыванию
)

func (s Sort) String() string {
	switch s {
	case Asc:
		return "asc"
	case Desc:
		return "desc"
	}
	return ""
}

// const Asc = "asc"
// const Desc = "desc"
const (
	MemoryBase      intConst = iota //основание конвертации памяти
	Multiplier      intConst = iota //множитель для конвертации памяти и времени ожидания завершения работы сервера
	ServerErrorCode intConst = iota // ошибка сервера 500
)

func (c intConst) Value() int64 {
	switch c {
	case MemoryBase: //основание конвертации памяти
		return 1000
	case Multiplier: ////множитель для конвертации памяти и времени ожидания завершения работы сервера
		return 4
	case ServerErrorCode:
		return 500
	}
	return 0
}

//const MemoryBase = 1000 //основание конвертации памяти
//const Multiplier = 4    //множитель для конвертации памяти и времени ожидания завершения работы сервера
