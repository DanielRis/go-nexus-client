package security

const (
	// PrivilegeDomains
	PrivilegeDomainAll                = "*"
	PrivilegeDomainAPIKey             = "apikey"
	PrivilegeDomainAnalytics          = "analytics"
	PrivilegeDomainAtlas              = "atlas"
	PrivilegeDomainBlobstores         = "blobstores"
	PrivilegeDomainBundles            = "bundles"
	PrivilegeDomainCapabilities       = "capabilities"
	PrivilegeDomainComponent          = "component"
	PrivilegeDomainDatastores         = "datastores"
	PrivilegeDomainHealthcheck        = "healthcheck"
	PrivilegeDomainHealthcheckSummary = "healthchecksummary"
	PrivilegeDomainIQViolationSummery = "iq-violation-summary"
	PrivilegeDomainLDAP               = "ldap"
	PrivilegeDomainLicensing          = "licensing"
	PrivilegeDomainLogging            = "logging"
	PrivilegeDomainMetrics            = "metrics"
	PrivilegeDomainPrivileges         = "privileges"
	PrivilegeDomainRoles              = "roles"
	PrivilegeDomainSearch             = "search"
	PrivilegeDomainSelectors          = "selectors"
	PrivilegeDomainSettings           = "settings"
	PrivilegeDomainSSLTruststore      = "ssl-truststore"
	PrivilegeDomainTasks              = "tasks"
	PrivilegeDomainUsers              = "users"
	PrivilegeDomainUsersChangePW      = "userschangepw"
	PrivilegeDomainWonderland         = "wonderland"

	// PrivilegeTypes
	PrivilegeTypeApplication     = "application"
	PrivilegeTypeContentSelector = "repository-content-selector"
	PrivilegeTypeRepositoryAdmin = "repository-admin"
	PrivilegeTypeRepositoryView  = "repository-view"
	PrivilegeTypeScript          = "script"
	PrivilegeTypeWildcard        = "wildcard"
)

var (
	// PrivilegeDomains represents a string slice of supported privilege domains
	PrivilegeDomains []string = []string{
		PrivilegeDomainAll,
		PrivilegeDomainAPIKey,
		PrivilegeDomainAnalytics,
		PrivilegeDomainAtlas,
		PrivilegeDomainBlobstores,
		PrivilegeDomainBundles,
		PrivilegeDomainCapabilities,
		PrivilegeDomainComponent,
		PrivilegeDomainDatastores,
		PrivilegeDomainHealthcheck,
		PrivilegeDomainHealthcheckSummary,
		PrivilegeDomainIQViolationSummery,
		PrivilegeDomainLDAP,
		PrivilegeDomainLicensing,
		PrivilegeDomainLogging,
		PrivilegeDomainMetrics,
		PrivilegeDomainPrivileges,
		PrivilegeDomainRoles,
		PrivilegeDomainSearch,
		PrivilegeDomainSelectors,
		PrivilegeDomainSettings,
		PrivilegeDomainSSLTruststore,
		PrivilegeDomainTasks,
		PrivilegeDomainUsers,
		PrivilegeDomainUsersChangePW,
		PrivilegeDomainWonderland,
	}
	// PrivilegeTypes represents a string slice of possible privilege types
	PrivilegeTypes []string = []string{
		PrivilegeTypeApplication,
		PrivilegeTypeContentSelector,
		PrivilegeTypeRepositoryAdmin,
		PrivilegeTypeRepositoryView,
		PrivilegeTypeScript,
		PrivilegeTypeWildcard,
	}
)

// Privilege data
type Privilege struct {
	Actions         []string `json:"actions,omitempty"`
	ContentSelector string   `json:"contentSelector,omitempty"`
	Description     string   `json:"description"`
	Domain          string   `json:"domain,omitempty"`
	Format          string   `json:"format,omitempty"`
	Name            string   `json:"name"`
	Pattern         string   `json:"pattern,omitempty"`
	ScriptName      string   `json:"scriptName,omitempty"`
	ReadOnly        bool     `json:"readOnly"`
	Repository      string   `json:"repository,omitempty"`
	Type            string   `json:"type"`
}
