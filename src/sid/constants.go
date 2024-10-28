package sid

const (
	WELLKNOWNSID_NOBODY                              = "S-1-0-0"
	WELLKNOWNSID_EVERYONE                            = "S-1-1-0"
	WELLKNOWNSID_LOCAL                               = "S-1-2-0"
	WELLKNOWNSID_CONSOLE_LOGON                       = "S-1-2-1"
	WELLKNOWNSID_CREATOR_OWNER                       = "S-1-3-0"
	WELLKNOWNSID_CREATOR_GROUP                       = "S-1-3-1"
	WELLKNOWNSID_CREATOR_OWNER_SERVER                = "S-1-3-2"
	WELLKNOWNSID_CREATOR_GROUP_SERVER                = "S-1-3-3"
	WELLKNOWNSID_NT_AUTHORITY                        = "S-1-5"
	WELLKNOWNSID_DIALUP                              = "S-1-5-1"
	WELLKNOWNSID_NETWORK                             = "S-1-5-2"
	WELLKNOWNSID_BATCH                               = "S-1-5-3"
	WELLKNOWNSID_INTERACTIVE                         = "S-1-5-4"
	WELLKNOWNSID_SERVICE                             = "S-1-5-6"
	WELLKNOWNSID_ANONYMOUS                           = "S-1-5-7"
	WELLKNOWNSID_PROXY                               = "S-1-5-8"
	WELLKNOWNSID_ENTERPRISE_DOMAIN_CONTROLLERS       = "S-1-5-9"
	WELLKNOWNSID_PRINCIPAL_SELF                      = "S-1-5-10"
	WELLKNOWNSID_AUTHENTICATED_USERS                 = "S-1-5-11"
	WELLKNOWNSID_RESTRICTED_CODE                     = "S-1-5-12"
	WELLKNOWNSID_TERMINAL_SERVER_USERS               = "S-1-5-13"
	WELLKNOWNSID_REMOTE_INTERACTIVE_LOGON            = "S-1-5-14"
	WELLKNOWNSID_THIS_ORGANIZATION                   = "S-1-5-15"
	WELLKNOWNSID_IUSR                                = "S-1-5-17"
	WELLKNOWNSID_LOCAL_SYSTEM                        = "S-1-5-18"
	WELLKNOWNSID_LOCAL_SERVICE                       = "S-1-5-19"
	WELLKNOWNSID_NETWORK_SERVICE                     = "S-1-5-20"
	WELLKNOWNSID_ADMINISTRATORS                      = "S-1-5-32-544"
	WELLKNOWNSID_USERS                               = "S-1-5-32-545"
	WELLKNOWNSID_GUESTS                              = "S-1-5-32-546"
	WELLKNOWNSID_POWER_USERS                         = "S-1-5-32-547"
	WELLKNOWNSID_ACCOUNT_OPERATORS                   = "S-1-5-32-548"
	WELLKNOWNSID_SERVER_OPERATORS                    = "S-1-5-32-549"
	WELLKNOWNSID_PRINT_OPERATORS                     = "S-1-5-32-550"
	WELLKNOWNSID_BACKUP_OPERATORS                    = "S-1-5-32-551"
	WELLKNOWNSID_REPLICATORS                         = "S-1-5-32-552"
	WELLKNOWNSID_NTLM_AUTHENTICATION                 = "S-1-5-64-10"
	WELLKNOWNSID_SCHANNEL_AUTHENTICATION             = "S-1-5-64-14"
	WELLKNOWNSID_DIGEST_AUTHENTICATION               = "S-1-5-64-21"
	WELLKNOWNSID_UNTRUSTED_LEVEL                     = "S-1-16-0"
	WELLKNOWNSID_LOW_INTEGRITY_LEVEL                 = "S-1-16-4096"
	WELLKNOWNSID_MEDIUM_INTEGRITY_LEVEL              = "S-1-16-8192"
	WELLKNOWNSID_MEDIUM_PLUS_INTEGRITY_LEVEL         = "S-1-16-8448"
	WELLKNOWNSID_HIGH_INTEGRITY_LEVEL                = "S-1-16-12288"
	WELLKNOWNSID_SYSTEM_INTEGRITY_LEVEL              = "S-1-16-16384"
	WELLKNOWNSID_PROTECTED_PROCESS                   = "S-1-16-20480"
	WELLKNOWNSID_SECURE_PROCESS                      = "S-1-16-28672"
	WELLKNOWNSID_ADMINISTRATOR_ACCOUNT               = "S-1-5-21-0-0-0-500"
	WELLKNOWNSID_GUEST_ACCOUNT                       = "S-1-5-21-0-0-0-501"
	WELLKNOWNSID_KRBTGT_ACCOUNT                      = "S-1-5-21-0-0-0-502"
	WELLKNOWNSID_DOMAIN_ADMINS                       = "S-1-5-21-0-0-0-512"
	WELLKNOWNSID_DOMAIN_USERS                        = "S-1-5-21-0-0-0-513"
	WELLKNOWNSID_DOMAIN_GUESTS                       = "S-1-5-21-0-0-0-514"
	WELLKNOWNSID_DOMAIN_COMPUTERS                    = "S-1-5-21-0-0-0-515"
	WELLKNOWNSID_DOMAIN_CONTROLLERS                  = "S-1-5-21-0-0-0-516"
	WELLKNOWNSID_CERT_PUBLISHERS                     = "S-1-5-21-0-0-0-517"
	WELLKNOWNSID_SCHEMA_ADMINS                       = "S-1-5-21-0-0-0-518"
	WELLKNOWNSID_ENTERPRISE_ADMINS                   = "S-1-5-21-0-0-0-519"
	WELLKNOWNSID_GROUP_POLICY_CREATOR_OWNERS         = "S-1-5-21-0-0-0-520"
	WELLKNOWNSID_READ_ONLY_DOMAIN_CONTROLLERS        = "S-1-5-21-0-0-0-521"
	WELLKNOWNSID_CLONEABLE_DOMAIN_CONTROLLERS        = "S-1-5-21-0-0-0-522"
	WELLKNOWNSID_RAS_SERVERS_GROUP                   = "S-1-5-21-0-0-0-553"
	WELLKNOWNSID_PRE_WINDOWS_2000_COMPATIBLE_ACCESS  = "S-1-5-32-554"
	WELLKNOWNSID_REMOTE_DESKTOP_USERS                = "S-1-5-32-555"
	WELLKNOWNSID_NETWORK_CONFIGURATION_OPERATORS     = "S-1-5-32-556"
	WELLKNOWNSID_INCOMING_FOREST_TRUST_BUILDERS      = "S-1-5-32-557"
	WELLKNOWNSID_PERFORMANCE_MONITOR_USERS           = "S-1-5-32-558"
	WELLKNOWNSID_PERFORMANCE_LOG_USERS               = "S-1-5-32-559"
	WELLKNOWNSID_WINDOWS_AUTHORIZATION_ACCESS_GROUP  = "S-1-5-32-560"
	WELLKNOWNSID_TERMINAL_SERVER_LICENSE_SERVERS     = "S-1-5-32-561"
	WELLKNOWNSID_DISTRIBUTED_COM_USERS               = "S-1-5-32-562"
	WELLKNOWNSID_CRYPTOGRAPHIC_OPERATORS             = "S-1-5-32-569"
	WELLKNOWNSID_EVENT_LOG_READERS                   = "S-1-5-32-573"
	WELLKNOWNSID_CERTIFICATE_SERVICE_DCOM_ACCESS     = "S-1-5-32-574"
	WELLKNOWNSID_RDS_REMOTE_ACCESS_SERVERS           = "S-1-5-32-575"
	WELLKNOWNSID_RDS_ENDPOINT_SERVERS                = "S-1-5-32-576"
	WELLKNOWNSID_RDS_MANAGEMENT_SERVERS              = "S-1-5-32-577"
	WELLKNOWNSID_HYPER_V_ADMINISTRATORS              = "S-1-5-32-578"
	WELLKNOWNSID_ACCESS_CONTROL_ASSISTANCE_OPERATORS = "S-1-5-32-579"
	WELLKNOWNSID_REMOTE_MANAGEMENT_USERS             = "S-1-5-32-580"
)

// WellKnownSIDs maps some well-known SIDs to their names.
var WellKnownSIDs = map[string]string{
	// Operating system-defined SIDs
	"S-1-0-0":  "Nobody",
	"S-1-1-0":  "Everyone",
	"S-1-2-0":  "Local",
	"S-1-2-1":  "Console Logon",
	"S-1-3-0":  "Creator Owner",
	"S-1-3-1":  "Creator Group",
	"S-1-3-2":  "Creator Owner Server",
	"S-1-3-3":  "Creator Group Server",
	"S-1-5":    "NT Authority",
	"S-1-5-1":  "Dialup",
	"S-1-5-2":  "Network",
	"S-1-5-3":  "Batch",
	"S-1-5-4":  "Interactive",
	"S-1-5-6":  "Service",
	"S-1-5-7":  "Anonymous",
	"S-1-5-8":  "Proxy",
	"S-1-5-9":  "Enterprise Domain Controllers",
	"S-1-5-10": "Principal Self",
	"S-1-5-11": "Authenticated Users",
	"S-1-5-12": "Restricted Code",
	"S-1-5-13": "Terminal Server Users",
	"S-1-5-14": "Remote Interactive Logon",
	"S-1-5-15": "This Organization",
	"S-1-5-17": "IUSR",
	"S-1-5-18": "Local System",
	"S-1-5-19": "Local Service",
	"S-1-5-20": "Network Service",

	// Built-in system groups
	"S-1-5-32-544": "Administrators",
	"S-1-5-32-545": "Users",
	"S-1-5-32-546": "Guests",
	"S-1-5-32-547": "Power Users",
	"S-1-5-32-548": "Account Operators",
	"S-1-5-32-549": "Server Operators",
	"S-1-5-32-550": "Print Operators",
	"S-1-5-32-551": "Backup Operators",
	"S-1-5-32-552": "Replicators",

	// Logon types
	"S-1-5-64-10": "NTLM Authentication",
	"S-1-5-64-14": "SChannel Authentication",
	"S-1-5-64-21": "Digest Authentication",

	// Mandatory integrity levels
	"S-1-16-0":     "Untrusted Level",
	"S-1-16-4096":  "Low Integrity Level",
	"S-1-16-8192":  "Medium Integrity Level",
	"S-1-16-8448":  "Medium Plus Integrity Level",
	"S-1-16-12288": "High Integrity Level",
	"S-1-16-16384": "System Integrity Level",
	"S-1-16-20480": "Protected Process",
	"S-1-16-28672": "Secure Process",

	// Special identity groups
	"S-1-5-21-0-0-0-500": "Administrator Account",
	"S-1-5-21-0-0-0-501": "Guest Account",
	"S-1-5-21-0-0-0-502": "KRBTGT Account",
	"S-1-5-21-0-0-0-512": "Domain Admins",
	"S-1-5-21-0-0-0-513": "Domain Users",
	"S-1-5-21-0-0-0-514": "Domain Guests",
	"S-1-5-21-0-0-0-515": "Domain Computers",
	"S-1-5-21-0-0-0-516": "Domain Controllers",
	"S-1-5-21-0-0-0-517": "Cert Publishers",
	"S-1-5-21-0-0-0-518": "Schema Admins",
	"S-1-5-21-0-0-0-519": "Enterprise Admins",
	"S-1-5-21-0-0-0-520": "Group Policy Creator Owners",
	"S-1-5-21-0-0-0-521": "Read-Only Domain Controllers",
	"S-1-5-21-0-0-0-522": "Cloneable Domain Controllers",
	"S-1-5-21-0-0-0-553": "RAS Servers Group",

	// Others
	"S-1-5-32-554": "Pre-Windows 2000 Compatible Access",
	"S-1-5-32-555": "Remote Desktop Users",
	"S-1-5-32-556": "Network Configuration Operators",
	"S-1-5-32-557": "Incoming Forest Trust Builders",
	"S-1-5-32-558": "Performance Monitor Users",
	"S-1-5-32-559": "Performance Log Users",
	"S-1-5-32-560": "Windows Authorization Access Group",
	"S-1-5-32-561": "Terminal Server License Servers",
	"S-1-5-32-562": "Distributed COM Users",
	"S-1-5-32-569": "Cryptographic Operators",
	"S-1-5-32-573": "Event Log Readers",
	"S-1-5-32-574": "Certificate Service DCOM Access",
	"S-1-5-32-575": "RDS Remote Access Servers",
	"S-1-5-32-576": "RDS Endpoint Servers",
	"S-1-5-32-577": "RDS Management Servers",
	"S-1-5-32-578": "Hyper-V Administrators",
	"S-1-5-32-579": "Access Control Assistance Operators",
	"S-1-5-32-580": "Remote Management Users",
}
