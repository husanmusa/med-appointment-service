package test

import (
	"context"
	pb "github.com/husanmusa/med-appointment-service/genproto/appointment_service"
	"github.com/husanmusa/med-appointment-service/storage/postgres"
	"reflect"
	"testing"
)

func Test_appointmentRepo_Create(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *pb.Appointment
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.Appointment
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				ctx: context.Background(),
				in: &pb.Appointment{
					DoctorId:    "405777ce-9e46-4aa0-bf61-3af62df0721e",
					PatientId:   "36b1efef-82ed-4d20-b402-0b20577f2815",
					AppointDate: "2023-05-04",
					AppointTime: "15:00",
				}},
			want: &pb.Appointment{
				DoctorId:    "405777ce-9e46-4aa0-bf61-3af62df0721e",
				PatientId:   "36b1efef-82ed-4d20-b402-0b20577f2815",
				AppointDate: "2023-05-04",
				AppointTime: "15:00",
				AppointId:   7,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := postgres.AppointmentRepo{
				Db: db,
			}
			got, err := a.Create(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_appointmentRepo_Get(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *pb.AppointmentId
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.Appointment
		wantErr bool
	}{
		{
			"test1",
			args{
				context.Background(),
				&pb.AppointmentId{Id: "f8c5e607-2d80-4119-b50c-9086992f7dfb"},
			},
			&pb.Appointment{
				DoctorId:    "405777ce-9e46-4aa0-bf61-3af62df0721e",
				PatientId:   "36b1efef-82ed-4d20-b402-0b20577f2815",
				AppointDate: "2023-05-04",
				AppointTime: "15:00",
				AppointId:   7,
				Status:      1,
				CreatedAt:   "2023-05-04 20:07:05",
				UpdatedAt:   "2023-05-04 20:07:05",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := postgres.AppointmentRepo{
				Db: db,
			}
			got, err := a.Get(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, \n\t\t\t wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, \n\t\t\t\t\t\t want %v", got, tt.want)
			}
		})
	}
}

func Test_appointmentRepo_Cancel(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *pb.AppointmentId
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"test1",
			args{
				context.Background(),
				&pb.AppointmentId{Id: "bf8ff53c-09b1-4b90-8170-9068c19770f5"},
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := postgres.AppointmentRepo{
				Db: db,
			}
			if err := a.Cancel(tt.args.ctx, tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("Cancel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
