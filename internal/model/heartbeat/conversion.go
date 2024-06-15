package heartbeatmodel

import (
	"github.com/jackc/pgx/v5/pgtype"
	storehb "github.com/soumya-codes/postgres-static-shard/internal/store/heartbeat"
)

func ConvertToStoreHeartbeatModel(hb *Heartbeat) (*storehb.Heartbeat, error) {
	var pgMID pgtype.Text
	err := pgMID.Scan(hb.MachineID)
	if err != nil || !pgMID.Valid {
		return nil, err
	}

	var pgLHB pgtype.Int8
	err = pgLHB.Scan(hb.LastHeartbeat)
	if err != nil || !pgLHB.Valid {
		return nil, err
	}

	return &storehb.Heartbeat{
		MachineID:     pgMID,
		LastHeartbeat: pgLHB,
	}, nil
}

func ConvertToHeartbeatModel(hb *storehb.Heartbeat) *Heartbeat {
	return &Heartbeat{
		MachineID:     hb.MachineID.String,
		LastHeartbeat: hb.LastHeartbeat.Int64,
	}
}
