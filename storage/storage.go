package storage

import (
	"context"
	pb "github.com/husanmusa/med-appointment-service/genproto/appointment_service"
)

type StorageI interface {
	Appointment() AppointmentServiceI
	CloseDB()
}

type AppointmentServiceI interface {
	Create(ctx context.Context, req *pb.Appointment) (*pb.Appointment, error)
	Get(ctx context.Context, req *pb.AppointmentId) (*pb.Appointment, error)
	Cancel(ctx context.Context, req *pb.AppointmentId) error
}
