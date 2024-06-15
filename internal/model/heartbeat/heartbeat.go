package heartbeatmodel

type Heartbeat struct {
	MachineID     string `json:"machine_id"`
	LastHeartbeat int64  `json:"last_heartbeat"`
}
