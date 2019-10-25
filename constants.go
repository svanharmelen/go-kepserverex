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

// ArrayBlockSize represents an array block size.
type ArrayBlockSize int

// List of available array block sizes.
const (
	ArrayBlockSize_30   ArrayBlockSize = 30
	ArrayBlockSize_60   ArrayBlockSize = 60
	ArrayBlockSize_120  ArrayBlockSize = 120
	ArrayBlockSize_240  ArrayBlockSize = 240
	ArrayBlockSize_480  ArrayBlockSize = 480
	ArrayBlockSize_960  ArrayBlockSize = 960
	ArrayBlockSize_1920 ArrayBlockSize = 1920
	ArrayBlockSize_3840 ArrayBlockSize = 3840
)

// BautRate represents a baut rate.
type BautRate int

// List of available baut rates.
const (
	BautRate_300    BautRate = 300
	BautRate_600    BautRate = 600
	BautRate_1200   BautRate = 1200
	BautRate_2400   BautRate = 2400
	BautRate_4800   BautRate = 4800
	BautRate_9600   BautRate = 9600
	BautRate_14400  BautRate = 14400
	BautRate_19200  BautRate = 19200
	BautRate_28800  BautRate = 28800
	BautRate_38400  BautRate = 38400
	BautRate_56000  BautRate = 56000
	BautRate_57600  BautRate = 57600
	BautRate_115200 BautRate = 115200
	BautRate_128000 BautRate = 128000
	BautRate_256000 BautRate = 256000
)

// ByteOrder represents a byte order.
type ByteOrder int

// List of available byte orders.
const (
	BigEndian ByteOrder = iota
	LittleEndian
)

// DataBits represents a data bit size.
type DataBits int

// List of available data bit sizes.
const (
	DataBits_5 DataBits = 5
	DataBits_6 DataBits = 6
	DataBits_7 DataBits = 7
	DataBits_8 DataBits = 8
)

// DataType represents a data type.
type DataType int

// List of available data types.
const (
	DataType_Default DataType = iota - 1
	DataType_String
	DataType_Boolean
	DataType_Char
	DataType_Byte
	DataType_Short
	DataType_Word
	DataType_Long
	DataType_DWord
	DataType_Float
	DataType_Double
	DataType_BCD
	DataType_LBCD
	DataType_Date
	DataType_LLong
	DataType_Qword
	DataType_StringArray
	DataType_BooleanArray
	DataType_CharArray
	DataType_ByteArray
	DataType_ShortArray
	DataType_WordArray
	DataType_LongArray
	DataType_DWordArray
	DataType_FloatArray
	DataType_DoubleArray
	DataType_BCDArray
	DataType_LBCDArray
	DataType_DateArray
	DataType_LLongArray
	DataType_QwordArray
)

// ClientAccess represents a client access type.
type ClientAccess int

// List of available client access types.
const (
	ClientAccess_ReadOnly ClientAccess = iota
	ClientAccess_ReadWrite
)

// ConnectionPriority represents a connection priority.
type ConnectionPriority int

// List of available connection priorities.
const (
	ConnectionPriority_Lowest  ConnectionPriority = 0
	ConnectionPriority_Low     ConnectionPriority = 64
	ConnectionPriority_Medium  ConnectionPriority = 128
	ConnectionPriority_High    ConnectionPriority = 192
	ConnectionPriority_Highest ConnectionPriority = 255
)

// ControlLogixEthernetModel represents a ControlLogix Ethernet model.
type ControlLogixEthernetModel int

// List of available ControlLogix Ethernet models.
const (
	ControlLogix_5500 ControlLogixEthernetModel = iota
	CompactLogix_5300
	FlexLogix_5400
	SoftLogix_5800
	DH_Gateway_PLC_5
	DH_Gateway_SLC_5_04
	ControlNetGateway_PLC_5C
	EIP_Gateway_MicroLogix
	EIP_Gateway_SLC_Fixed
	EIP_Gateway_SLC_Modular
	EIP_Gateway_PLC_5
	Serial_Gateway_ControlLogix
	Serial_Gateway_CompactLogix
	Serial_Gateway_FlexLogix
	Serial_Gateway_SoftLogix
	ENI_ControlLogix_5500
	ENI_CompactLogix_5300
	ENI_FlexLogix_5400
	ENI_SoftLogix_5800
	ENI_SLC_500_Fixed_IO
	ENI_SLC_500_Modular_IO
	ENI_PLC_5
	MicroLogix_1100
	MicroLogix_1400
)

// DeadbandType represents a deadband type.
type DeadbandType int

// List of available deadband types.
const (
	DeadbandType_None DeadbandType = iota
	DeadbandType_Percent
	DeadbandType_Absolute
)

// FloatingPointValues represents a floating point option.
type FloatingPointValues int

// List of available floating point options.
const (
	ReplaceWithZero FloatingPointValues = iota
	Unmodified
)

// FlowControl represents a flow control option.
type FlowControl int

// List of available flow control options.
const (
	FlowControl_None FlowControl = iota
	FlowControl_DTR
	FlowControl_RTS
	FlowControl_RTSDTR
	FlowControl_RTSAlways
	FlowControl_RTSManual
)

// IDFormat represents an ID format.
type IDFormat int

// List of available ID formats.
const (
	IDFormat_Octal IDFormat = iota
	IDFormat_Decimal
	IDFormat_Hex
)

// InactivityWatchdog represents an inactivity watchdog timeout
type InactivityWatchdog int

// List of available inactivity watchdog seconds.
const (
	InactivityWatchdog_8   InactivityWatchdog = 8
	InactivityWatchdog_16  InactivityWatchdog = 16
	InactivityWatchdog_32  InactivityWatchdog = 32
	InactivityWatchdog_64  InactivityWatchdog = 64
	InactivityWatchdog_128 InactivityWatchdog = 128
)

// ImportMethod represents a database import method.
type ImportMethod int

// List of available database import methods.
const (
	CreateFromDevice ImportMethod = iota
	CreateFromImportFile
)

// LinkType represents a link type.
type LinkType int

// List of available link types.
const (
	LinkType_PG LinkType = iota + 1
	LinkType_OP
	LinkType_PC
)

// MaxPDUSize represents a maximum PDU size mode.
type MaxPDUSize int

// List of available maximum PDU sizes.
const (
	MaxPDUSize_240 MaxPDUSize = 240
	MaxPDUSize_480 MaxPDUSize = 480
	MaxPDUSize_960 MaxPDUSize = 960
)

// MessageMode represents a message mode.
type MessageMode int

// List of available message modes.
const (
	MessageMode_None MessageMode = iota + 1
	MessageMode_Sign
	MessageMode_SignAndEncrypt
)

// NetworkMode represents a network mode.
type NetworkMode int

// List of available network modes.
const (
	NetworkMode_LoadBalanced NetworkMode = iota
	NetworkMode_Priority
)

// OnDeviceStartup represents a on device startup mode.
type OnDeviceStartup int

// List of available on device startup modes.
const (
	DoNotGenerateOnStartup OnDeviceStartup = iota
	AlwaysGenerateOnStartup
	GenerateOnFirstStartup
)

// OnDuplicateTag represents a on duplicate tag mode.
type OnDuplicateTag int

// List of available on duplicate tag modes.
const (
	DeleteOnCreate OnDuplicateTag = iota
	OverwriteAsNecessary
	DoNotOverwrite
	DoNotOverwriteLogError
)

// OPCUAClientModel represents an OPC UA Client model.
type OPCUAClientModel int

// List of available OPC UA Client models.
const (
	OPCUA OPCUAClientModel = iota
)

// OptimizationMethod represents an optimization method.
type OptimizationMethod int

// List of available optimization methods.
const (
	WriteAllValuesForAllTags OptimizationMethod = iota
	WriteOnlyLatestValueForNonBooleanTags
	WriteOnlyLatestValueForAllTags
)

// Parity represents a parity option.
type Parity int

// List of available parity options.
const (
	Parity_None Parity = 78
	Parity_Odd  Parity = 79
	Parity_Even Parity = 69
)

// PhysicalMedium represents a physical medium.
type PhysicalMedium int

// List of available physical mediums.
const (
	PhysicalMedium_None PhysicalMedium = iota
	PhysicalMedium_COMPort
	PhysicalMedium_Modem
	PhysicalMedium_EthernetEncap
)

// Protocol represents a protocol.
type Protocol int

// List of available protocols.
const (
	UDP Protocol = iota
	TCPIP
)

// ProtocolMode represents a protocol mode.
type ProtocolMode int

// List of available protocol modes.
const (
	Symbolic ProtocolMode = iota
	LogicalNonBlocking
	LogicalBlocking
)

// ReadProcessing represents a read processing option.
type ReadProcessing int

// List of available read processing options.
const (
	ReadProcessing_Ignore ReadProcessing = iota
	ReadProcessing_Fail
)

// RequestSize represents a request size.
type RequestSize int

// List of available request sizes.
const (
	RequestSize_32  RequestSize = 32
	RequestSize_64  RequestSize = 64
	RequestSize_128 RequestSize = 128
	RequestSize_232 RequestSize = 232
)

// ScanMode represents a scan mode.
type ScanMode int

// List of available scan modes.
const (
	RespectClientSpecifiedScanRate ScanMode = iota
	RequestDataNoFasterThenScanRate
	RequestAllDataAtScanRate
	DoNotScanDemandPollOnly
	RespectTagSpecifiedScanRate
)

// ScaledDataType represents a scaled data type.
type ScaledDataType int

// List of available scaled data types.
const (
	ScaledDataType_Char ScaledDataType = iota + 2
	ScaledDataType_Byte
	ScaledDataType_Short
	ScaledDataType_Word
	ScaledDataType_Long
	ScaledDataType_DWord
	ScaledDataType_Float
	ScaledDataType_Double
)

// ScalingType represents a scaling type.
type ScalingType int

// List of available scaling types.
const (
	ScalingType_None ScalingType = iota
	ScalingType_Linear
	ScalingType_SquareRoot
)

// SecurityPolicy represents a security policy.
type SecurityPolicy int

// List of available security policies.
const (
	SecurityPolicy_None           SecurityPolicy = 0
	SecurityPolicy_Basic256SHA256 SecurityPolicy = 3
)

// SiemensS5AS511Model represents a Siemens S5 (AS511) model.
type SiemensS5AS511Model int

// List of available Siemens S5 (AS511) models.
const (
	S5_AS511_90U SiemensS5AS511Model = iota
	S5_AS511_95U
	S5_AS511_100U100
	S5_AS511_100U101
	S5_AS511_100U103
	S5_AS511_101U
	S5_AS511_115U941
	S5_AS511_115U942
	S5_AS511_115U943
	S5_AS511_115U944
	S5_AS511_115U945
	S5_AS511_135U921
	S5_AS511_135U922
	S5_AS511_135U928
	S5_AS511_155U946
	S5_AS511_155U947
)

// SiemensTCPIPEthernetModel represents a Siemens TCP/IP Ethernet model.
type SiemensTCPIPEthernetModel int

// List of available Siemens TCP/IP Ethernet models.
const (
	S7_200 SiemensTCPIPEthernetModel = iota
	S7_300
	S7_400
	S7_1200
	S7_1500
	Netlink_S7_300
	Netlink_S7_400
)

// StopBits represents a stop bit size.
type StopBits int

// List of available stop bit sizes.
const (
	StopBits_1 StopBits = 1
	StopBits_2 StopBits = 2
)

// TagHierarchy represents a tag hierarchy.
type TagHierarchy int

// List of available tag hierarchies.
const (
	Condensed TagHierarchy = iota
	Expanded
)

// UpdateMode represents an update mode.
type UpdateMode int

// List of available update modes.
const (
	UpdateMode_Exception UpdateMode = iota
	UpdateMode_Poll
)

// VirtualNetwork represents a virtual network.
type VirtualNetwork int

// List of available virtual networks.
const (
	VirtualNetwork_None VirtualNetwork = iota
	VirtualNetwork_1
	VirtualNetwork_2
	VirtualNetwork_3
	VirtualNetwork_4
	VirtualNetwork_5
	VirtualNetwork_6
	VirtualNetwork_7
	VirtualNetwork_8
	VirtualNetwork_9
	VirtualNetwork_10
	VirtualNetwork_11
	VirtualNetwork_12
	VirtualNetwork_13
	VirtualNetwork_14
	VirtualNetwork_15
	VirtualNetwork_16
	VirtualNetwork_17
	VirtualNetwork_18
	VirtualNetwork_19
	VirtualNetwork_20
	VirtualNetwork_21
	VirtualNetwork_22
	VirtualNetwork_23
	VirtualNetwork_24
	VirtualNetwork_25
	VirtualNetwork_26
	VirtualNetwork_27
	VirtualNetwork_28
	VirtualNetwork_29
	VirtualNetwork_30
	VirtualNetwork_31
	VirtualNetwork_32
	VirtualNetwork_33
	VirtualNetwork_34
	VirtualNetwork_35
	VirtualNetwork_36
	VirtualNetwork_37
	VirtualNetwork_38
	VirtualNetwork_39
	VirtualNetwork_40
	VirtualNetwork_41
	VirtualNetwork_42
	VirtualNetwork_43
	VirtualNetwork_44
	VirtualNetwork_45
	VirtualNetwork_46
	VirtualNetwork_47
	VirtualNetwork_48
	VirtualNetwork_49
	VirtualNetwork_50
	VirtualNetwork_51
	VirtualNetwork_52
	VirtualNetwork_53
	VirtualNetwork_54
	VirtualNetwork_55
	VirtualNetwork_56
	VirtualNetwork_57
	VirtualNetwork_58
	VirtualNetwork_59
	VirtualNetwork_60
	VirtualNetwork_61
	VirtualNetwork_62
	VirtualNetwork_63
	VirtualNetwork_64
	VirtualNetwork_65
	VirtualNetwork_66
	VirtualNetwork_67
	VirtualNetwork_68
	VirtualNetwork_69
	VirtualNetwork_70
	VirtualNetwork_71
	VirtualNetwork_72
	VirtualNetwork_73
	VirtualNetwork_74
	VirtualNetwork_75
	VirtualNetwork_76
	VirtualNetwork_77
	VirtualNetwork_78
	VirtualNetwork_79
	VirtualNetwork_80
	VirtualNetwork_81
	VirtualNetwork_82
	VirtualNetwork_83
	VirtualNetwork_84
	VirtualNetwork_85
	VirtualNetwork_86
	VirtualNetwork_87
	VirtualNetwork_88
	VirtualNetwork_89
	VirtualNetwork_90
	VirtualNetwork_91
	VirtualNetwork_92
	VirtualNetwork_93
	VirtualNetwork_94
	VirtualNetwork_95
	VirtualNetwork_96
	VirtualNetwork_97
	VirtualNetwork_98
	VirtualNetwork_99
	VirtualNetwork_100
	VirtualNetwork_101
	VirtualNetwork_102
	VirtualNetwork_103
	VirtualNetwork_104
	VirtualNetwork_105
	VirtualNetwork_106
	VirtualNetwork_107
	VirtualNetwork_108
	VirtualNetwork_109
	VirtualNetwork_110
	VirtualNetwork_111
	VirtualNetwork_112
	VirtualNetwork_113
	VirtualNetwork_114
	VirtualNetwork_115
	VirtualNetwork_116
	VirtualNetwork_117
	VirtualNetwork_118
	VirtualNetwork_119
	VirtualNetwork_120
	VirtualNetwork_121
	VirtualNetwork_122
	VirtualNetwork_123
	VirtualNetwork_124
	VirtualNetwork_125
	VirtualNetwork_126
	VirtualNetwork_127
	VirtualNetwork_128
	VirtualNetwork_129
	VirtualNetwork_130
	VirtualNetwork_131
	VirtualNetwork_132
	VirtualNetwork_133
	VirtualNetwork_134
	VirtualNetwork_135
	VirtualNetwork_136
	VirtualNetwork_137
	VirtualNetwork_138
	VirtualNetwork_139
	VirtualNetwork_140
	VirtualNetwork_141
	VirtualNetwork_142
	VirtualNetwork_143
	VirtualNetwork_144
	VirtualNetwork_145
	VirtualNetwork_146
	VirtualNetwork_147
	VirtualNetwork_148
	VirtualNetwork_149
	VirtualNetwork_150
	VirtualNetwork_151
	VirtualNetwork_152
	VirtualNetwork_153
	VirtualNetwork_154
	VirtualNetwork_155
	VirtualNetwork_156
	VirtualNetwork_157
	VirtualNetwork_158
	VirtualNetwork_159
	VirtualNetwork_160
	VirtualNetwork_161
	VirtualNetwork_162
	VirtualNetwork_163
	VirtualNetwork_164
	VirtualNetwork_165
	VirtualNetwork_166
	VirtualNetwork_167
	VirtualNetwork_168
	VirtualNetwork_169
	VirtualNetwork_170
	VirtualNetwork_171
	VirtualNetwork_172
	VirtualNetwork_173
	VirtualNetwork_174
	VirtualNetwork_175
	VirtualNetwork_176
	VirtualNetwork_177
	VirtualNetwork_178
	VirtualNetwork_179
	VirtualNetwork_180
	VirtualNetwork_181
	VirtualNetwork_182
	VirtualNetwork_183
	VirtualNetwork_184
	VirtualNetwork_185
	VirtualNetwork_186
	VirtualNetwork_187
	VirtualNetwork_188
	VirtualNetwork_189
	VirtualNetwork_190
	VirtualNetwork_191
	VirtualNetwork_192
	VirtualNetwork_193
	VirtualNetwork_194
	VirtualNetwork_195
	VirtualNetwork_196
	VirtualNetwork_197
	VirtualNetwork_198
	VirtualNetwork_199
	VirtualNetwork_200
	VirtualNetwork_201
	VirtualNetwork_202
	VirtualNetwork_203
	VirtualNetwork_204
	VirtualNetwork_205
	VirtualNetwork_206
	VirtualNetwork_207
	VirtualNetwork_208
	VirtualNetwork_209
	VirtualNetwork_210
	VirtualNetwork_211
	VirtualNetwork_212
	VirtualNetwork_213
	VirtualNetwork_214
	VirtualNetwork_215
	VirtualNetwork_216
	VirtualNetwork_217
	VirtualNetwork_218
	VirtualNetwork_219
	VirtualNetwork_220
	VirtualNetwork_221
	VirtualNetwork_222
	VirtualNetwork_223
	VirtualNetwork_224
	VirtualNetwork_225
	VirtualNetwork_226
	VirtualNetwork_227
	VirtualNetwork_228
	VirtualNetwork_229
	VirtualNetwork_230
	VirtualNetwork_231
	VirtualNetwork_232
	VirtualNetwork_233
	VirtualNetwork_234
	VirtualNetwork_235
	VirtualNetwork_236
	VirtualNetwork_237
	VirtualNetwork_238
	VirtualNetwork_239
	VirtualNetwork_240
	VirtualNetwork_241
	VirtualNetwork_242
	VirtualNetwork_243
	VirtualNetwork_244
	VirtualNetwork_245
	VirtualNetwork_246
	VirtualNetwork_247
	VirtualNetwork_248
	VirtualNetwork_249
	VirtualNetwork_250
	VirtualNetwork_251
	VirtualNetwork_252
	VirtualNetwork_253
	VirtualNetwork_254
	VirtualNetwork_255
	VirtualNetwork_256
	VirtualNetwork_257
	VirtualNetwork_258
	VirtualNetwork_259
	VirtualNetwork_260
	VirtualNetwork_261
	VirtualNetwork_262
	VirtualNetwork_263
	VirtualNetwork_264
	VirtualNetwork_265
	VirtualNetwork_266
	VirtualNetwork_267
	VirtualNetwork_268
	VirtualNetwork_269
	VirtualNetwork_270
	VirtualNetwork_271
	VirtualNetwork_272
	VirtualNetwork_273
	VirtualNetwork_274
	VirtualNetwork_275
	VirtualNetwork_276
	VirtualNetwork_277
	VirtualNetwork_278
	VirtualNetwork_279
	VirtualNetwork_280
	VirtualNetwork_281
	VirtualNetwork_282
	VirtualNetwork_283
	VirtualNetwork_284
	VirtualNetwork_285
	VirtualNetwork_286
	VirtualNetwork_287
	VirtualNetwork_288
	VirtualNetwork_289
	VirtualNetwork_290
	VirtualNetwork_291
	VirtualNetwork_292
	VirtualNetwork_293
	VirtualNetwork_294
	VirtualNetwork_295
	VirtualNetwork_296
	VirtualNetwork_297
	VirtualNetwork_298
	VirtualNetwork_299
	VirtualNetwork_300
	VirtualNetwork_301
	VirtualNetwork_302
	VirtualNetwork_303
	VirtualNetwork_304
	VirtualNetwork_305
	VirtualNetwork_306
	VirtualNetwork_307
	VirtualNetwork_308
	VirtualNetwork_309
	VirtualNetwork_310
	VirtualNetwork_311
	VirtualNetwork_312
	VirtualNetwork_313
	VirtualNetwork_314
	VirtualNetwork_315
	VirtualNetwork_316
	VirtualNetwork_317
	VirtualNetwork_318
	VirtualNetwork_319
	VirtualNetwork_320
	VirtualNetwork_321
	VirtualNetwork_322
	VirtualNetwork_323
	VirtualNetwork_324
	VirtualNetwork_325
	VirtualNetwork_326
	VirtualNetwork_327
	VirtualNetwork_328
	VirtualNetwork_329
	VirtualNetwork_330
	VirtualNetwork_331
	VirtualNetwork_332
	VirtualNetwork_333
	VirtualNetwork_334
	VirtualNetwork_335
	VirtualNetwork_336
	VirtualNetwork_337
	VirtualNetwork_338
	VirtualNetwork_339
	VirtualNetwork_340
	VirtualNetwork_341
	VirtualNetwork_342
	VirtualNetwork_343
	VirtualNetwork_344
	VirtualNetwork_345
	VirtualNetwork_346
	VirtualNetwork_347
	VirtualNetwork_348
	VirtualNetwork_349
	VirtualNetwork_350
	VirtualNetwork_351
	VirtualNetwork_352
	VirtualNetwork_353
	VirtualNetwork_354
	VirtualNetwork_355
	VirtualNetwork_356
	VirtualNetwork_357
	VirtualNetwork_358
	VirtualNetwork_359
	VirtualNetwork_360
	VirtualNetwork_361
	VirtualNetwork_362
	VirtualNetwork_363
	VirtualNetwork_364
	VirtualNetwork_365
	VirtualNetwork_366
	VirtualNetwork_367
	VirtualNetwork_368
	VirtualNetwork_369
	VirtualNetwork_370
	VirtualNetwork_371
	VirtualNetwork_372
	VirtualNetwork_373
	VirtualNetwork_374
	VirtualNetwork_375
	VirtualNetwork_376
	VirtualNetwork_377
	VirtualNetwork_378
	VirtualNetwork_379
	VirtualNetwork_380
	VirtualNetwork_381
	VirtualNetwork_382
	VirtualNetwork_383
	VirtualNetwork_384
	VirtualNetwork_385
	VirtualNetwork_386
	VirtualNetwork_387
	VirtualNetwork_388
	VirtualNetwork_389
	VirtualNetwork_390
	VirtualNetwork_391
	VirtualNetwork_392
	VirtualNetwork_393
	VirtualNetwork_394
	VirtualNetwork_395
	VirtualNetwork_396
	VirtualNetwork_397
	VirtualNetwork_398
	VirtualNetwork_399
	VirtualNetwork_400
	VirtualNetwork_401
	VirtualNetwork_402
	VirtualNetwork_403
	VirtualNetwork_404
	VirtualNetwork_405
	VirtualNetwork_406
	VirtualNetwork_407
	VirtualNetwork_408
	VirtualNetwork_409
	VirtualNetwork_410
	VirtualNetwork_411
	VirtualNetwork_412
	VirtualNetwork_413
	VirtualNetwork_414
	VirtualNetwork_415
	VirtualNetwork_416
	VirtualNetwork_417
	VirtualNetwork_418
	VirtualNetwork_419
	VirtualNetwork_420
	VirtualNetwork_421
	VirtualNetwork_422
	VirtualNetwork_423
	VirtualNetwork_424
	VirtualNetwork_425
	VirtualNetwork_426
	VirtualNetwork_427
	VirtualNetwork_428
	VirtualNetwork_429
	VirtualNetwork_430
	VirtualNetwork_431
	VirtualNetwork_432
	VirtualNetwork_433
	VirtualNetwork_434
	VirtualNetwork_435
	VirtualNetwork_436
	VirtualNetwork_437
	VirtualNetwork_438
	VirtualNetwork_439
	VirtualNetwork_440
	VirtualNetwork_441
	VirtualNetwork_442
	VirtualNetwork_443
	VirtualNetwork_444
	VirtualNetwork_445
	VirtualNetwork_446
	VirtualNetwork_447
	VirtualNetwork_448
	VirtualNetwork_449
	VirtualNetwork_450
	VirtualNetwork_451
	VirtualNetwork_452
	VirtualNetwork_453
	VirtualNetwork_454
	VirtualNetwork_455
	VirtualNetwork_456
	VirtualNetwork_457
	VirtualNetwork_458
	VirtualNetwork_459
	VirtualNetwork_460
	VirtualNetwork_461
	VirtualNetwork_462
	VirtualNetwork_463
	VirtualNetwork_464
	VirtualNetwork_465
	VirtualNetwork_466
	VirtualNetwork_467
	VirtualNetwork_468
	VirtualNetwork_469
	VirtualNetwork_470
	VirtualNetwork_471
	VirtualNetwork_472
	VirtualNetwork_473
	VirtualNetwork_474
	VirtualNetwork_475
	VirtualNetwork_476
	VirtualNetwork_477
	VirtualNetwork_478
	VirtualNetwork_479
	VirtualNetwork_480
	VirtualNetwork_481
	VirtualNetwork_482
	VirtualNetwork_483
	VirtualNetwork_484
	VirtualNetwork_485
	VirtualNetwork_486
	VirtualNetwork_487
	VirtualNetwork_488
	VirtualNetwork_489
	VirtualNetwork_490
	VirtualNetwork_491
	VirtualNetwork_492
	VirtualNetwork_493
	VirtualNetwork_494
	VirtualNetwork_495
	VirtualNetwork_496
	VirtualNetwork_497
	VirtualNetwork_498
	VirtualNetwork_499
	VirtualNetwork_500
)
