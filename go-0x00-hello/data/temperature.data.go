package data


type TemperatureModel struct {
	Name string
	Celcius float32
	Expected float32
}

var TempTestData []TemperatureModel = []TemperatureModel{
	{
		Name:     "freezing point",
		Celcius:  0,
		Expected: 32,
	},
	{
		Name:     "room temperature",
		Celcius:  25,
		Expected: 77,
	},
	{
		Name:     "body temperature",
		Celcius:  37,
		Expected: 98.6,
	},
	{
		Name:     "boiling point",
		Celcius:  100,
		Expected: 212,
	},
	{
		Name:     "negative temperature",
		Celcius:  -40,
		Expected: -40,
	},
	{
		Name:     "cold temperature",
		Celcius:  -10,
		Expected: 14,
	},
	{
		Name:     "decimal temperature",
		Celcius:  20.5,
		Expected: 68.9,
	},
	{
		Name:     "decimal negative temperature",
		Celcius:  -5.5,
		Expected: 22.1,
	},
}