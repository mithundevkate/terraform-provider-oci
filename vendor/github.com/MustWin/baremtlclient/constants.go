package baremtlsdk

type resourceName string

const (
	// Identity States
	ResourceCreated  = "CREATED"
	ResourceCreating = "CREATING"

	// Core API States
	ResourceProvisioning = "PROVISIONING"
	ResourceAvailable    = "AVAILABLE"
	ResourceTerminating  = "TERMINATING"
	ResourceTerminated   = "TERMINATED"

	identityServiceAPI        = "https://identity.us-az-phoenix-1.OracleIaaS.com"
	identityServiceAPIVersion = "v1"

	coreServiceAPI        = "https://core.us-az-phoenix-1.OracleIaaS.com"
	coreServiceAPIVersion = "v1"

	// Header Keys
	headerOPCIdempotencyToken = "opc-idempotency-token"
	headerOPCNextPage         = "opc-next-page"
	headerIfMatch             = "if-match"
	headerOPCRequestID        = "opc-request-id"

	// URL Query Keys
	queryAvailabilityDomain = "availabilityDomain"
	queryCompartmentID      = "compartmentId"
	queryGroupID            = "groupId"
	queryImageID            = "imageId"
	queryLimit              = "limit"
	queryPage               = "page"
	queryUserID             = "userId"

	// Identity Resources
	resourceAvailabilityDomains  resourceName = "availabilityDomains"
	resourceCompartments         resourceName = "compartments"
	resourceGroups               resourceName = "groups"
	resourcePolicies             resourceName = "policies"
	resourceUsers                resourceName = "users"
	resourceUserGroupMemberships resourceName = "userGroupMemberships"

	// Core Resources
	resourceCustomerPremiseEquipment resourceName = "cpes"
	resourceShapes                   resourceName = "shapes"

	apiKeys    = "apiKeys"
	uiPassword = "uiPassword"
)