//
// Copyright 2019, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package kepserverex

import (
	"fmt"
	"net/url"
)

// DeviceService handles communication with the device related methods
// of the KEPServerEX API.
type DeviceService struct {
	client *Client
}

// Device represents a KEPServerEX device.
type Device struct {
	Name              string `json:"common.ALLTYPES_NAME"`
	Description       string `json:"common.ALLTYPES_DESCRIPTION"`
	UniqueID          int64  `json:"servermain.DEVICE_UNIQUE_ID"`
	ProjectID         int64  `json:"PROJECT_ID"`
	ChannelAssignment string `json:"servermain.DEVICE_CHANNEL_ASSIGNMENT"`
	Driver            string `json:"servermain.MULTIPLE_TYPES_DEVICE_DRIVER"`
}

// ControlLogixEthernetDevice represents an ControlLogix Ethernet device.
type ControlLogixEthernetDevice struct {
	Name                                 string                    `json:"common.ALLTYPES_NAME"`
	Description                          string                    `json:"common.ALLTYPES_DESCRIPTION"`
	UniqueID                             int64                     `json:"servermain.DEVICE_UNIQUE_ID"`
	ProjectID                            int64                     `json:"PROJECT_ID"`
	ChannelAssignment                    string                    `json:"servermain.DEVICE_CHANNEL_ASSIGNMENT"`
	Driver                               string                    `json:"servermain.MULTIPLE_TYPES_DEVICE_DRIVER"`
	Model                                ControlLogixEthernetModel `json:"servermain.DEVICE_MODEL"`
	IDFormat                             IDFormat                  `json:"servermain.DEVICE_ID_FORMAT"`
	IDString                             string                    `json:"servermain.DEVICE_ID_STRING"`
	IDHexadecimal                        int                       `json:"servermain.DEVICE_ID_HEXADECIMAL"`
	IDDecimal                            int                       `json:"servermain.DEVICE_ID_DECIMAL"`
	IDOctal                              int                       `json:"servermain.DEVICE_ID_OCTAL"`
	DataCollection                       bool                      `json:"servermain.DEVICE_DATA_COLLECTION"`
	Simulated                            bool                      `json:"servermain.DEVICE_SIMULATED"`
	ScanMode                             ScanMode                  `json:"servermain.DEVICE_SCAN_MODE"`
	ScanRate                             int                       `json:"servermain.DEVICE_SCAN_MODE_RATE_MS"`
	InitialUpdatesFromCache              bool                      `json:"servermain.DEVICE_SCAN_MODE_PROVIDE_INITIAL_UPDATES_FROM_CACHE"`
	ConnectionTimeout                    int                       `json:"servermain.DEVICE_CONNECTION_TIMEOUT_SECONDS"`
	RequestTimeout                       int                       `json:"servermain.DEVICE_REQUEST_TIMEOUT_MILLISECONDS"`
	AttemptsBeforeTimeout                int                       `json:"servermain.DEVICE_RETRY_ATTEMPTS"`
	InterRequestDelay                    int                       `json:"servermain.DEVICE_INTER_REQUEST_DELAY_MILLISECONDS"`
	DemoteOnFailure                      bool                      `json:"servermain.DEVICE_AUTO_DEMOTION_ENABLE_ON_COMMUNICATIONS_FAILURES"`
	TimeoutToDemote                      int                       `json:"servermain.DEVICE_AUTO_DEMOTION_DEMOTE_AFTER_SUCCESSIVE_TIMEOUTS"`
	DemotionPeriod                       int                       `json:"servermain.DEVICE_AUTO_DEMOTION_PERIOD_MS"`
	DiscardRequestsWhenDemoted           bool                      `json:"servermain.DEVICE_AUTO_DEMOTION_DISCARD_WRITES"`
	OnDeviceStartup                      OnDeviceStartup           `json:"servermain.DEVICE_TAG_GENERATION_ON_STARTUP"`
	OnDuplicateTag                       OnDuplicateTag            `json:"servermain.DEVICE_TAG_GENERATION_DUPLICATE_HANDLING"`
	ParentGroup                          string                    `json:"servermain.DEVICE_TAG_GENERATION_GROUP"`
	AllowAutomaticallyGeneratedSubgroups bool                      `json:"servermain.DEVICE_TAG_GENERATION_ALLOW_SUB_GROUPS"`
	TCPIPPort                            int                       `json:"controllogix_ethernet.DEVICE_PORT_NUMBER"`
	ConnectionSize                       int                       `json:"controllogix_ethernet.DEVICE_CONNECTION_SIZE_BYTES"`
	InactivityWatchdog                   InactivityWatchdog        `json:"controllogix_ethernet.DEVICE_INACTIVITY_WATCHDOG_SECONDS"`
	ArrayBlockSize                       ArrayBlockSize            `json:"controllogix_ethernet.DEVICE_ARRAY_BLOCK_SIZE_ELEMENTS"`
	ProtocolMode                         ProtocolMode              `json:"controllogix_ethernet.DEVICE_PROTOCOL_MODE"`
	SynchronizeAfterOnlineEdits          bool                      `json:"controllogix_ethernet.DEVICE_ONLINE_EDITS"`
	SynchronizeAfterOfflineEdits         bool                      `json:"controllogix_ethernet.DEVICE_OFFLINE_EDITS"`
	TerminateStringDataAtLEN             bool                      `json:"controllogix_ethernet.DEVICE_AUTOMATICALLY_READ_STRING_LENGTH"`
	DefaultDataType                      DataType                  `json:"controllogix_ethernet.DEVICE_DEFAULT_DATA_TYPE"`
	PerformanceStatistics                bool                      `json:"controllogix_ethernet.DEVICE_ENABLE_PERFORMANCE_STATISTICS"`
	DatabaseImportMethod                 ImportMethod              `json:"controllogix_ethernet.DEVICE_DATABASE_IMPORT_METHOD"`
	TagImportFile                        string                    `json:"controllogix_ethernet.DEVICE_TAG_IMPORT_FILE"`
	TagDescriptions                      bool                      `json:"controllogix_ethernet.DEVICE_DISPLAY_DESCRIPTIONS"`
	LimitNameLength                      bool                      `json:"controllogix_ethernet.DEVICE_LIMIT_TAG_NAMES"`
	TagHierarchy                         TagHierarchy              `json:"controllogix_ethernet.DEVICE_TAG_HIERARCHY"`
	ImposeArrayLimit                     bool                      `json:"controllogix_ethernet.DEVICE_IMPOSE_ARRAY_ELEMENT_COUNT_LIMIT"`
	ArrayCountUpperLimit                 int                       `json:"controllogix_ethernet.DEVICE_ARRAY_ELEMENT_COUNT_LIMIT"`
	RemoteTCPIPPort                      int                       `json:"controllogix_ethernet.DEVICE_CL_ENET_PORT_NUMBER"`
	RequestSize                          RequestSize               `json:"controllogix_ethernet.DEVICE_REQUEST_SIZE"`
	AllowFunctionFileBlockWrites         bool                      `json:"controllogix_ethernet.DEVICE_PERFORM_BLOCK_WRITES"`
	Slot                                 struct {
		Module      []int `json:"controllogix_ethernet.DEVICE_SLOT_CONFIGURATION_MODULE"`
		InputWords  []int `json:"controllogix_ethernet.DEVICE_SLOT_CONFIGURATION_INPUT_WORDS"`
		OutputWords []int `json:"controllogix_ethernet.DEVICE_SLOT_CONFIGURATION_OUTPUT_WORDS"`
	} `json:"controllogix_ethernet.DEVICE_SLOT_CONFIGURATION"`
}

// ControlLogixEthernetDeviceOptions represents all ControlLogix Ethernet device options.
type ControlLogixEthernetDeviceOptions struct {
	Name                                 string                    `json:"common.ALLTYPES_NAME,omitempty"`
	Description                          string                    `json:"common.ALLTYPES_DESCRIPTION,omitempty"`
	Driver                               string                    `json:"servermain.MULTIPLE_TYPES_DEVICE_DRIVER"`
	Model                                ControlLogixEthernetModel `json:"servermain.DEVICE_MODEL"`
	IDFormat                             IDFormat                  `json:"servermain.DEVICE_ID_FORMAT,omitempty"`
	IDString                             string                    `json:"servermain.DEVICE_ID_STRING,omitempty"`
	IDHexadecimal                        int                       `json:"servermain.DEVICE_ID_HEXADECIMAL,omitempty"`
	IDDecimal                            int                       `json:"servermain.DEVICE_ID_DECIMAL,omitempty"`
	IDOctal                              int                       `json:"servermain.DEVICE_ID_OCTAL,omitempty"`
	DataCollection                       *bool                     `json:"servermain.DEVICE_DATA_COLLECTION,omitempty"`
	Simulated                            *bool                     `json:"servermain.DEVICE_SIMULATED,omitempty"`
	ScanMode                             *ScanMode                 `json:"servermain.DEVICE_SCAN_MODE,omitempty"`
	ScanRate                             *int                      `json:"servermain.DEVICE_SCAN_MODE_RATE_MS,omitempty"`
	InitialUpdatesFromCache              *bool                     `json:"servermain.DEVICE_SCAN_MODE_PROVIDE_INITIAL_UPDATES_FROM_CACHE,omitempty"`
	ConnectionTimeout                    *int                      `json:"servermain.DEVICE_CONNECTION_TIMEOUT_SECONDS,omitempty"`
	RequestTimeout                       *int                      `json:"servermain.DEVICE_REQUEST_TIMEOUT_MILLISECONDS,omitempty"`
	AttemptsBeforeTimeout                *int                      `json:"servermain.DEVICE_RETRY_ATTEMPTS,omitempty"`
	InterRequestDelay                    *int                      `json:"servermain.DEVICE_INTER_REQUEST_DELAY_MILLISECONDS,omitempty"`
	DemoteOnFailure                      *bool                     `json:"servermain.DEVICE_AUTO_DEMOTION_ENABLE_ON_COMMUNICATIONS_FAILURES,omitempty"`
	TimeoutToDemote                      *int                      `json:"servermain.DEVICE_AUTO_DEMOTION_DEMOTE_AFTER_SUCCESSIVE_TIMEOUTS,omitempty"`
	DemotionPeriod                       *int                      `json:"servermain.DEVICE_AUTO_DEMOTION_PERIOD_MS,omitempty"`
	DiscardRequestsWhenDemoted           *bool                     `json:"servermain.DEVICE_AUTO_DEMOTION_DISCARD_WRITES,omitempty"`
	OnDeviceStartup                      *OnDeviceStartup          `json:"servermain.DEVICE_TAG_GENERATION_ON_STARTUP,omitempty"`
	OnDuplicateTag                       *OnDuplicateTag           `json:"servermain.DEVICE_TAG_GENERATION_DUPLICATE_HANDLING,omitempty"`
	ParentGroup                          *string                   `json:"servermain.DEVICE_TAG_GENERATION_GROUP,omitempty"`
	AllowAutomaticallyGeneratedSubgroups *bool                     `json:"servermain.DEVICE_TAG_GENERATION_ALLOW_SUB_GROUPS,omitempty"`
	TCPIPPort                            *int                      `json:"controllogix_ethernet.DEVICE_PORT_NUMBER,omitempty"`
	ConnectionSize                       *int                      `json:"controllogix_ethernet.DEVICE_CONNECTION_SIZE_BYTES,omitempty"`
	InactivityWatchdog                   *InactivityWatchdog       `json:"controllogix_ethernet.DEVICE_INACTIVITY_WATCHDOG_SECONDS,omitempty"`
	ArrayBlockSize                       *ArrayBlockSize           `json:"controllogix_ethernet.DEVICE_ARRAY_BLOCK_SIZE_ELEMENTS,omitempty"`
	ProtocolMode                         *ProtocolMode             `json:"controllogix_ethernet.DEVICE_PROTOCOL_MODE,omitempty"`
	SynchronizeAfterOnlineEdits          *bool                     `json:"controllogix_ethernet.DEVICE_ONLINE_EDITS,omitempty"`
	SynchronizeAfterOfflineEdits         *bool                     `json:"controllogix_ethernet.DEVICE_OFFLINE_EDITS,omitempty"`
	TerminateStringDataAtLEN             *bool                     `json:"controllogix_ethernet.DEVICE_AUTOMATICALLY_READ_STRING_LENGTH,omitempty"`
	DefaultDataType                      *DataType                 `json:"controllogix_ethernet.DEVICE_DEFAULT_DATA_TYPE,omitempty"`
	PerformanceStatistics                *bool                     `json:"controllogix_ethernet.DEVICE_ENABLE_PERFORMANCE_STATISTICS,omitempty"`
	DatabaseImportMethod                 *ImportMethod             `json:"controllogix_ethernet.DEVICE_DATABASE_IMPORT_METHOD,omitempty"`
	TagImportFile                        *string                   `json:"controllogix_ethernet.DEVICE_TAG_IMPORT_FILE,omitempty"`
	TagDescriptions                      *bool                     `json:"controllogix_ethernet.DEVICE_DISPLAY_DESCRIPTIONS,omitempty"`
	LimitNameLength                      *bool                     `json:"controllogix_ethernet.DEVICE_LIMIT_TAG_NAMES,omitempty"`
	TagHierarchy                         *TagHierarchy             `json:"controllogix_ethernet.DEVICE_TAG_HIERARCHY,omitempty"`
	ImposeArrayLimit                     *bool                     `json:"controllogix_ethernet.DEVICE_IMPOSE_ARRAY_ELEMENT_COUNT_LIMIT,omitempty"`
	ArrayCountUpperLimit                 *int                      `json:"controllogix_ethernet.DEVICE_ARRAY_ELEMENT_COUNT_LIMIT,omitempty"`
	RemoteTCPIPPort                      *int                      `json:"controllogix_ethernet.DEVICE_CL_ENET_PORT_NUMBER,omitempty"`
	RequestSize                          *RequestSize              `json:"controllogix_ethernet.DEVICE_REQUEST_SIZE,omitempty"`
	AllowFunctionFileBlockWrites         *bool                     `json:"controllogix_ethernet.DEVICE_PERFORM_BLOCK_WRITES,omitempty"`
}

// OPCUAClientDevice represents an OPC UA client device.
type OPCUAClientDevice struct {
	Name                       string             `json:"common.ALLTYPES_NAME"`
	Description                string             `json:"common.ALLTYPES_DESCRIPTION"`
	UniqueID                   int64              `json:"servermain.DEVICE_UNIQUE_ID"`
	ProjectID                  int64              `json:"PROJECT_ID"`
	ChannelAssignment          string             `json:"servermain.DEVICE_CHANNEL_ASSIGNMENT"`
	Driver                     string             `json:"servermain.MULTIPLE_TYPES_DEVICE_DRIVER"`
	Model                      OPCUAClientModel   `json:"servermain.DEVICE_MODEL"`
	DataCollection             bool               `json:"servermain.DEVICE_DATA_COLLECTION"`
	Simulated                  bool               `json:"servermain.DEVICE_SIMULATED"`
	ScanMode                   ScanMode           `json:"servermain.DEVICE_SCAN_MODE"`
	ScanRate                   int                `json:"servermain.DEVICE_SCAN_MODE_RATE_MS"`
	InitialUpdatesFromCache    bool               `json:"servermain.DEVICE_SCAN_MODE_PROVIDE_INITIAL_UPDATES_FROM_CACHE"`
	PublishingInterval         int                `json:"opcuaclient.DEVICE_SUBSCRIPTION_PUBLISHING_INTERVAL_MILLISECONDS"`
	MaxNotificationsPerPublish int                `json:"opcuaclient.DEVICE_SUBSCRIPTION_MAX_NOTIFICATIONS_PER_PUBLISH"`
	UpdateMode                 UpdateMode         `json:"opcuaclient.DEVICE_SUBSCRIPTION_UPDATE_MODE"`
	RegisteredReadWrite        bool               `json:"opcuaclient.DEVICE_SUBSCRIPTION_REGISTERED_READWRITE"`
	MaxItemsPerRead            int                `json:"opcuaclient.DEVICE_MAX_ITEMS_PER_READ"`
	MaxItemsPerWrite           int                `json:"opcuaclient.DEVICE_MAX_ITEMS_PER_WRITE"`
	ReadTimeout                int                `json:"opcuaclient.DEVICE_READ_TIMEOUT_MS"`
	WriteTimeout               int                `json:"opcuaclient.DEVICE_WRITE_TIMEOUT_MS"`
	ReadAfterWrite             bool               `json:"opcuaclient.DEVICE_READ_AFTER_WRITE"`
	LifetimeCount              int                `json:"opcuaclient.DEVICE_CONNECTION_LIFETIME_COUNT"`
	KeepAliveCount             int                `json:"opcuaclient.DEVICE_CONNECTION_MAX_KEEP_ALIVE"`
	ConnectionPriority         ConnectionPriority `json:"opcuaclient.DEVICE_CONNECTION_PRIORITY"`
	SampleInterval             int                `json:"opcuaclient.DEVICE_MONITORED_ITEMS_SAMPLE_INTERVAL_MILLISECONDS"`
	QueueSize                  int                `json:"opcuaclient.DEVICE_MONITORED_ITEMS_QUEUE_SIZE"`
	DiscardOldest              bool               `json:"opcuaclient.DEVICE_MONITORED_ITEMS_DISCARD_OLDEST"`
	DeadbandType               DeadbandType       `json:"opcuaclient.DEVICE_MONITORED_ITEMS_DEADBAND_TYPE"`
	DeadbandValue              int                `json:"opcuaclient.DEVICE_MONITORED_ITEMS_DEADBAND_VALUE"`
	SelectImportItems          []int              `json:"opcuaclient.DEVICE_MONITORED_ITEMS_SELECT_IMPORT"`
}

// OPCUAClientDeviceOptions represents all OPC UA client device options.
type OPCUAClientDeviceOptions struct {
	Name                       string              `json:"common.ALLTYPES_NAME,omitempty"`
	Description                string              `json:"common.ALLTYPES_DESCRIPTION,omitempty"`
	UniqueID                   int64               `json:"servermain.DEVICE_UNIQUE_ID,omitempty"`
	Model                      OPCUAClientModel    `json:"servermain.DEVICE_MODEL,omitempty"`
	DataCollection             *bool               `json:"servermain.DEVICE_DATA_COLLECTION,omitempty"`
	Simulated                  *bool               `json:"servermain.DEVICE_SIMULATED,omitempty"`
	ScanMode                   *ScanMode           `json:"servermain.DEVICE_SCAN_MODE,omitempty"`
	ScanRate                   *int                `json:"servermain.DEVICE_SCAN_MODE_RATE_MS,omitempty"`
	InitialUpdatesFromCache    *bool               `json:"servermain.DEVICE_SCAN_MODE_PROVIDE_INITIAL_UPDATES_FROM_CACHE,omitempty"`
	PublishingInterval         *int                `json:"opcuaclient.DEVICE_SUBSCRIPTION_PUBLISHING_INTERVAL_MILLISECONDS,omitempty"`
	MaxNotificationsPerPublish *int                `json:"opcuaclient.DEVICE_SUBSCRIPTION_MAX_NOTIFICATIONS_PER_PUBLISH,omitempty"`
	UpdateMode                 *UpdateMode         `json:"opcuaclient.DEVICE_SUBSCRIPTION_UPDATE_MODE,omitempty"`
	RegisteredReadWrite        *bool               `json:"opcuaclient.DEVICE_SUBSCRIPTION_REGISTERED_READWRITE,omitempty"`
	MaxItemsPerRead            *int                `json:"opcuaclient.DEVICE_MAX_ITEMS_PER_READ,omitempty"`
	MaxItemsPerWrite           *int                `json:"opcuaclient.DEVICE_MAX_ITEMS_PER_WRITE,omitempty"`
	ReadTimeout                *int                `json:"opcuaclient.DEVICE_READ_TIMEOUT_MS,omitempty"`
	WriteTimeout               *int                `json:"opcuaclient.DEVICE_WRITE_TIMEOUT_MS,omitempty"`
	ReadAfterWrite             *bool               `json:"opcuaclient.DEVICE_READ_AFTER_WRITE,omitempty"`
	LifetimeCount              *int                `json:"opcuaclient.DEVICE_CONNECTION_LIFETIME_COUNT,omitempty"`
	KeepAliveCount             *int                `json:"opcuaclient.DEVICE_CONNECTION_MAX_KEEP_ALIVE,omitempty"`
	ConnectionPriority         *ConnectionPriority `json:"opcuaclient.DEVICE_CONNECTION_PRIORITY,omitempty"`
	SampleInterval             *int                `json:"opcuaclient.DEVICE_MONITORED_ITEMS_SAMPLE_INTERVAL_MILLISECONDS,omitempty"`
	QueueSize                  *int                `json:"opcuaclient.DEVICE_MONITORED_ITEMS_QUEUE_SIZE,omitempty"`
	DiscardOldest              *bool               `json:"opcuaclient.DEVICE_MONITORED_ITEMS_DISCARD_OLDEST,omitempty"`
	DeadbandType               *DeadbandType       `json:"opcuaclient.DEVICE_MONITORED_ITEMS_DEADBAND_TYPE,omitempty"`
	DeadbandValue              *int                `json:"opcuaclient.DEVICE_MONITORED_ITEMS_DEADBAND_VALUE,omitempty"`
	SelectImportItems          []int               `json:"opcuaclient.DEVICE_MONITORED_ITEMS_SELECT_IMPORT,omitempty"`
}

// SiemensS5AS511Device represents a Siemens S5 (AS511) device.
type SiemensS5AS511Device struct {
	Name                       string              `json:"common.ALLTYPES_NAME"`
	Description                string              `json:"common.ALLTYPES_DESCRIPTION"`
	UniqueID                   int64               `json:"servermain.DEVICE_UNIQUE_ID"`
	ProjectID                  int64               `json:"PROJECT_ID"`
	ChannelAssignment          string              `json:"servermain.DEVICE_CHANNEL_ASSIGNMENT"`
	Driver                     string              `json:"servermain.MULTIPLE_TYPES_DEVICE_DRIVER"`
	Model                      SiemensS5AS511Model `json:"servermain.DEVICE_MODEL"`
	IDFormat                   IDFormat            `json:"servermain.DEVICE_ID_FORMAT"`
	IDString                   string              `json:"servermain.DEVICE_ID_STRING"`
	IDHexadecimal              int                 `json:"servermain.DEVICE_ID_HEXADECIMAL"`
	IDDecimal                  int                 `json:"servermain.DEVICE_ID_DECIMAL"`
	IDOctal                    int                 `json:"servermain.DEVICE_ID_OCTAL"`
	DataCollection             bool                `json:"servermain.DEVICE_DATA_COLLECTION"`
	Simulated                  bool                `json:"servermain.DEVICE_SIMULATED"`
	ScanMode                   ScanMode            `json:"servermain.DEVICE_SCAN_MODE"`
	ScanRate                   int                 `json:"servermain.DEVICE_SCAN_MODE_RATE_MS"`
	InitialUpdatesFromCache    bool                `json:"servermain.DEVICE_SCAN_MODE_PROVIDE_INITIAL_UPDATES_FROM_CACHE"`
	IPAddress                  string              `json:"servermain.DEVICE_ETHERNET_COMMUNICATIONS_IP"`
	Port                       int                 `json:"servermain.DEVICE_ETHERNET_COMMUNICATIONS_PORT"`
	Protocol                   Protocol            `json:"servermain.DEVICE_ETHERNET_COMMUNICATIONS_PROTOCOL"`
	ConnectionTimeout          int                 `json:"servermain.DEVICE_CONNECTION_TIMEOUT_SECONDS"`
	RequestTimeout             int                 `json:"servermain.DEVICE_REQUEST_TIMEOUT_MILLISECONDS"`
	AttemptsBeforeTimeout      int                 `json:"servermain.DEVICE_RETRY_ATTEMPTS"`
	DemoteOnFailure            bool                `json:"servermain.DEVICE_AUTO_DEMOTION_ENABLE_ON_COMMUNICATIONS_FAILURES"`
	TimeoutToDemote            int                 `json:"servermain.DEVICE_AUTO_DEMOTION_DEMOTE_AFTER_SUCCESSIVE_TIMEOUTS"`
	DemotionPeriod             int                 `json:"servermain.DEVICE_AUTO_DEMOTION_PERIOD_MS"`
	DiscardRequestsWhenDemoted bool                `json:"servermain.DEVICE_AUTO_DEMOTION_DISCARD_WRITES"`
}

// SiemensS5AS511DeviceOptions represents all Siemens S5 (AS511) device options.
type SiemensS5AS511DeviceOptions struct {
	Name                       string              `json:"common.ALLTYPES_NAME,omitempty"`
	Description                string              `json:"common.ALLTYPES_DESCRIPTION,omitempty"`
	UniqueID                   int64               `json:"servermain.DEVICE_UNIQUE_ID,omitempty"`
	Driver                     string              `json:"servermain.MULTIPLE_TYPES_DEVICE_DRIVER"`
	Model                      SiemensS5AS511Model `json:"servermain.DEVICE_MODEL,omitempty"`
	IDFormat                   IDFormat            `json:"servermain.DEVICE_ID_FORMAT,omitempty"`
	IDString                   string              `json:"servermain.DEVICE_ID_STRING,omitempty"`
	IDHexadecimal              int                 `json:"servermain.DEVICE_ID_HEXADECIMAL,omitempty"`
	IDDecimal                  int                 `json:"servermain.DEVICE_ID_DECIMAL,omitempty"`
	IDOctal                    int                 `json:"servermain.DEVICE_ID_OCTAL,omitempty"`
	DataCollection             *bool               `json:"servermain.DEVICE_DATA_COLLECTION,omitempty"`
	Simulated                  *bool               `json:"servermain.DEVICE_SIMULATED,omitempty"`
	ScanMode                   *ScanMode           `json:"servermain.DEVICE_SCAN_MODE,omitempty"`
	ScanRate                   *int                `json:"servermain.DEVICE_SCAN_MODE_RATE_MS,omitempty"`
	InitialUpdatesFromCache    *bool               `json:"servermain.DEVICE_SCAN_MODE_PROVIDE_INITIAL_UPDATES_FROM_CACHE,omitempty"`
	IPAddress                  *string             `json:"servermain.DEVICE_ETHERNET_COMMUNICATIONS_IP,omitempty"`
	Port                       *int                `json:"servermain.DEVICE_ETHERNET_COMMUNICATIONS_PORT,omitempty"`
	Protocol                   *Protocol           `json:"servermain.DEVICE_ETHERNET_COMMUNICATIONS_PROTOCOL,omitempty"`
	ConnectionTimeout          *int                `json:"servermain.DEVICE_CONNECTION_TIMEOUT_SECONDS,omitempty"`
	RequestTimeout             *int                `json:"servermain.DEVICE_REQUEST_TIMEOUT_MILLISECONDS,omitempty"`
	AttemptsBeforeTimeout      *int                `json:"servermain.DEVICE_RETRY_ATTEMPTS,omitempty"`
	DemoteOnFailure            *bool               `json:"servermain.DEVICE_AUTO_DEMOTION_ENABLE_ON_COMMUNICATIONS_FAILURES,omitempty"`
	TimeoutToDemote            *int                `json:"servermain.DEVICE_AUTO_DEMOTION_DEMOTE_AFTER_SUCCESSIVE_TIMEOUTS,omitempty"`
	DemotionPeriod             *int                `json:"servermain.DEVICE_AUTO_DEMOTION_PERIOD_MS,omitempty"`
	DiscardRequestsWhenDemoted *bool               `json:"servermain.DEVICE_AUTO_DEMOTION_DISCARD_WRITES,omitempty"`
}

// SiemensTCPIPEthernetDevice represents a Siemens TCP/IP Ethernet device.
type SiemensTCPIPEthernetDevice struct {
	Name                                 string                    `json:"common.ALLTYPES_NAME"`
	Description                          string                    `json:"common.ALLTYPES_DESCRIPTION"`
	UniqueID                             int64                     `json:"servermain.DEVICE_UNIQUE_ID"`
	ProjectID                            int64                     `json:"PROJECT_ID"`
	ChannelAssignment                    string                    `json:"servermain.DEVICE_CHANNEL_ASSIGNMENT"`
	Driver                               string                    `json:"servermain.MULTIPLE_TYPES_DEVICE_DRIVER"`
	Model                                SiemensTCPIPEthernetModel `json:"servermain.DEVICE_MODEL"`
	IDFormat                             IDFormat                  `json:"servermain.DEVICE_ID_FORMAT"`
	IDString                             string                    `json:"servermain.DEVICE_ID_STRING"`
	IDHexadecimal                        int                       `json:"servermain.DEVICE_ID_HEXADECIMAL"`
	IDDecimal                            int                       `json:"servermain.DEVICE_ID_DECIMAL"`
	IDOctal                              int                       `json:"servermain.DEVICE_ID_OCTAL"`
	DataCollection                       bool                      `json:"servermain.DEVICE_DATA_COLLECTION"`
	Simulated                            bool                      `json:"servermain.DEVICE_SIMULATED"`
	ScanMode                             ScanMode                  `json:"servermain.DEVICE_SCAN_MODE"`
	ScanRate                             int                       `json:"servermain.DEVICE_SCAN_MODE_RATE_MS"`
	InitialUpdatesFromCache              bool                      `json:"servermain.DEVICE_SCAN_MODE_PROVIDE_INITIAL_UPDATES_FROM_CACHE"`
	ConnectionTimeout                    int                       `json:"servermain.DEVICE_CONNECTION_TIMEOUT_SECONDS"`
	RequestTimeout                       int                       `json:"servermain.DEVICE_REQUEST_TIMEOUT_MILLISECONDS"`
	AttemptsBeforeTimeout                int                       `json:"servermain.DEVICE_RETRY_ATTEMPTS"`
	InterRequestDelay                    int                       `json:"servermain.DEVICE_INTER_REQUEST_DELAY_MILLISECONDS"`
	DemoteOnFailure                      bool                      `json:"servermain.DEVICE_AUTO_DEMOTION_ENABLE_ON_COMMUNICATIONS_FAILURES"`
	TimeoutToDemote                      int                       `json:"servermain.DEVICE_AUTO_DEMOTION_DEMOTE_AFTER_SUCCESSIVE_TIMEOUTS"`
	DemotionPeriod                       int                       `json:"servermain.DEVICE_AUTO_DEMOTION_PERIOD_MS"`
	DiscardRequestsWhenDemoted           bool                      `json:"servermain.DEVICE_AUTO_DEMOTION_DISCARD_WRITES"`
	OnDeviceStartup                      OnDeviceStartup           `json:"servermain.DEVICE_TAG_GENERATION_ON_STARTUP"`
	OnDuplicateTag                       OnDuplicateTag            `json:"servermain.DEVICE_TAG_GENERATION_DUPLICATE_HANDLING"`
	ParentGroup                          string                    `json:"servermain.DEVICE_TAG_GENERATION_GROUP"`
	AllowAutomaticallyGeneratedSubgroups bool                      `json:"servermain.DEVICE_TAG_GENERATION_ALLOW_SUB_GROUPS"`
	PortNumber                           int                       `json:"siemens_tcpip_ethernet.DEVICE_COMMUNICATIONS_PORT_NUMBER"`
	MPIID                                int                       `json:"siemens_tcpip_ethernet.DEVICE_COMMUNICATIONS_MPI_ID"`
	MaxPDUSize                           MaxPDUSize                `json:"siemens_tcpip_ethernet.DEVICE_S7_COMMUNICATIONS_MAX_PDU"`
	LocalTSAP                            int                       `json:"siemens_tcpip_ethernet.DEVICE_S7_COMMUNICATIONS_200_LOCAL_TSAP"`
	RemoteTSAP                           int                       `json:"siemens_tcpip_ethernet.DEVICE_S7_COMMUNICATIONS_200_REMOTE_TSAP"`
	LinkType                             LinkType                  `json:"siemens_tcpip_ethernet.DEVICE_S7_COMMUNICATIONS_300_400_1200_1500_LINK_TYPE"`
	CPURack                              int                       `json:"siemens_tcpip_ethernet.DEVICE_S7_COMMUNICATIONS_CPU_RACK"`
	CPUSlot                              int                       `json:"siemens_tcpip_ethernet.DEVICE_S7_COMMUNICATIONS_CPU_SLOT"`
	ByteOrder                            ByteOrder                 `json:"siemens_tcpip_ethernet.DEVICE_ADDRESSING_BYTE_ORDER"`
	TagImportType                        int                       `json:"siemens_tcpip_ethernet.DEVICE_TAG_IMPORT_TYPE"`
	CodePage                             int64                     `json:"siemens_tcpip_ethernet.DEVICE_TAG_IMPORT_CODE_PAGE"`
	Step7Project                         string                    `json:"siemens_tcpip_ethernet.DEVICE_TAG_IMPORT_STEP_7_PROJECT_FILE"`
	ProgramPath                          string                    `json:"siemens_tcpip_ethernet.DEVICE_TAG_IMPORT_PROGRAM_PATH"`
	TIAPortalExporterFile                string                    `json:"siemens_tcpip_ethernet.DEVICE_TAG_IMPORT_TIA_EXPORT_FILE"`
}

// SiemensTCPIPEthernetDeviceOptions represents all Siemens TCP/IP Ethernet device options.
type SiemensTCPIPEthernetDeviceOptions struct {
	Name                                 string                    `json:"common.ALLTYPES_NAME,omitempty"`
	Description                          string                    `json:"common.ALLTYPES_DESCRIPTION,omitempty"`
	UniqueID                             int64                     `json:"servermain.DEVICE_UNIQUE_ID,omitempty"`
	Driver                               string                    `json:"servermain.MULTIPLE_TYPES_DEVICE_DRIVER"`
	Model                                SiemensTCPIPEthernetModel `json:"servermain.DEVICE_MODEL,omitempty"`
	IDFormat                             IDFormat                  `json:"servermain.DEVICE_ID_FORMAT,omitempty"`
	IDString                             string                    `json:"servermain.DEVICE_ID_STRING,omitempty"`
	IDHexadecimal                        int                       `json:"servermain.DEVICE_ID_HEXADECIMAL,omitempty"`
	IDDecimal                            int                       `json:"servermain.DEVICE_ID_DECIMAL,omitempty"`
	IDOctal                              int                       `json:"servermain.DEVICE_ID_OCTAL,omitempty"`
	DataCollection                       *bool                     `json:"servermain.DEVICE_DATA_COLLECTION,omitempty"`
	Simulated                            *bool                     `json:"servermain.DEVICE_SIMULATED,omitempty"`
	ScanMode                             *ScanMode                 `json:"servermain.DEVICE_SCAN_MODE,omitempty"`
	ScanRate                             *int                      `json:"servermain.DEVICE_SCAN_MODE_RATE_MS,omitempty"`
	InitialUpdatesFromCache              *bool                     `json:"servermain.DEVICE_SCAN_MODE_PROVIDE_INITIAL_UPDATES_FROM_CACHE,omitempty"`
	ConnectionTimeout                    *int                      `json:"servermain.DEVICE_CONNECTION_TIMEOUT_SECONDS,omitempty"`
	RequestTimeout                       *int                      `json:"servermain.DEVICE_REQUEST_TIMEOUT_MILLISECONDS,omitempty"`
	AttemptsBeforeTimeout                *int                      `json:"servermain.DEVICE_RETRY_ATTEMPTS,omitempty"`
	InterRequestDelay                    *int                      `json:"servermain.DEVICE_INTER_REQUEST_DELAY_MILLISECONDS,omitempty"`
	DemoteOnFailure                      *bool                     `json:"servermain.DEVICE_AUTO_DEMOTION_ENABLE_ON_COMMUNICATIONS_FAILURES,omitempty"`
	TimeoutToDemote                      *int                      `json:"servermain.DEVICE_AUTO_DEMOTION_DEMOTE_AFTER_SUCCESSIVE_TIMEOUTS,omitempty"`
	DemotionPeriod                       *int                      `json:"servermain.DEVICE_AUTO_DEMOTION_PERIOD_MS,omitempty"`
	DiscardRequestsWhenDemoted           *bool                     `json:"servermain.DEVICE_AUTO_DEMOTION_DISCARD_WRITES,omitempty"`
	OnDeviceStartup                      *OnDeviceStartup          `json:"servermain.DEVICE_TAG_GENERATION_ON_STARTUP,omitempty"`
	OnDuplicateTag                       *OnDuplicateTag           `json:"servermain.DEVICE_TAG_GENERATION_DUPLICATE_HANDLING,omitempty"`
	ParentGroup                          *string                   `json:"servermain.DEVICE_TAG_GENERATION_GROUP,omitempty"`
	AllowAutomaticallyGeneratedSubgroups *bool                     `json:"servermain.DEVICE_TAG_GENERATION_ALLOW_SUB_GROUPS,omitempty"`
	PortNumber                           *int                      `json:"siemens_tcpip_ethernet.DEVICE_COMMUNICATIONS_PORT_NUMBER,omitempty"`
	MPIID                                *int                      `json:"siemens_tcpip_ethernet.DEVICE_COMMUNICATIONS_MPI_ID,omitempty"`
	MaxPDUSize                           *MaxPDUSize               `json:"siemens_tcpip_ethernet.DEVICE_S7_COMMUNICATIONS_MAX_PDU,omitempty"`
	LocalTSAP                            *int                      `json:"siemens_tcpip_ethernet.DEVICE_S7_COMMUNICATIONS_200_LOCAL_TSAP,omitempty"`
	RemoteTSAP                           *int                      `json:"siemens_tcpip_ethernet.DEVICE_S7_COMMUNICATIONS_200_REMOTE_TSAP,omitempty"`
	LinkType                             *LinkType                 `json:"siemens_tcpip_ethernet.DEVICE_S7_COMMUNICATIONS_300_400_1200_1500_LINK_TYPE,omitempty"`
	CPURack                              *int                      `json:"siemens_tcpip_ethernet.DEVICE_S7_COMMUNICATIONS_CPU_RACK,omitempty"`
	CPUSlot                              *int                      `json:"siemens_tcpip_ethernet.DEVICE_S7_COMMUNICATIONS_CPU_SLOT,omitempty"`
	ByteOrder                            *ByteOrder                `json:"siemens_tcpip_ethernet.DEVICE_ADDRESSING_BYTE_ORDER,omitempty"`
	TagImportType                        *int                      `json:"siemens_tcpip_ethernet.DEVICE_TAG_IMPORT_TYPE,omitempty"`
	CodePage                             *int64                    `json:"siemens_tcpip_ethernet.DEVICE_TAG_IMPORT_CODE_PAGE,omitempty"`
	Step7Project                         *string                   `json:"siemens_tcpip_ethernet.DEVICE_TAG_IMPORT_STEP_7_PROJECT_FILE,omitempty"`
	ProgramPath                          *string                   `json:"siemens_tcpip_ethernet.DEVICE_TAG_IMPORT_PROGRAM_PATH,omitempty"`
	TIAPortalExporterFile                *string                   `json:"siemens_tcpip_ethernet.DEVICE_TAG_IMPORT_TIA_EXPORT_FILE,omitempty"`
}

// ListDevices gets a list of devices.
func (s *DeviceService) ListDevices(channel string) ([]*Device, error) {
	u := fmt.Sprintf("channels/%s/devices", url.PathEscape(channel))

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var devices []*Device
	if err = s.client.Do(req, &devices); err != nil {
		return nil, err
	}

	return devices, nil
}

// CreateControlLogixEthernetDevice creates a new ControlLogix Ethernet device.
func (s *DeviceService) CreateControlLogixEthernetDevice(channel string, options *ControlLogixEthernetDeviceOptions) error {
	return s.createDevice(channel, options)
}

// CreateOPCUAClientDevice creates a new OPC UA client device.
func (s *DeviceService) CreateOPCUAClientDevice(channel string, options *OPCUAClientDeviceOptions) error {
	return s.createDevice(channel, options)
}

// CreateSiemensS5AS511Device creates a new Siemens S5 (AS511) device.
func (s *DeviceService) CreateSiemensS5AS511Device(channel string, options *SiemensS5AS511DeviceOptions) error {
	return s.createDevice(channel, options)
}

// CreateSiemensTCPIPEthernetDevice creates a new Siemens TCP/IP Ethernet device.
func (s *DeviceService) CreateSiemensTCPIPEthernetDevice(channel string, options *SiemensTCPIPEthernetDeviceOptions) error {
	return s.createDevice(channel, options)
}

func (s *DeviceService) createDevice(channel string, v interface{}) error {
	u := fmt.Sprintf("channels/%s/devices", url.PathEscape(channel))
	req, err := s.client.NewRequest("POST", u, v)
	if err != nil {
		return err
	}
	return s.client.Do(req, nil)
}

// GetControlLogixEthernetDevice gets a ControlLogix Ethernet device.
func (s *DeviceService) GetControlLogixEthernetDevice(channel, name string) (*ControlLogixEthernetDevice, error) {
	var device *ControlLogixEthernetDevice
	v, err := s.getDevice(channel, name, device)
	if err != nil {
		return nil, err
	}
	return v.(*ControlLogixEthernetDevice), nil
}

// GetOPCUAClientDevice gets an OPC UA client device.
func (s *DeviceService) GetOPCUAClientDevice(channel, name string) (*OPCUAClientDevice, error) {
	var device *OPCUAClientDevice
	v, err := s.getDevice(channel, name, device)
	if err != nil {
		return nil, err
	}
	return v.(*OPCUAClientDevice), nil
}

// GetSiemensS5AS511Device gets an Siemens S5 (AS511) device.
func (s *DeviceService) GetSiemensS5AS511Device(channel, name string) (*SiemensS5AS511Device, error) {
	var device *SiemensS5AS511Device
	v, err := s.getDevice(channel, name, device)
	if err != nil {
		return nil, err
	}
	return v.(*SiemensS5AS511Device), nil
}

// GetSiemensTCPIPEthernetDevice gets an Siemens TCP/IP Ethernet device.
func (s *DeviceService) GetSiemensTCPIPEthernetDevice(channel, name string) (*SiemensTCPIPEthernetDevice, error) {
	var device *SiemensTCPIPEthernetDevice
	v, err := s.getDevice(channel, name, device)
	if err != nil {
		return nil, err
	}
	return v.(*SiemensTCPIPEthernetDevice), nil
}

// getDevice gets a specific device.
func (s *DeviceService) getDevice(channel, name string, v interface{}) (interface{}, error) {
	u := fmt.Sprintf("channels/%s/devices/%s", url.PathEscape(channel), url.PathEscape(name))

	req, err := s.client.NewRequest("GET", u, v)
	if err != nil {
		return nil, err
	}

	var c *Device
	if err = s.client.Do(req, c); err != nil {
		return nil, err
	}

	return c, nil
}

// UpdateControlLogixEthernetDevice updates an existing ControlLogix Ethernet device.
func (s *DeviceService) UpdateControlLogixEthernetDevice(channel string, options *ControlLogixEthernetDevice) error {
	return s.createDevice(channel, options)
}

// UpdateOPCUAClientDevice updates an existing OPC UA client device.
func (s *DeviceService) UpdateOPCUAClientDevice(channel, name string, options *OPCUAClientDeviceOptions) error {
	return s.updateDevice(channel, name, options)
}

// UpdateSiemensS5AS511Device updates an existing Siemens S5 (AS511) device.
func (s *DeviceService) UpdateSiemensS5AS511Device(channel, name string, options *SiemensS5AS511DeviceOptions) error {
	return s.updateDevice(channel, name, options)
}

// UpdateSiemensTCPIPEthernetDevice updates an existing Siemens TCP/IP Ethernet device.
func (s *DeviceService) UpdateSiemensTCPIPEthernetDevice(channel, name string, options *SiemensTCPIPEthernetDeviceOptions) error {
	return s.updateDevice(channel, name, options)
}

func (s *DeviceService) updateDevice(channel, name string, v interface{}) error {
	u := fmt.Sprintf("channels/%s/devices/%s", url.PathEscape(channel), url.PathEscape(name))
	req, err := s.client.NewRequest("PUT", u, v)
	if err != nil {
		return err
	}
	return s.client.Do(req, nil)
}

// DeleteDevice deletes a device.
func (s *DeviceService) DeleteDevice(channel, name string) error {
	u := fmt.Sprintf("channels/%s/devices/%s", url.PathEscape(channel), url.PathEscape(name))
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return err
	}
	return s.client.Do(req, nil)
}
