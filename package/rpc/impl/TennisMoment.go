package impl

import (
	"context"
	"encoding/json"

	"github.com/zhangjiacheng-iHealth/IHCommunity/package/dao/impl"
	"github.com/zhangjiacheng-iHealth/IHCommunity/package/model"
	"github.com/zhangjiacheng-iHealth/IHCommunity/proto"
)

type IHCommunityRPCImpl struct {
	dao *impl.IHCommunityDaoImpl
}

func NewIHCommunityControllerImpl() *IHCommunityRPCImpl {
	return &IHCommunityRPCImpl{dao: impl.NewIHCommunityDaoImpl()}
}
func (impl *IHCommunityRPCImpl) AddPlayer(ctx context.Context, request *proto.PlayerInfoRequest) (*proto.PlayerInfoResponse, error) {
	Id := request.GetId()
	LoginName := request.GetLoginName()
	Name := request.GetName()
	Icon := request.GetIcon()
	Sex := request.GetSex()
	Age := request.GetAge()
	YearsPlayed := request.GetYearsPlayed()
	Height := request.GetHeight()
	Width := request.GetWidth()
	Grip := request.GetGrip()
	Backhand := request.GetBackhand()
	Points := request.GetPoints()
	IsAdult := request.GetIsAdult()
	CareerStats := request.GetCareerStats()

	impl.dao.AddPlayer(ctx, model.Player{
		Id:            Id,
		LoginName:     LoginName,
		Name:          Name,
		Icon:          Icon,
		Sex:           Sex,
		Age:           Age,
		YearsPlayed:   YearsPlayed,
		Height:        Height,
		Width:         Width,
		Grip:          Grip,
		Backhand:      Backhand,
		Points:        Points,
		IsAdult:       IsAdult,
		CareerStatsId: CareerStats.Id,
	})
	return &proto.PlayerInfoResponse{Code: 0, Msg: "", Count: 1, Data: "true"}, nil
}

// func (impl *PlayerInfoRPCImpl) FindPlayerByKey(ctx context.Context, request *proto.PlayerInfoRequest) (*proto.PlayerInfoResponse, error) {
// 	key := request.GetInfoKey()
// 	Player := impl.dao.GetPlayerByUid(ctx, key)
// 	info, _ := json.Marshal(Player)
// 	return &proto.PlayerInfoResponse{Code: 0, Msg: "", Count: 1, Data: string(info)}, nil
// }

func (impl *IHCommunityRPCImpl) UpdatePlayer(ctx context.Context, request *proto.PlayerInfoRequest) (*proto.PlayerInfoResponse, error) {
	Id := request.GetId()
	LoginName := request.GetLoginName()
	Name := request.GetName()
	Icon := request.GetIcon()
	Sex := request.GetSex()
	Age := request.GetAge()
	YearsPlayed := request.GetYearsPlayed()
	Height := request.GetHeight()
	Width := request.GetWidth()
	Grip := request.GetGrip()
	Backhand := request.GetBackhand()
	Points := request.GetPoints()
	IsAdult := request.GetIsAdult()
	CareerStats := request.GetCareerStats()
	impl.dao.UpdatePlayer(ctx, model.PlayerResponse{
		Id:          Id,
		LoginName:   LoginName,
		Name:        Name,
		Icon:        Icon,
		Sex:         Sex,
		Age:         Age,
		YearsPlayed: YearsPlayed,
		Height:      Height,
		Width:       Width,
		Grip:        Grip,
		Backhand:    Backhand,
		Points:      Points,
		IsAdult:     IsAdult,
		CareerStats: CareerStats,
	})
	return &proto.PlayerInfoResponse{Code: 0, Msg: "", Count: 1, Data: "true"}, nil
}

// func (impl *PlayerInfoRPCImpl) DeleteById(ctx context.Context, request *proto.PlayerInfoRequest) (*proto.PlayerInfoResponse, error) {
// 	id := request.GetId()
// 	impl.dao.DeletePlayerById(ctx, id)
// 	return &proto.PlayerInfoResponse{Code: 0, Msg: "", Count: 1, Data: "true"}, nil
// }

func (impl *IHCommunityRPCImpl) SearchPlayer(ctx context.Context, request *proto.PlayerInfoRequest) (*proto.PlayerInfoResponse, error) {
	id := request.GetId()
	Players, _ := impl.dao.GetPlayerInfo(ctx, id)
	infos, _ := json.Marshal(Players)
	return &proto.PlayerInfoResponse{Code: 0, Msg: "", Count: 1, Data: string(infos)}, nil
}

func (impl *IHCommunityRPCImpl) GetPlayerInfo(ctx context.Context, req *proto.PlayerInfoRequest) (resp *proto.PlayerInfoResponse, err error) {
	id := req.GetId()
	Players, _ := impl.dao.GetPlayerInfo(ctx, id)
	infos, _ := json.Marshal(Players)
	return &proto.PlayerInfoResponse{Code: 0, Msg: "", Count: 1, Data: string(infos)}, nil
}

func (impl *IHCommunityRPCImpl) AddFriend(ctx context.Context, request *proto.PlayerInfoRequest) (*proto.PlayerInfoResponse, error) {
	id := request.GetId()
	Players, _ := impl.dao.GetPlayerInfo(ctx, id)
	infos, _ := json.Marshal(Players)
	return &proto.PlayerInfoResponse{Code: 0, Msg: "", Count: 1, Data: string(infos)}, nil
}
func (impl *IHCommunityRPCImpl) DeleteFriend(ctx context.Context, request *proto.PlayerInfoRequest) (*proto.PlayerInfoResponse, error) {
	id := request.GetId()
	Players, _ := impl.dao.GetPlayerInfo(ctx, id)
	infos, _ := json.Marshal(Players)
	return &proto.PlayerInfoResponse{Code: 0, Msg: "", Count: 1, Data: string(infos)}, nil
}
