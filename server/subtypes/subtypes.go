package subtypes

// responSebody - тело ответа
type ResponseStruct struct {
	ErrorCode     int            //код ошибки 0-нет ошибки 1-есть ошибка
	ErrorMessage  string         //описание ошибки
	Data          []EntityStruct //структура каталога
	Root          string         //root, структура которого изучалась
	FullSize      int64          //полный размер обрабатываемой директории
	LoadingTime   string         //время работы сканера
	DateOfRequest string         //дата запроса
	TimeOfRequest string         //время запроса
}

// entitySruct - содержит имя, тип и размер папки/файла
type EntityStruct struct {
	Name          string `json:"Имя"`    //Имя объекта
	EntityType    string `json:"Тип"`    //Тип объекта
	Size          int64  `json:"-"`      //Размер объекта в байтах
	SizeFormatted string `json:"Размер"` //Форматированный размер объекта
}

// envParam - переменная окружения
type EnvParam struct {
	Port string //значение порта
	Root string //начальная директория
}

// ClearResponse - очистка результатов прошлого запроса
func ClearResponse() {
	ResponseBody.ErrorCode = 0
	ResponseBody.ErrorMessage = ""
	ResponseBody.Data = nil
	ResponseBody.Root = ""
	ResponseBody.FullSize = 0
	ResponseBody.LoadingTime = ""
	ResponseBody.TimeOfRequest = ""
	ResponseBody.DateOfRequest = ""
}

var ConfigParam EnvParam
var ResponseBody ResponseStruct

const Asc = "asc"       //флаг сортировки по возрастанию
const Desc = "desc"     //флаг сортировки по убыванию
const MemoryBase = 1000 //основание конвертации памяти
const Multiplier = 4
