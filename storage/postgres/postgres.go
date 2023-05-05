package postgres

import (
	"context"
	"fmt"
	"github.com/husanmusa/med-appointment-service/config"
	"github.com/husanmusa/med-appointment-service/storage"
	"github.com/jackc/pgx/v4"
	"log"
	"time"

	"github.com/saidamir98/udevs_pkg/logger"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db           *pgxpool.Pool
	log          logger.LoggerI
	appointmentI storage.AppointmentServiceI
	doctorI      storage.DoctorI
	patientI     storage.PatientI
}

func (s *Store) CloseDB() {
	s.db.Close()
}

type PGXStdLogger struct{}

func (l *PGXStdLogger) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	args := make([]interface{}, 0, len(data)+2) // making space for arguments + level + msg
	var query bool
	args = append(args, level, msg, "WARNING!!! SLOW_QUERY")
	for k, v := range data {
		args = append(args, fmt.Sprintf("%s=%v", k, v))

		if k == "time" {
			t := v.(time.Duration)

			if t > time.Millisecond*500 {
				query = true
			} else {
				query = false
			}
		}
	}

	if query {
		log.Println(args...)
	}
}

func NewPostgres(ctx context.Context, cfg config.Config, log logger.LoggerI) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnections
	config.ConnConfig.LogLevel = pgx.LogLevelInfo
	config.ConnConfig.Logger = &PGXStdLogger{}

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return &Store{
		db:  pool,
		log: log,
	}, err
}

func (s *Store) Appointment() storage.AppointmentServiceI {
	if s.appointmentI == nil {
		s.appointmentI = NewAppointmentRepo(s.db)
	}

	return s.appointmentI
}

func (s *Store) Doctor() storage.DoctorI {
	if s.doctorI == nil {
		s.doctorI = NewDoctorRepo(s.db)
	}

	return s.doctorI
}

func (s *Store) Patient() storage.PatientI {
	if s.patientI == nil {
		s.patientI = NewPatientRepo(s.db)
	}

	return s.patientI
}
