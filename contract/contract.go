package contract

//
//import (
//	"encoding/hex"
//	"time"
//
//	"github.com/EducationEKT/EKT/context"
//	"github.com/EducationEKT/EKT/core/types"
//	"github.com/EducationEKT/EKT/core/userevent"
//)
//
//const Contract_Call_Timeout = 200 * time.Millisecond
//
//func Run(ctx *context.Sticker, tx userevent.Transaction, account *types.Account) (*userevent.TransactionReceipt, []byte) {
//	c := getContract(ctx, tx.To, account)
//	if c == nil {
//		return userevent.ContractRefuseTx(tx), nil
//	}
//	receipt, data := CallWithTimeout(c, tx)
//	//receipt, data := c.Call(tx)
//	if receipt == nil {
//		receipt = userevent.ContractRefuseTx(tx)
//	}
//	return receipt, data
//}
//
//type ReceiptAndContractData struct {
//	Reciept      *userevent.TransactionReceipt
//	ContractData []byte
//}
//
//func CallWithTimeout(c IContract, tx userevent.Transaction) (*userevent.TransactionReceipt, []byte) {
//	ch := make(chan ReceiptAndContractData)
//	go async_call(c, tx, ch)
//	for {
//		select {
//		case <-time.After(Contract_Call_Timeout):
//			reciept := userevent.NewTransactionReceipt(tx, false, userevent.FailType_CONTRACT_TIMEOUT)
//			return &reciept, nil
//		case result := <-ch:
//			return result.Reciept, result.ContractData
//		}
//	}
//}
//
//func async_call(c IContract, tx userevent.Transaction, ch chan ReceiptAndContractData) {
//	receipt, data := c.Call(tx)
//	if receipt == nil {
//		receipt = userevent.ContractRefuseTx(tx)
//	}
//	result := ReceiptAndContractData{
//		Reciept:      receipt,
//		ContractData: data,
//	}
//	ch <- result
//}
//
//func InitContractAccount(tx userevent.Transaction, account *types.Account) bool {
//	switch hex.EncodeToString(tx.To[:32]) {
//	case SYSTEM_AUTHOR:
//		switch hex.EncodeToString(tx.To[32:]) {
//		case EKT_GAS_BANCOR_CONTRACT:
//			contract := types.NewContractAccount(tx.To[32:], nil, types.ContractData{})
//			contract.Gas = 1e8
//			if account.Contracts == nil {
//				account.Contracts = make(map[string]types.ContractAccount)
//			}
//			account.Contracts[hex.EncodeToString(tx.To[32:])] = *contract
//			return true
//		}
//	}
//	return false
//}
