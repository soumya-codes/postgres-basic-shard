package heartbeat

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soumya-codes/postgres-static-shard/internal/db/postgres"
	heartbeatmodel "github.com/soumya-codes/postgres-static-shard/internal/model/heartbeat"
	"github.com/soumya-codes/postgres-static-shard/internal/response"
	db "github.com/soumya-codes/postgres-static-shard/internal/store/heartbeat"
)

func GetLastHeartbeat(c *gin.Context) {
	machineID := c.Param("machine_id")
	shardConn, err := postgres.GetShardConnection(machineID)
	if err != nil {
		fmt.Printf("error getting shard connection: %v", err)
		respErr := response.MapErrorToErrorResponse(response.ErrCodeInvalidInput, response.ErrorDetail{
			Field:   "machine_id",
			Message: fmt.Sprintf("invalid machine_id: %s", machineID),
		})
		response.RespondWithError(c, respErr.RespCode, respErr.ResponseError)
		return
	}

	hbStoreModel, err := heartbeatmodel.ConvertToStoreHeartbeatModel(&heartbeatmodel.Heartbeat{
		MachineID: machineID,
	})
	if err != nil {
		fmt.Printf("error converting to store model: %v", err)
		respErr := response.MapErrorToErrorResponse(response.ErrCodeInvalidInput, response.ErrorDetail{
			Field:   "machine_id",
			Message: fmt.Sprintf("invalid machine_id: %v", machineID),
		})
		response.RespondWithError(c, respErr.RespCode, respErr.ResponseError)
		return
	}

	q := db.New(shardConn.Conn)
	lastHeartbeat, err := q.GetLastHeartbeat(c, hbStoreModel.MachineID)
	if err != nil {
		fmt.Printf("error getting last heartbeat from db: %v", err)
		respErr := response.MapErrorToErrorResponse(response.ErrCodeInvalidInput, response.ErrorDetail{
			Field:   "machine_id",
			Message: fmt.Sprintf("invalid machine_id: %v", machineID),
		})
		response.RespondWithError(c, respErr.RespCode, respErr.ResponseError)
		return
	}

	response.RespondWithJSON(c, http.StatusOK, map[string]int64{"last_heartbeat": lastHeartbeat.Int64})
}

func UpdateLastHeartbeat(c *gin.Context) {
	machineID := c.Param("machine_id")

	var requestBody struct {
		LastHeartbeat int64 `json:"last_heartbeat"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		fmt.Println("error binding request body: ", err)
		respErr := response.MapErrorToErrorResponse(response.ErrCodeInvalidInput, response.ErrorDetail{
			Field:   "request_body",
			Message: fmt.Sprintf("invalid request body: %#v", requestBody),
		})
		response.RespondWithError(c, respErr.RespCode, respErr.ResponseError)
		return
	}

	if requestBody.LastHeartbeat <= 0 {
		respErr := response.MapErrorToErrorResponse(response.ErrCodeInvalidInput, response.ErrorDetail{
			Field:   "heartbeat",
			Message: "last_heartbeat should be greater than 0",
		})
		response.RespondWithError(c, respErr.RespCode, respErr.ResponseError)
		return
	}

	shardConn, err := postgres.GetShardConnection(machineID)
	if err != nil {
		fmt.Printf("error getting shard connection: %v", err)
		respErr := response.MapErrorToErrorResponse(response.ErrCodeInternalServerError, response.ErrorDetail{
			Field:   "",
			Message: "internal server error",
		})
		response.RespondWithError(c, respErr.RespCode, respErr.ResponseError)
		return
	}

	hbStoreModel, err := heartbeatmodel.ConvertToStoreHeartbeatModel(&heartbeatmodel.Heartbeat{
		MachineID:     machineID,
		LastHeartbeat: requestBody.LastHeartbeat,
	})
	if err != nil {
		fmt.Printf("error converting to store model: %v", err)
		respErr := response.MapErrorToErrorResponse(response.ErrCodeInvalidInput, response.ErrorDetail{
			Field:   "request_body",
			Message: fmt.Sprintf("invalid request body: %#v", requestBody),
		})
		response.RespondWithError(c, respErr.RespCode, respErr.ResponseError)
		return
	}

	q := db.New(shardConn.Conn)
	_, err = q.UpdateLastHeartbeat(c, db.UpdateLastHeartbeatParams{
		MachineID:     hbStoreModel.MachineID,
		LastHeartbeat: hbStoreModel.LastHeartbeat,
	})
	if err != nil {
		fmt.Printf("error updating last heartbeat from machine-id: %v %v", machineID, err)
		respErr := response.MapErrorToErrorResponse(response.ErrCodeInvalidInput, response.ErrorDetail{
			Field:   "machine_id",
			Message: fmt.Sprintf("invalid machine_id: %s", machineID),
		})
		response.RespondWithError(c, respErr.RespCode, respErr.ResponseError)
		return
	}

	response.RespondWithJSON(c, http.StatusOK, map[string]string{"message": "Updated successfully"})
}
