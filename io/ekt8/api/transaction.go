package api

import (
	"encoding/json"

	"github.com/EducationEKT/EKT/io/ekt8/core/common"
	"github.com/EducationEKT/EKT/io/ekt8/dispatcher"
	"github.com/EducationEKT/xserver/x_err"
	"github.com/EducationEKT/xserver/x_http/x_req"
	"github.com/EducationEKT/xserver/x_http/x_resp"
	"github.com/EducationEKT/xserver/x_http/x_router"
)

func init() {
	x_router.Post("/transaction/api/newTransaction", newTransaction)
}

func newTransaction(req *x_req.XReq) (*x_resp.XRespContainer, *x_err.XErr) {
	var tx common.Transaction
	err := json.Unmarshal(req.Body, &tx)
	if err != nil {
		return nil, x_err.New(-1, err.Error())
	}
	dispatcher.GetDisPatcher().NewTransaction(&tx)
	// TODO 根据mainchain的记录进行转发
	return x_resp.Success("success"), nil
}
