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

type SharedAccountProfile struct {
	Balance       *Balance       `json:"balance"`
	Barred        string         `json:"barred"`
	FirstName     string         `json:"firstName"`
	FrozenBalance *FrozenBalance `json:"frozenBalance"`
	Grade         string         `json:"grade"`
	LastName      string         `json:"lastName"`
	Msisdn        string         `json:"msisdn"`
	Suspended     string         `json:"suspended"`
	Type_         string         `json:"type"`
	UserId        string         `json:"userId"`
}