// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0
// source: parcel/parcel.proto

package parcelpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Status int32

const (
	Status_UNKNOWN    Status = 0
	Status_PENDING    Status = 1
	Status_IN_TRANSIT Status = 2
	Status_DELIVERED  Status = 3
	Status_CANCELED   Status = 4
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "UNKNOWN",
		1: "PENDING",
		2: "IN_TRANSIT",
		3: "DELIVERED",
		4: "CANCELED",
	}
	Status_value = map[string]int32{
		"UNKNOWN":    0,
		"PENDING":    1,
		"IN_TRANSIT": 2,
		"DELIVERED":  3,
		"CANCELED":   4,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_parcel_parcel_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_parcel_parcel_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_parcel_parcel_proto_rawDescGZIP(), []int{0}
}

type Parcel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrackNumber       string                 `protobuf:"bytes,1,opt,name=track_number,json=trackNumber,proto3" json:"track_number,omitempty"`
	Name              string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ManagerTelegramId int64                  `protobuf:"varint,3,opt,name=manager_telegram_id,json=managerTelegramId,proto3" json:"manager_telegram_id,omitempty"`
	Recipient         string                 `protobuf:"bytes,4,opt,name=recipient,proto3" json:"recipient,omitempty"`
	ArrivalAddress    string                 `protobuf:"bytes,5,opt,name=arrival_address,json=arrivalAddress,proto3" json:"arrival_address,omitempty"`
	ForecastDate      *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=forecast_date,json=forecastDate,proto3" json:"forecast_date,omitempty"`
	Description       string                 `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	Status            Status                 `protobuf:"varint,8,opt,name=status,proto3,enum=parcelpb.Status" json:"status,omitempty"`
}

func (x *Parcel) Reset() {
	*x = Parcel{}
	mi := &file_parcel_parcel_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Parcel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Parcel) ProtoMessage() {}

func (x *Parcel) ProtoReflect() protoreflect.Message {
	mi := &file_parcel_parcel_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Parcel.ProtoReflect.Descriptor instead.
func (*Parcel) Descriptor() ([]byte, []int) {
	return file_parcel_parcel_proto_rawDescGZIP(), []int{0}
}

func (x *Parcel) GetTrackNumber() string {
	if x != nil {
		return x.TrackNumber
	}
	return ""
}

func (x *Parcel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Parcel) GetManagerTelegramId() int64 {
	if x != nil {
		return x.ManagerTelegramId
	}
	return 0
}

func (x *Parcel) GetRecipient() string {
	if x != nil {
		return x.Recipient
	}
	return ""
}

func (x *Parcel) GetArrivalAddress() string {
	if x != nil {
		return x.ArrivalAddress
	}
	return ""
}

func (x *Parcel) GetForecastDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ForecastDate
	}
	return nil
}

func (x *Parcel) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Parcel) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_UNKNOWN
}

type Checkpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time         *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Place        string                 `protobuf:"bytes,2,opt,name=place,proto3" json:"place,omitempty"`
	Description  string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	ParcelStatus Status                 `protobuf:"varint,4,opt,name=parcel_status,json=parcelStatus,proto3,enum=parcelpb.Status" json:"parcel_status,omitempty"`
}

func (x *Checkpoint) Reset() {
	*x = Checkpoint{}
	mi := &file_parcel_parcel_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Checkpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Checkpoint) ProtoMessage() {}

func (x *Checkpoint) ProtoReflect() protoreflect.Message {
	mi := &file_parcel_parcel_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Checkpoint.ProtoReflect.Descriptor instead.
func (*Checkpoint) Descriptor() ([]byte, []int) {
	return file_parcel_parcel_proto_rawDescGZIP(), []int{1}
}

func (x *Checkpoint) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Checkpoint) GetPlace() string {
	if x != nil {
		return x.Place
	}
	return ""
}

func (x *Checkpoint) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Checkpoint) GetParcelStatus() Status {
	if x != nil {
		return x.ParcelStatus
	}
	return Status_UNKNOWN
}

var File_parcel_parcel_proto protoreflect.FileDescriptor

var file_parcel_parcel_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x61, 0x72, 0x63, 0x65, 0x6c, 0x2f, 0x70, 0x61, 0x72, 0x63, 0x65, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x61, 0x72, 0x63, 0x65, 0x6c, 0x70, 0x62, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xc3, 0x02, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x63, 0x65, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x74, 0x65,
	0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x11, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x54, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d,
	0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74,
	0x12, 0x27, 0x0a, 0x0f, 0x61, 0x72, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x72, 0x72, 0x69, 0x76,
	0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x3f, 0x0a, 0x0d, 0x66, 0x6f, 0x72,
	0x65, 0x63, 0x61, 0x73, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x66, 0x6f,
	0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70,
	0x61, 0x72, 0x63, 0x65, 0x6c, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xab, 0x01, 0x0a, 0x0a, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a,
	0x0d, 0x70, 0x61, 0x72, 0x63, 0x65, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70, 0x61, 0x72, 0x63, 0x65, 0x6c, 0x70, 0x62, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x63, 0x65, 0x6c, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2a, 0x4f, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50,
	0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4e, 0x5f, 0x54,
	0x52, 0x41, 0x4e, 0x53, 0x49, 0x54, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x45, 0x4c, 0x49,
	0x56, 0x45, 0x52, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x41, 0x4e, 0x43, 0x45,
	0x4c, 0x45, 0x44, 0x10, 0x04, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x75, 0x62, 0x6c, 0x69, 0x6b, 0x65, 0x72, 0x2f, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x2d, 0x70, 0x61, 0x72, 0x63, 0x65, 0x6c, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x70, 0x61, 0x72, 0x63, 0x65, 0x6c, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_parcel_parcel_proto_rawDescOnce sync.Once
	file_parcel_parcel_proto_rawDescData = file_parcel_parcel_proto_rawDesc
)

func file_parcel_parcel_proto_rawDescGZIP() []byte {
	file_parcel_parcel_proto_rawDescOnce.Do(func() {
		file_parcel_parcel_proto_rawDescData = protoimpl.X.CompressGZIP(file_parcel_parcel_proto_rawDescData)
	})
	return file_parcel_parcel_proto_rawDescData
}

var file_parcel_parcel_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_parcel_parcel_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_parcel_parcel_proto_goTypes = []any{
	(Status)(0),                   // 0: parcelpb.Status
	(*Parcel)(nil),                // 1: parcelpb.Parcel
	(*Checkpoint)(nil),            // 2: parcelpb.Checkpoint
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_parcel_parcel_proto_depIdxs = []int32{
	3, // 0: parcelpb.Parcel.forecast_date:type_name -> google.protobuf.Timestamp
	0, // 1: parcelpb.Parcel.status:type_name -> parcelpb.Status
	3, // 2: parcelpb.Checkpoint.time:type_name -> google.protobuf.Timestamp
	0, // 3: parcelpb.Checkpoint.parcel_status:type_name -> parcelpb.Status
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_parcel_parcel_proto_init() }
func file_parcel_parcel_proto_init() {
	if File_parcel_parcel_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_parcel_parcel_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_parcel_parcel_proto_goTypes,
		DependencyIndexes: file_parcel_parcel_proto_depIdxs,
		EnumInfos:         file_parcel_parcel_proto_enumTypes,
		MessageInfos:      file_parcel_parcel_proto_msgTypes,
	}.Build()
	File_parcel_parcel_proto = out.File
	file_parcel_parcel_proto_rawDesc = nil
	file_parcel_parcel_proto_goTypes = nil
	file_parcel_parcel_proto_depIdxs = nil
}
