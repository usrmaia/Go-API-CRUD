package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"google.golang.org/grpc"

	"github.com/usrmaia/Go-API-CRUD/pb"
	"github.com/usrmaia/Go-API-CRUD/src/model"
	"github.com/usrmaia/Go-API-CRUD/src/view"
)

var ClientConn *grpc.ClientConn
var SendMessageClient pb.SendMessageClient

func ClientConnDial(target string) {
	var err error
	ClientConn, err = grpc.Dial(target, grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	SendMessageClient = pb.NewSendMessageClient(ClientConn)
}

func OpenDB() {
	SendMessageClient := pb.NewSendMessageClient(ClientConn)

	req := &pb.RequestDataSourceName{
		DataSourceName: "root:250721@tcp(172.17.0.2:3306)/suzana_motorcycle_parts",
	}

	res, err := SendMessageClient.OpenDB(context.Background(), req)

	if err != nil {
		log.Fatal("Erro ao abrir banco ", err)
	}

	fmt.Println("Database Connection Status:", res.GetStatus())
}

func Home(w http.ResponseWriter, r *http.Request) {
	RequestMessage := &pb.RequestMessage{
		Message: "Test home",
	}

	res, err := SendMessageClient.Home(context.Background(), RequestMessage)

	if err != nil {
		fmt.Println("Error Send Home Server API:", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Println(res.GetStatus())
	w.WriteHeader(http.StatusOK)
}

func ReturnAPart(w http.ResponseWriter, r *http.Request, id int) {
	RequestPartID := &pb.RequestPartID{
		Id: int64(id),
	}
	res, err := SendMessageClient.ReturnAPart(context.Background(), RequestPartID)

	if err != nil {
		fmt.Println("Error Send ReturnAPart Server API:", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	part := model.Part{
		Id:    res.GetId(),
		Name:  res.GetName(),
		Brand: res.GetBrand(),
		Value: res.GetValue(),
	}

	view.ResponsePart(w, part)
}

func ReturnParts(w http.ResponseWriter, r *http.Request) {
	RequestMessage := &pb.RequestMessage{}
	res, err := SendMessageClient.ReturnParts(context.Background(), RequestMessage)

	if err != nil {
		fmt.Println("Error Send ReturnParts Server API:", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var Parts []model.Part
	for _, part := range res.GetParts() {
		Parts = append(Parts, model.Part{
			Id:    part.GetId(),
			Name:  part.GetName(),
			Brand: part.GetBrand(),
			Value: part.GetValue(),
		})
	}

	view.ResponseParts(w, Parts)
}

func AddPart(w http.ResponseWriter, r *http.Request) {
	var data []byte
	var err error
	data, err = ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println("Error ReadAll Request AddPart Server API:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var Part model.Part
	json.Unmarshal(data, &Part)

	RequestAdd := &pb.RequestAdd{
		Name:  Part.Name,
		Brand: Part.Brand,
		Value: Part.Value,
	}

	res, err := SendMessageClient.AddPart(context.Background(), RequestAdd)

	if err != nil {
		fmt.Println("Error Send AddPart Server API:", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	Part.Id = res.GetId()

	view.ResponsePart(w, Part)
}

func DelPart(w http.ResponseWriter, r *http.Request, id int) {
	RequestPartID := &pb.RequestPartID{
		Id: int64(id),
	}

	res, err := SendMessageClient.DelPart(context.Background(), RequestPartID)

	if err != nil {
		fmt.Println("Error Send DelPart Server API:", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	Part := model.Part{
		Id:    res.GetId(),
		Name:  res.GetName(),
		Brand: res.GetBrand(),
		Value: res.GetValue(),
	}

	view.ResponsePart(w, Part)
}

func UpPart(w http.ResponseWriter, r *http.Request) {
	var data []byte
	var err error
	data, err = ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println("Error ReadAll Request UpPart Server API:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var Part model.Part
	json.Unmarshal(data, &Part)

	RequestUp := &pb.RequestUp{
		Id:    int64(Part.Id),
		Name:  Part.Name,
		Brand: Part.Brand,
		Value: Part.Value,
	}

	_, err = SendMessageClient.UpPart(context.Background(), RequestUp)

	if err != nil {
		fmt.Println("Error Send UpPart Server API:", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	view.ResponsePart(w, Part)
}
