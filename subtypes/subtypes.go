package subtypes

// entitySruct - содержит имя, тип и размер папки/файла
type EntityStruct struct {
	Name          string `json:"Имя"`    //Имя объекта
	EntityType    string `json:"Тип"`    //Тип объекта
	Size          int64  `json:"-"`      //Размер объекта в байтах
	SizeFormatted string `json:"Размер"` //Форматированный размер объекта
}

// envParam - переменная окружения
type EnvParam struct {
	Port string //значение
}

var ConfigParam EnvParam

const Asc = "asc"       //флаг сортировки по возрастанию
const Desc = "desc"     //флаг сортировки по убыванию
const MemoryBase = 1000 //основание конвертации памяти
