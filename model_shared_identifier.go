/*
 * Documentation
 *
 * Resource for api
 *
 * API version: 1.0.0
 * Contact: andictl@andi.dev
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package mobilemoneysdk

type SharedIdentifier struct {
	EncryptedPinCode string `json:"encryptedPinCode"`
	Id               string `json:"id"`
	IdType           string `json:"idType"`
	PinCode          string `json:"pinCode"`
	Wallet           string `json:"wallet"`
}
