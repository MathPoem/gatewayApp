package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"net/http"
)

const (
	baseQueryUrl = "https://www.rusprofile.ru"
	searchQuery  = "ajax.php?query=%s&action=search"
)

type GatewayService struct {
	UnimplementedGatewayServiceServer
}

type Organization struct {
	Inn  string `json:"inn"`
	Kpp  string `json:"kpp"`
	Name string `json:"name"`
	Head string `json:"head"`
}

func (s *GatewayService) FindORG(ctx context.Context, inn *RequestINN) (*ResponseORG, error) {
	_, err := fetchGeneralInfo(ctx, inn)

	if err != nil {
		return nil, err
	}

	return &ResponseORG{
			Inn:  inn.Inn,
			Kpp:  "kpp",
			Name: "name",
			Head: "head",
		},
		nil
}

func fetchGeneralInfo(ctx context.Context, inn *RequestINN) (*Organization, error) {
	searchURL := "https://www.rusprofile.ru/search?query=%s&search_inactive=0&type=ul"

	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf(searchURL, inn.Inn), nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)

	var data struct {
		Items []Organization `json:"ul"`
	}

	if err := json.Unmarshal(b, &data); err != nil {
		return nil, status.Errorf(codes.Internal, "couldn't unmarshal to json: %v", err)
	}
	fmt.Println(data)

	//if resp.Request.Response != nil {
	//	urlById, _ := resp.Request.Response.Location()
	//	reqById, _ := http.NewRequestWithContext(ctx, "GET", urlById.String(), nil)
	//	resp, err := http.DefaultClient.Do(reqById)
	//	if err != nil {
	//		return nil, err
	//	}
	//	fmt.Println(resp.)
	//}

	return nil, nil
}
