package bot

import (
	"encoding/json"
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"net/http"
	"strconv"
	"sync"
)

var (
	KeyAll = tb.InlineButton{Unique:"i1", Text:"Все случаи",}
	KeyByCountry = tb.InlineButton{Unique:"i2",Text:"Выбрать страну"}
	KeyKz = tb.InlineButton{Unique:"i3",Text:"Казахстан"}
	inlineKeys = [][]tb.InlineButton{
		[]tb.InlineButton{KeyAll,KeyKz},
		[]tb.InlineButton{KeyByCountry},
	}
	country = ""

)

func NewEndpointsFactory() *endpointsFactory {
	return &endpointsFactory{}
}

type endpointsFactory struct {}

func (ef *endpointsFactory) Start(b *tb.Bot) func (m *tb.Message){
	return func(m *tb.Message) {
		b.Send(m.Sender,fmt.Sprintf("Привет %s. Добро пожаловать в бот Covid-19",m.Sender.FirstName),&tb.ReplyMarkup{InlineKeyboard:inlineKeys})

	}
}


func (ef *endpointsFactory) AllCases(b *tb.Bot) func(c *tb.Callback){
	return func(c *tb.Callback) {
		re,err:= http.Get("https://coronavirus-19-api.herokuapp.com/all")
		if err!=nil {
			fmt.Print(err.Error())
		}
		cas := AllCases{}
		decoder := json.NewDecoder(re.Body)
		error := decoder.Decode(&cas)
		if error!=nil {
			fmt.Print(error.Error())
		}else{
			b.Edit(c.Message,fmt.Sprintf("Все случаи: %v \nСмертей: %v \nВыздоровел: %v",strconv.Itoa(cas.Cases),strconv.Itoa(cas.Deaths),strconv.Itoa(cas.Recovered)))
			b.EditReplyMarkup(c.Message,&tb.ReplyMarkup{InlineKeyboard:inlineKeys})g
		}
	}
}

func (ef *endpointsFactory) CheckKz(b *tb.Bot) func(c *tb.Callback){
	return func(c *tb.Callback) {
		re,err:= http.Get("https://coronavirus-19-api.herokuapp.com/countries/Kazakhstan")
		if err!=nil {
			fmt.Print(err.Error())
		}
		cas := Countries{}

		decoder := json.NewDecoder(re.Body)
		error := decoder.Decode(&cas)

		if error!=nil {
			fmt.Print(error.Error())
		}else{
			b.Edit(c.Message,fmt.Sprintf("Все случаи: %v\nСегодняшние происшествия: %v\nСмертей: %v\nСегодняшние смерти: %v\nВыздоровел: %v\nКритический: %v",cas.Cases,cas.TodayCases,cas.Deaths,cas.TodayDeaths,cas.Recovered,cas.Critical))
			b.EditReplyMarkup(c.Message,&tb.ReplyMarkup{InlineKeyboard:inlineKeys})
		}
	}
}

func (ef *endpointsFactory) FindByCountry(b *tb.Bot) func(c *tb.Callback){
	return func(c *tb.Callback) {
		var wg sync.WaitGroup
		wg.Add(1)
		GetCountry(ef,b,c,wg)


	}
}
func GetCountry(ef *endpointsFactory,b *tb.Bot,c *tb.Callback, wg sync.WaitGroup) {
	defer wg.Done()
	b.Send(c.Sender,"Напишите название страны (На английском и без ошибок)")
	ch := make(chan string)
	go func(){
		b.Handle(tb.OnText, func(m *tb.Message) {
			ch <- m.Text
		})
	}()
	res := <-ch
	re,err:= http.Get("https://coronavirus-19-api.herokuapp.com/countries/"+res)
	if err!=nil {
		fmt.Print(err.Error()+" 1111")
		wg.Add(1)
		GetCountry(ef,b,c,wg)
	}else{
		cas := Countries{}

		decoder := json.NewDecoder(re.Body)
		error := decoder.Decode(&cas)

		if error!=nil {
			fmt.Print(error.Error())
			wg.Add(1)
			GetCountry(ef,b,c,wg)
		}else{
			b.Send(c.Sender,fmt.Sprintf("Название страны: %v\nВсе случаи: %v\nСегодняшние происшествия: %v\nСмертей: %v\nСегодняшние смерти: %v\nВыздоровел: %v\nКритический: %v",cas.Country,cas.Cases,cas.TodayCases,cas.Deaths,cas.TodayDeaths,cas.Recovered,cas.Critical),&tb.ReplyMarkup{InlineKeyboard:inlineKeys})
		}
	}

}
