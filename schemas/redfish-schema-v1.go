//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

// releaseStatus is The term specifies the public release status of a property
// or schema.
type releaseStatus string

const (
	// StandardreleaseStatus means Standard
	StandardreleaseStatus releaseStatus = "Standard"
	// InformationalreleaseStatus means Informational
	InformationalreleaseStatus releaseStatus = "Informational"
	// WorkInProgressreleaseStatus means WorkInProgress
	WorkInProgressreleaseStatus releaseStatus = "WorkInProgress"
	// InDevelopmentreleaseStatus means InDevelopment
	InDevelopmentreleaseStatus releaseStatus = "InDevelopment"
)

// actionResponse Contains the definition of the response body for a given
// action if it does not follow the Redfish Error format.
type actionResponse struct {
}

// enumDeprecated shall be applied to a value in order to specify that the value
// is deprecated. The value of the string should explain the deprecation,
// including new value to be used. The value can be supported in new and
// existing implementations, but usage in new implementations is discouraged.
// Deprecated values are likely to be removed in a future major version of the
// schema.
type enumDeprecated struct {
}

// enumDescriptions shall contain informative language related to the
// enumeration values of the property.
type enumDescriptions struct {
}

// enumLongDescriptions shall contain normative language relating to the
// enumeration values of the property.
type enumLongDescriptions struct {
}

// enumTranslations shall contain properties with names that match the
// enumerations and values that contain the translation of the enumeration
// value. The language used shall match the 'language' value for this schema.
type enumTranslations struct {
}

// enumVersionAdded shall be applied to a value in order to specify when the
// value was added. The value of the string should contain the resource version
// where the value was added.
type enumVersionAdded struct {
}

// enumVersionDeprecated shall be applied to a value in order to specify when
// the value was deprecated. The value of the string should contain the resource
// version where the value was deprecated.
type enumVersionDeprecated struct {
}
