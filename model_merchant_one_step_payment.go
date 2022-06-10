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

type MerchantOneStepPayment struct {
	Amount *Amount `json:"amount"`
	Customer *Customer `json:"customer"`
	Partner *Partner `json:"partner"`
	Reference string `json:"reference"`
}
