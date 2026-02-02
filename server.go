package main

import (
	pb "github.com/christian-schueler/go-grpc/fmdata_proto"
)

type server struct {
	pb.UnimplementedFmDataServiceServer
}

func (s *server) GetFmItems(context Context, request *pb.FmDataRequest) (*pb.FmData, error) {
	return &pb.FmData{
		Items: []*pb.Item{
			{
				Id:          "4d9105a9-aa94-490f-a45d-c5dcbe73a39e",
				Type:        "country",
				Name:        "Deutschland",
				Nummer:      "1",
				Description: nil,
				Parent:      nil,
				ParentType:  nil,
				Status:      1,
			},
		},
		[]*pb.Item{
			{
				Id:          "cbb8f0ea-91e8-423b-a39c-c1fd7bde12be",
				Type:        "city",
				Name:        "Hildburghausen",
				Nummer:      "1",
				Description: nil,
				Parent:      "4d9105a9-aa94-490f-a45d-c5dcbe73a39e",
				ParentType:  "country",
				Status:      1,
			},
		},
		[]*pb.Item{
			{
				Id:          "e95e31f1-1ca0-4330-b1a2-66989d18010d",
				Type:        "city",
				Name:        "Suhl",
				Nummer:      "1",
				Description: nil,
				Parent:      "4d9105a9-aa94-490f-a45d-c5dcbe73a39e",
				ParentType:  "country",
				Status:      1,
			},
		},
		[]*pb.Item{
			{
				Id:          "8b8b88b0-d519-44b6-9938-4bfe60dddab9",
				Type:        "city",
				Name:        "Coburg",
				Nummer:      "1",
				Description: nil,
				Parent:      "4d9105a9-aa94-490f-a45d-c5dcbe73a39e",
				ParentType:  "country",
				Status:      1,
			},
		},
	}, nil
}

func (s *server) GetFmItem(context Context, request *pb.ItemRequest) (*pb.Item, error) {
}
