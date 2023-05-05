package test

import (
	"context"
	pb "github.com/husanmusa/med-appointment-service/genproto/appointment_service"
	"github.com/husanmusa/med-appointment-service/storage/postgres"
	"reflect"
	"testing"
)

func Test_patientRepo_Create(t *testing.T) {
	type args struct {
		ctx context.Context
		req *pb.Patient
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test1",
			args{
				context.Background(),
				&pb.Patient{
					Name:   "Bemor Bemorov",
					Age:    25,
					Gender: 1,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &postgres.PatientRepo{
				Db: db,
			}
			if err := r.Create(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_patientRepo_Get(t *testing.T) {

	type args struct {
		ctx context.Context
		req *pb.PatientId
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.Patient
		wantErr bool
	}{
		{
			"test1",
			args{
				context.Background(),
				&pb.PatientId{PatientId: "537cf771-4c06-4309-bc7c-62d995a00e42"},
			},
			&pb.Patient{
				PatientId: "537cf771-4c06-4309-bc7c-62d995a00e42",
				Name:      "Bemor Bemorov",
				Age:       25,
				Gender:    1,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &postgres.PatientRepo{
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

func Test_patientRepo_GetList(t *testing.T) {
	type args struct {
		ctx context.Context
		req *pb.GetListPatientsRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.GetListPatientsResponse
		wantErr bool
	}{
		{
			"test1",
			args{
				context.Background(),
				&pb.GetListPatientsRequest{},
			},
			&pb.GetListPatientsResponse{
				Patients: []*pb.Patient{
					{
						PatientId: "43b77f9c-0f9e-49cc-ab00-f8c7ed5a6d62",
						Name:      "Bemor Bemorov",
						Age:       25,
						Gender:    1,
					},
				},
				Count: 1,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &postgres.PatientRepo{
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

func Test_patientRepo_Update(t *testing.T) {
	type args struct {
		ctx context.Context
		req *pb.Patient
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
				&pb.Patient{
					PatientId: "43b77f9c-0f9e-49cc-ab00-f8c7ed5a6d62",
					Name:      "Bemor Bemorov",
					Age:       25,
					Gender:    1,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &postgres.PatientRepo{
				Db: db,
			}
			if err := r.Update(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_patientRepo_Delete(t *testing.T) {
	type args struct {
		ctx context.Context
		req *pb.PatientId
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
				&pb.PatientId{PatientId: "43b77f9c-0f9e-49cc-ab00-f8c7ed5a6d63"},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &postgres.PatientRepo{
				Db: db,
			}
			if err := r.Delete(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
