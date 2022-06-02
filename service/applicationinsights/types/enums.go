// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type CloudWatchEventSource string

// Enum values for CloudWatchEventSource
const (
	CloudWatchEventSourceEc2        CloudWatchEventSource = "EC2"
	CloudWatchEventSourceCodeDeploy CloudWatchEventSource = "CODE_DEPLOY"
	CloudWatchEventSourceHealth     CloudWatchEventSource = "HEALTH"
	CloudWatchEventSourceRds        CloudWatchEventSource = "RDS"
)

// Values returns all known values for CloudWatchEventSource. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CloudWatchEventSource) Values() []CloudWatchEventSource {
	return []CloudWatchEventSource{
		"EC2",
		"CODE_DEPLOY",
		"HEALTH",
		"RDS",
	}
}

type ConfigurationEventResourceType string

// Enum values for ConfigurationEventResourceType
const (
	ConfigurationEventResourceTypeCloudwatchAlarm ConfigurationEventResourceType = "CLOUDWATCH_ALARM"
	ConfigurationEventResourceTypeCloudwatchLog   ConfigurationEventResourceType = "CLOUDWATCH_LOG"
	ConfigurationEventResourceTypeCloudformation  ConfigurationEventResourceType = "CLOUDFORMATION"
	ConfigurationEventResourceTypeSsmAssociation  ConfigurationEventResourceType = "SSM_ASSOCIATION"
)

// Values returns all known values for ConfigurationEventResourceType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ConfigurationEventResourceType) Values() []ConfigurationEventResourceType {
	return []ConfigurationEventResourceType{
		"CLOUDWATCH_ALARM",
		"CLOUDWATCH_LOG",
		"CLOUDFORMATION",
		"SSM_ASSOCIATION",
	}
}

type ConfigurationEventStatus string

// Enum values for ConfigurationEventStatus
const (
	ConfigurationEventStatusInfo  ConfigurationEventStatus = "INFO"
	ConfigurationEventStatusWarn  ConfigurationEventStatus = "WARN"
	ConfigurationEventStatusError ConfigurationEventStatus = "ERROR"
)

// Values returns all known values for ConfigurationEventStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ConfigurationEventStatus) Values() []ConfigurationEventStatus {
	return []ConfigurationEventStatus{
		"INFO",
		"WARN",
		"ERROR",
	}
}

type DiscoveryType string

// Enum values for DiscoveryType
const (
	DiscoveryTypeResourceGroupBased DiscoveryType = "RESOURCE_GROUP_BASED"
	DiscoveryTypeAccountBased       DiscoveryType = "ACCOUNT_BASED"
)

// Values returns all known values for DiscoveryType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DiscoveryType) Values() []DiscoveryType {
	return []DiscoveryType{
		"RESOURCE_GROUP_BASED",
		"ACCOUNT_BASED",
	}
}

type FeedbackKey string

// Enum values for FeedbackKey
const (
	FeedbackKeyInsightsFeedback FeedbackKey = "INSIGHTS_FEEDBACK"
)

// Values returns all known values for FeedbackKey. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (FeedbackKey) Values() []FeedbackKey {
	return []FeedbackKey{
		"INSIGHTS_FEEDBACK",
	}
}

type FeedbackValue string

// Enum values for FeedbackValue
const (
	FeedbackValueNotSpecified FeedbackValue = "NOT_SPECIFIED"
	FeedbackValueUseful       FeedbackValue = "USEFUL"
	FeedbackValueNotUseful    FeedbackValue = "NOT_USEFUL"
)

// Values returns all known values for FeedbackValue. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FeedbackValue) Values() []FeedbackValue {
	return []FeedbackValue{
		"NOT_SPECIFIED",
		"USEFUL",
		"NOT_USEFUL",
	}
}

type GroupingType string

// Enum values for GroupingType
const (
	GroupingTypeAccountBased GroupingType = "ACCOUNT_BASED"
)

// Values returns all known values for GroupingType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (GroupingType) Values() []GroupingType {
	return []GroupingType{
		"ACCOUNT_BASED",
	}
}

type LogFilter string

// Enum values for LogFilter
const (
	LogFilterError LogFilter = "ERROR"
	LogFilterWarn  LogFilter = "WARN"
	LogFilterInfo  LogFilter = "INFO"
)

// Values returns all known values for LogFilter. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (LogFilter) Values() []LogFilter {
	return []LogFilter{
		"ERROR",
		"WARN",
		"INFO",
	}
}

type OsType string

// Enum values for OsType
const (
	OsTypeWindows OsType = "WINDOWS"
	OsTypeLinux   OsType = "LINUX"
)

// Values returns all known values for OsType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (OsType) Values() []OsType {
	return []OsType{
		"WINDOWS",
		"LINUX",
	}
}

type SeverityLevel string

// Enum values for SeverityLevel
const (
	SeverityLevelInformative SeverityLevel = "Informative"
	SeverityLevelLow         SeverityLevel = "Low"
	SeverityLevelMedium      SeverityLevel = "Medium"
	SeverityLevelHigh        SeverityLevel = "High"
)

// Values returns all known values for SeverityLevel. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SeverityLevel) Values() []SeverityLevel {
	return []SeverityLevel{
		"Informative",
		"Low",
		"Medium",
		"High",
	}
}

type Status string

// Enum values for Status
const (
	StatusIgnore    Status = "IGNORE"
	StatusResolved  Status = "RESOLVED"
	StatusPending   Status = "PENDING"
	StatusRecurring Status = "RECURRING"
)

// Values returns all known values for Status. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Status) Values() []Status {
	return []Status{
		"IGNORE",
		"RESOLVED",
		"PENDING",
		"RECURRING",
	}
}

type Tier string

// Enum values for Tier
const (
	TierCustom                             Tier = "CUSTOM"
	TierDefault                            Tier = "DEFAULT"
	TierDotNetCore                         Tier = "DOT_NET_CORE"
	TierDotNetWorker                       Tier = "DOT_NET_WORKER"
	TierDotNetWebTier                      Tier = "DOT_NET_WEB_TIER"
	TierDotNetWeb                          Tier = "DOT_NET_WEB"
	TierSqlServer                          Tier = "SQL_SERVER"
	TierSqlServerAlwaysonAvailabilityGroup Tier = "SQL_SERVER_ALWAYSON_AVAILABILITY_GROUP"
	TierMysql                              Tier = "MYSQL"
	TierPostgresql                         Tier = "POSTGRESQL"
	TierJavaJmx                            Tier = "JAVA_JMX"
	TierOracle                             Tier = "ORACLE"
	TierSapHanaMultiNode                   Tier = "SAP_HANA_MULTI_NODE"
	TierSapHanaSingleNode                  Tier = "SAP_HANA_SINGLE_NODE"
	TierSapHanaHighAvailability            Tier = "SAP_HANA_HIGH_AVAILABILITY"
	TierSqlServerFailoverClusterInstance   Tier = "SQL_SERVER_FAILOVER_CLUSTER_INSTANCE"
	TierSharepoint                         Tier = "SHAREPOINT"
	TierActiveDirectory                    Tier = "ACTIVE_DIRECTORY"
)

// Values returns all known values for Tier. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Tier) Values() []Tier {
	return []Tier{
		"CUSTOM",
		"DEFAULT",
		"DOT_NET_CORE",
		"DOT_NET_WORKER",
		"DOT_NET_WEB_TIER",
		"DOT_NET_WEB",
		"SQL_SERVER",
		"SQL_SERVER_ALWAYSON_AVAILABILITY_GROUP",
		"MYSQL",
		"POSTGRESQL",
		"JAVA_JMX",
		"ORACLE",
		"SAP_HANA_MULTI_NODE",
		"SAP_HANA_SINGLE_NODE",
		"SAP_HANA_HIGH_AVAILABILITY",
		"SQL_SERVER_FAILOVER_CLUSTER_INSTANCE",
		"SHAREPOINT",
		"ACTIVE_DIRECTORY",
	}
}
