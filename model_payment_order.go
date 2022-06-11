/*
 * Documentation
 *
 * Resource for api
 *
 * API version: 1.0.0
 * Contact: contact@akouendy.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type PaymentOrder struct {
	AppId string `json:"AppId"`
	AppName string `json:"AppName"`
	Customer string `json:"Customer"`
	ExpireDate time.Time `json:"ExpireDate"`
	ExpireNotify bool `json:"ExpireNotify"`
	ID string `json:"ID"`
	LastPaymentDate time.Time `json:"LastPaymentDate"`
	Status string `json:"Status"`
	TotalAmount string `json:"TotalAmount"`
}
