package tibrv

const (
	// TibrvOK import C API define : TIBRV_OK
	TibrvOK = 0
	// TibrvINITFailure import C API define : TIBRV_INIT_FAILURE
	TibrvINITFailure = 1
	// TibrvInvalidTransport import C API define : TIBRV_INVALID_TRANSPORT
	TibrvInvalidTransport = 2
	// TibrvInvalidArg import C API define : TIBRV_INVALID_ARG
	TibrvInvalidArg = 3
	// TibrvNotInitialized import C API define : TIBRV_NOT_INITIALIZED
	TibrvNotInitialized = 4
	// TibrvArgConflict import C API define : TIBRV_ARG_CONFLICT
	TibrvArgConflict = 5
	// TibrvServiceNotFound import C API define : TIBRV_SERVICE_NOT_FOUND
	TibrvServiceNotFound = 16
	// TibrvNetworkNotFound import C API define : TIBRV_NETWORK_NOT_FOUND
	TibrvNetworkNotFound = 17
	// TibrvDaemonNotFound import C API define : TIBRV_DAEMON_NOT_FOUND
	TibrvDaemonNotFound = 18
	// TibrvNoMemory import C API define : TIBRV_NO_MEMORY
	TibrvNoMemory = 19
	// TibrvInvalidSubject import C API define : TIBRV_INVALID_SUBJECT
	TibrvInvalidSubject = 20
	// TibrvDaemonNotConnected import C API define : TIBRV_DAEMON_NOT_CONNECTED
	TibrvDaemonNotConnected = 21
	// TibrvVersionMismatch import C API define : TIBRV_VERSION_MISMATCH
	TibrvVersionMismatch = 22
	// TibrvSubjectCollision import C API define : TIBRV_SUBJECT_COLLISION
	TibrvSubjectCollision = 23
	// TibrvVcNotConnected import C API define : TIBRV_VC_NOT_CONNECTED
	TibrvVcNotConnected = 24
	// TibrvNotPermitted import C API define : TIBRV_NOT_PERMITTED
	TibrvNotPermitted = 27
	// TibrvInvalidName import C API define : TIBRV_INVALID_NAME
	TibrvInvalidName = 30
	// TibrvInvalidType import C API define : TIBRV_INVALID_TYPE
	TibrvInvalidType = 31
	// TibrvInvalidSize import C API define : TIBRV_INVALID_SIZE
	TibrvInvalidSize = 32
	// TibrvInvalidCount import C API define : TIBRV_INVALID_COUNT
	TibrvInvalidCount = 33
	// TibrvNotFound import C API define : TIBRV_NOT_FOUND
	TibrvNotFound = 35
	// TibrvIDInUse import C API define : TIBRV_ID_IN_USE
	TibrvIDInUse = 36
	// TibrvIDConflict import C API define : TIBRV_ID_CONFICT
	TibrvIDConflict = 37
	// TibrvConversionFailed import C API define : TIBRV_CONVERSION_FAILED
	TibrvConversionFailed = 38
	// TibrvReservedHandler import C API define : TIBRV_RESERVED_HANDLER
	TibrvReservedHandler = 39
	// TibrvEncoderFailed import C API define : TIBRV_ENCODER_FAILED
	TibrvEncoderFailed = 40
	// TibrvDecoderFailed import C API define : TIBRV_DECODER_FAILED
	TibrvDecoderFailed = 41
	// TibrvInvalidMsg import C API define : TIBRV_INVALID_MSG
	TibrvInvalidMsg = 42
	// TibrvInvalidField import C API define : TIBRV_INVALID_FIELD
	TibrvInvalidField = 43
	// TibrvInvalidInstance import C API define : TIBRV_INVALID_INSTANCE
	TibrvInvalidInstance = 44
	// TibrvCorruptMsg import C API define : TIBRV_CORRUPT_MSG
	TibrvCorruptMsg = 45
	// TibrvEncodingMismatch import C API define : TIBRV_ENCODING_MISMATCH
	TibrvEncodingMismatch = 46
	// TibrvTimeout import C API define : TIBRV_TIMEOUT
	TibrvTimeout = 50
	// TibrvIntr import C API define : TIBRV_INTR
	TibrvIntr = 51
	// TibrvInvalidDispatchable import C API define : TIBRV_INVALID_DISPATCHABLE
	TibrvInvalidDispatchable = 52
	// TibrvInvalidDispatcher import C API define : TIBRV_INVALID_DISPATCHER
	TibrvInvalidDispatcher = 53
	// TibrvInvalidEvent import C API define : TIBRV_INVALID_EVENT
	TibrvInvalidEvent = 60
	// TibrvInvalidCallback import C API define : TIBRV_INVALID_CALLBACK
	TibrvInvalidCallback = 61
	// TibrvInvalidQueue import C API define : TIBRV_INVALID_QUEUE
	TibrvInvalidQueue = 62
	// TibrvInvalidQueueGroup import C API define : TIBRV_INVALID_QUEUE_GROUP
	TibrvInvalidQueueGroup = 63
	// TibrvInvalidTimeInterval import C API define : TIBRV_INVALID_TIME_INTERVAL
	TibrvInvalidTimeInterval = 64
	// TibrvInvalidIOSource import C API define : TIBRV_INVALID_IO_SOURCE
	TibrvInvalidIOSource = 65
	// TibrvInvalidIOCondition import C API define : TIBRV_INVALID_IO_CONDITION
	TibrvInvalidIOCondition = 66
	// TibrvSocketLimit import C API define : TIBRV_SOCKET_LIMIT
	TibrvSocketLimit = 67
	// TibrvOsError import C API define : TIBRV_OS_ERROR
	TibrvOsError = 68
	// TibrvInsufficientBuffer import C API define : TIBRV_INSUFFICIENT_BUFFER
	TibrvInsufficientBuffer = 70
	// TibrvEOF import C API define : TIBRV_EOF
	TibrvEOF = 71
	// TibrvInvalidFile import C API define : TIBRV_INVALID_FILE
	TibrvInvalidFile = 72
	// TibrvFileNotFound import C API define : TIBRV_FILE_NOT_FOUND
	TibrvFileNotFound = 73
	// TibrvIOFailed import C API define : TIBRV_IO_FAILED
	TibrvIOFailed = 74
	// TibrvNotFileOwner import C API define : TIBRV_NOT_FILE_OWNER
	TibrvNotFileOwner = 80
	// TibrvUuserpassMismatch import C API define : TIBRV_USERPASS_MISMATCH
	TibrvUuserpassMismatch = 81
	// TibrvTooManyNeighboors import C API define : TIBRV_TOO_MANY_NEIGHBORS
	TibrvTooManyNeighboors = 90
	// TibrvAlreadyExists import C API define : TIBRV_ALREADY_EXISTS
	TibrvAlreadyExists = 91
	// TibrvPortBusy import C API define : TIBRV_PORT_BUSY
	TibrvPortBusy = 100
	// TibrvDeliveryFailed import C API define : TIBRV_DELIVERY_FAILED
	TibrvDeliveryFailed = 101
	// TibrvQueueLimit import C API define : TIBRV_QUEUE_LIMIT
	TibrvQueueLimit = 102
	// TibrvInvalidContentDesc import C API define : TIBRV_INVALID_CONTENT_DESC
	TibrvInvalidContentDesc = 110
	// TibrvInvalidSerializedBuffer import C API define : TIBRV_INVALID_SERIALIZED_BUFFER
	TibrvInvalidSerializedBuffer = 111
	// TibrvDescriptorNotFound import C API define : TIBRV_DESCRIPTOR_NOT_FOUND
	TibrvDescriptorNotFound = 115
	// TibrvCorruptSerializedBuffer import C API define : TIBRV_CORRUPT_SERIALIZED_BUFFER
	TibrvCorruptSerializedBuffer = 116
	// TibrvIPMOnly import C API define : TIBRV_IPM_ONLY
	TibrvIPMOnly = 117
)
