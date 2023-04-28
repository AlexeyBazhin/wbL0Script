package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/google/uuid"
	stan "github.com/nats-io/stan.go"
)

type (
	Data struct {
		Type  string `json:"type"`
		Model `json:"model"`
	}
	Model struct {
		Id                uuid.UUID `json:"order_uid"`
		TrackNumber       string    `json:"track_number"`
		Entry             string    `json:"entry"`
		Delivery          `json:"delivery"`
		Payment           `json:"payment"`
		Items             []Item    `json:"items"`
		Locale            string    `json:"locale"`
		InternalSignature string    `json:"internal_signature"`
		CustomerId        string    `json:"customer_id"`
		DeliveryService   string    `json:"delivery_service"`
		Shardkey          string    `json:"shardkey"`
		SmId              int       `json:"sm_id"`
		DateCreated       time.Time `json:"date_created"`
		OofShard          string    `json:"oof_shard"`
	}
	Delivery struct {
		Name    string `json:"name"`
		Phone   string `json:"phone"`
		Zip     string `json:"zip"`
		City    string `json:"city"`
		Address string `json:"address"`
		Region  string `json:"region"`
		Email   string `json:"email"`
	}
	Payment struct {
		Transaction  uuid.UUID `json:"transaction"`
		RequestId    string    `json:"request_id"`
		Currency     string    `json:"currency"`
		Provider     string    `json:"provider"`
		Amount       int       `json:"amount"`
		PaymentDt    int       `json:"payment_dt"`
		Bank         string    `json:"bank"`
		DeliveryCost int       `json:"delivery_cost"`
		GoodsTotal   int       `json:"goods_total"`
		CustomFee    int       `json:"custom_fee"`
	}
	Item struct {
		ChrtId      int    `json:"chrt_id"`
		TrackNumber string `json:"track_number"`
		Price       int    `json:"price"`
		RId         string `json:"rid"`
		Name        string `json:"name"`
		Sale        int    `json:"sale"`
		Size        string `json:"size"`
		TotalPrice  int    `json:"total_price"`
		NmId        int    `json:"nm_id"`
		Brand       string `json:"brand"`
		Status      int    `json:"status"`
	}
)

// var jsonData = []byte(`
// {
// 	"order_uid": "20354d7a-e4fe-47af-8ff6-187bca92f3f9",
// 	"track_number": "WBILMTESTTRACK",
// 	"entry": "WBIL",
// 	"delivery": {
// 	  "name": "Test Testov",
// 	  "phone": "+9720000000",
// 	  "zip": "2639809",
// 	  "city": "Kiryat Mozkin",
// 	  "address": "Ploshad Mira 15",
// 	  "region": "Kraiot",
// 	  "email": "test@gmail.com"
// 	},
// 	"payment": {
// 	  "transaction": "20354d7a-e4fe-47af-8ff6-187bca92f3f9",
// 	  "request_id": "",
// 	  "currency": "USD",
// 	  "provider": "wbpay",
// 	  "amount": 1817,
// 	  "payment_dt": 1637907727,
// 	  "bank": "alpha",
// 	  "delivery_cost": 1500,
// 	  "goods_total": 317,
// 	  "custom_fee": 0
// 	},
// 	"items": [
// 	  {
// 		"chrt_id": 9934930,
// 		"track_number": "WBILMTESTTRACK",
// 		"price": 453,
// 		"rid": "ab4219087a764ae0btest",
// 		"name": "Mascaras",
// 		"sale": 30,
// 		"size": "0",
// 		"total_price": 317,
// 		"nm_id": 2389212,
// 		"brand": "Vivienne Sabo",
// 		"status": 202
// 	  }
// 	],
// 	"locale": "en",
// 	"internal_signature": "",
// 	"customer_id": "test",
// 	"delivery_service": "meest",
// 	"shardkey": "9",
// 	"sm_id": 99,
// 	"date_created": "2021-11-26T06:22:19Z",
// 	"oof_shard": "1"
// }
// `)

func main() {
	publish()
	
	// var order Model
	// if err := json.Unmarshal(jsonData, &order); err != nil {
	// 	fmt.Println(fmt.Errorf("failed to unmarshal json data: %s", err.Error()))
	// }
	// fmt.Println(order)
	// model := makeNewJSON()
	// data := Data{
	// 	Type:  "order",
	// 	Model: model,
	// }
	// if dataByte, err := json.Marshal(data); err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println(string(dataByte))
	// 	// newModel := Model{}
	// 	// if err := json.Unmarshal(dataByte, &newModel); err != nil {
	// 	// 	panic(err)
	// 	// 	//fmt.Println(fmt.Errorf("failed to unmarshal json data: %s", err.Error()))
	// 	// }
	// 	// fmt.Println(newModel)
	// }
}
func publish() {
	sc, err := stan.Connect("amethyst-cluster", "clientID")
	if err != nil {
		panic(err)
	}

	defer sc.Close()

	for i := 0; i < 5; i++ {
		model := makeNewJSON()
		data := Data{
			Type:  "order",
			Model: model,
		}
		if dataByte, err := json.Marshal(data); err != nil {
			panic(err)
		} else {
			fmt.Println(data.Id)
			if err := sc.Publish("models", dataByte); err != nil {
				panic(err)
			}
		}
	}
}
func makeNewJSON() Model {
	rand.Seed(time.Now().UnixNano())
	orderUid := uuid.New()
	trackNumber := strconv.Itoa(rand.Intn(100)) + "WBILM" + strconv.Itoa(rand.Intn(100))

	itemsCount := rand.Intn(5) + 1
	items := make([]Item, itemsCount)
	for i := 0; i < itemsCount; i++ {
		items[i] = Item{
			ChrtId:      rand.Intn(100000),
			TrackNumber: trackNumber,
			Price:       rand.Intn(5000),
			RId:         "ab4219087a764ae0btest",
			Name:        strconv.Itoa(rand.Intn(100)) + "Mascars",
			Sale:        rand.Intn(100),
			Size:        "0",
			TotalPrice:  317,
			NmId:        2389212,
			Brand:       "Vivienne",
			Status:      202,
		}
	}
	return Model{
		Id:                orderUid,
		TrackNumber:       trackNumber,
		Entry:             "WBIL",
		Locale:            "en",
		InternalSignature: "",
		CustomerId:        "test",
		DeliveryService:   "meest",
		Shardkey:          "9",
		SmId:              99,
		DateCreated:       time.Now(),
		OofShard:          "1",
		Delivery: Delivery{
			Name:    "Test Testov" + strconv.Itoa(rand.Intn(100)),
			Phone:   "+9720000000",
			Zip:     "2639809",
			City:    "Kiryat Mozkin",
			Address: "Ploshad",
			Region:  "Kraiot",
			Email:   "test@gmail.com",
		},
		Payment: Payment{
			Transaction:  orderUid,
			RequestId:    "",
			Currency:     "USD",
			Provider:     "wbpay",
			Amount:       rand.Intn(2000),
			PaymentDt:    1637907727,
			Bank:         "alpha",
			DeliveryCost: rand.Intn(2000),
			GoodsTotal:   rand.Intn(500),
			CustomFee:    0,
		},
		Items: items,
	}
}
