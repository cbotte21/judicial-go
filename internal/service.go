package internal

import (
	"context"
	"errors"
	hive "github.com/cbotte21/hive-go/pb"
	"github.com/cbotte21/judicial-go/internal/schema"
	"github.com/cbotte21/judicial-go/pb"
	"github.com/cbotte21/microservice-common/pkg/datastore"
	"github.com/cbotte21/microservice-common/pkg/jwtParser"
	"time"
)

type Judicial struct {
	JwtSecret        *jwtParser.JwtSecret
	HiveClient       *hive.HiveServiceClient
	MongoBanClient   *datastore.MongoClient[schema.Ban]
	MongoUnbanClient *datastore.MongoClient[schema.Unban]
	pb.UnimplementedJudicialServiceServer
}

func NewJudicial(hiveClient *hive.HiveServiceClient, mongoBanClient *datastore.MongoClient[schema.Ban], mongoUnbanClient *datastore.MongoClient[schema.Unban]) Judicial {
	return Judicial{HiveClient: hiveClient, MongoBanClient: mongoBanClient, MongoUnbanClient: mongoUnbanClient}
}

func canBan(role int) error {
	if role > 0 {
		return nil
	}
	return errors.New("insignificant permissions")
}

func (judicial *Judicial) Ban(ctx context.Context, banRequest *pb.BanRequest) (*pb.BanResponse, error) {
	admin, err := judicial.JwtSecret.Redeem(banRequest.GetGod())

	if err == nil {
		err = canBan(admin.Role)
		if err == nil {
			err := judicial.MongoBanClient.Create(schema.Ban{
				Player:    banRequest.XId,
				God:       admin.Id,
				Reason:    banRequest.GetReason(),
				Expiry:    banRequest.GetExpiry().String(),
				Timestamp: time.Now().String(),
			})
			if err == nil { //Success
				_, _ = (*judicial.HiveClient).ForceDisconnect(context.Background(), &hive.DisconnectRequest{Id: banRequest.GetXId()})
				return &pb.BanResponse{Status: true}, nil
			}
		}
	}
	return &pb.BanResponse{Status: false}, err
}

func (judicial *Judicial) Unban(ctx context.Context, unbanRequest *pb.UnbanRequest) (*pb.UnbanResponse, error) {
	admin, err := judicial.JwtSecret.Redeem(unbanRequest.GetGod())

	if err == nil {
		err = canBan(admin.Role)
		if err == nil {
			err := judicial.MongoBanClient.Delete(schema.Ban{Player: unbanRequest.GetXId()})
			if err == nil {
				err := judicial.MongoUnbanClient.Create(schema.Unban{
					Player:    unbanRequest.GetXId(),
					God:       admin.Id,
					Timestamp: time.Now().String(),
				})
				if err == nil { //Success
					return &pb.UnbanResponse{Status: true}, nil
				}
			}
		}
	}
	return &pb.UnbanResponse{Status: false}, err
}

func (judicial *Judicial) Integrity(ctx context.Context, integrityRequest *pb.IntegrityRequest) (*pb.IntegrityResponse, error) {
	ban, err := judicial.MongoBanClient.Find(schema.Ban{Player: integrityRequest.GetXId()})
	if err != nil { // Player is not banned
		return &pb.IntegrityResponse{Status: true}, nil
	}
	return &pb.IntegrityResponse{Status: false}, errors.New("player is banned until " + ban.Expiry)
}
