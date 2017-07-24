/* 
 * SCOIR API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.0.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type PreAct struct {

	// Id is assigned when the score is created
	Id string `json:"Id,omitempty"`

	// UpdatedBy is managed by SCOIR
	UpdatedBy string `json:"UpdatedBy,omitempty"`

	// UpdatedAt is managed by SCOIR
	UpdatedAt string `json:"UpdatedAt,omitempty"`

	TestDate string `json:"TestDate"`

	EnglishScore int32 `json:"EnglishScore"`

	MathScore int32 `json:"MathScore"`

	ReadingScore int32 `json:"ReadingScore"`

	ScienceScore int32 `json:"ScienceScore"`

	STEMScore int32 `json:"STEMScore"`
}