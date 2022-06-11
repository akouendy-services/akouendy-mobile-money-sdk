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

type MerchantInitPaymentQrcode struct {
	Amount int64 `json:"Amount"`
	Code string `json:"Code"`
	Desc string `json:"Desc"`
	Reference string `json:"Reference"`
}
