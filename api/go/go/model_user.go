/*
 * test
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// User - 
type User struct {

	// Unique identifier for the given user.
	Id int32 `json:"id"`

	FirstName string `json:"firstName"`

	LastName string `json:"lastName"`

	Email string `json:"email"`

	DateOfBirth string `json:"dateOfBirth,omitempty"`

	// Set to true if the user's email has been verified.
	EmailVerified bool `json:"emailVerified"`

	// The date that the user was created.
	CreateDate string `json:"createDate,omitempty"`
}
