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

type SharedChangePinIdentifier struct {
	Msisdn string `json:"Msisdn"`
	EncryptedNewPinCode string `json:"encryptedNewPinCode"`
	EncryptedPinCode string `json:"encryptedPinCode"`
	IdType string `json:"idType"`
	NewPinCode string `json:"newPinCode"`
	PinCode string `json:"pinCode"`
}
