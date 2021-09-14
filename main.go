package main

/*
Етап №1:
створити сервіс (REST API) органайзера-калердаря з функціональністю:
- додавати події, нагадування.
- редагувати їх, змінювати назву, час, опис...
- видаляти події
- переглядати перелік подій на день, тиждень, місяць, рік (з пітримкою фільтрації по ознакам)
*/

type responseCode struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
}

type Date struct {
	Day   string `json:"day"`
	Month string `json:"month"`
	Year  string `json:"year"`
}

type Events struct {
	Name string `json:"name"`
	Time Date   `json:"time"`
	Desc string `json:"desc"`
}

var EventsData = map[string]Events{}

func init() {
	EventsData["1"] = Events{
		Name: "Google Cloud",
		Time: Date{Day: "01", Month: "03", Year: "2021"},
		Desc: "Description Google event",
	}

	EventsData["2"] = Events{
		Name: "Amazon AWS",
		Time: Date{Day: "12", Month: "03", Year: "2021"},
		Desc: "Description Amazon AWS event",
	}

	EventsData["3"] = Events{
		Name: "Microsoft Azure",
		Time: Date{Day: "11", Month: "03", Year: "2022"},
		Desc: "Description Microsoft Azure event",
	}

	EventsData["4"] = Events{
		Name: "Yahoo event",
		Time: Date{Day: "29", Month: "03", Year: "2021"},
		Desc: "Description Yahoo event event",
	}

}

func main() {
	router()
}
