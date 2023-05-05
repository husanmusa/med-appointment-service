package test

import (
	"context"
	pb "github.com/husanmusa/med-appointment-service/genproto/appointment_service"
	"github.com/husanmusa/med-appointment-service/storage/postgres"
	"reflect"
	"testing"
)

func Test_doctorRepo_Create(t *testing.T) {
	type args struct {
		ctx context.Context
		req *pb.Doctor
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test1",
			args{
				context.Background(),
				&pb.Doctor{
					Name:       "Doktor Doktorova",
					Age:        44,
					Role:       "Pediatr",
					Polyclinic: "45-sonli poliklinika",
					Gender:     0,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &postgres.DoctorRepo{
				Db: db,
			}
			if err := r.Create(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_doctorRepo_Get(t *testing.T) {

	type args struct {
		ctx context.Context
		req *pb.DoctorId
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.Doctor
		wantErr bool
	}{
		{
			"test1",
			args{
				context.Background(),
				&pb.DoctorId{DoctorId: "9dded5e0-f7a7-4165-a5be-409c1a5c6f81"},
			},
			&pb.Doctor{
				DoctorId:   "9dded5e0-f7a7-4165-a5be-409c1a5c6f81",
				Name:       "Doktor Doktorova",
				Age:        44,
				Role:       "Pediatr",
				Polyclinic: "45-sonli poliklinika",
				Gender:     0,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &postgres.DoctorRepo{
				Db: db,
			}
			got, err := r.Get(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_doctorRepo_GetList(t *testing.T) {
	type args struct {
		ctx context.Context
		req *pb.GetListDoctorsRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.GetListDoctorsResponse
		wantErr bool
	}{
		{
			"test1",
			args{
				context.Background(),
				&pb.GetListDoctorsRequest{},
			},
			&pb.GetListDoctorsResponse{
				Doctors: []*pb.Doctor{
					{
						DoctorId:   "9dded5e0-f7a7-4165-a5be-409c1a5c6f81",
						Name:       "Doktor Doktorova",
						Age:        44,
						Role:       "Pediatr",
						Polyclinic: "45-sonli poliklinika",
						Gender:     0,
					},
				},
				Count: 1,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &postgres.DoctorRepo{
				Db: db,
			}
			got, err := r.GetList(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_doctorRepo_Update(t *testing.T) {
	type args struct {
		ctx context.Context
		req *pb.Doctor
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
				&pb.Doctor{
					DoctorId:   "9dded5e0-f7a7-4165-a5be-409c1a5c6f81",
					Name:       "Doktor Doktorova",
					Age:        44,
					Role:       "Pediatr",
					Polyclinic: "45-sonli poliklinika",
					Gender:     0,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &postgres.DoctorRepo{
				Db: db,
			}
			if err := r.Update(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_doctorRepo_Delete(t *testing.T) {
	type args struct {
		ctx context.Context
		req *pb.DoctorId
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
				&pb.DoctorId{DoctorId: "9dded5e0-f7a7-4165-a5be-409c1a5c6f82"},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &postgres.DoctorRepo{
				Db: db,
			}
			if err := r.Delete(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
