// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EmrServiceClient is the client API for EmrService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmrServiceClient interface {
	PatientInfo(ctx context.Context, in *Request, opts ...grpc.CallOption) (*InfoResponse, error)
	ServiceList(ctx context.Context, in *ServiceListRequest, opts ...grpc.CallOption) (*ServiceListResponse, error)
	OPDScreening(ctx context.Context, in *VisitRequest, opts ...grpc.CallOption) (*OPDScreenResponse, error)
	OPDDiagnosis(ctx context.Context, in *VisitRequest, opts ...grpc.CallOption) (*OPDDiagResponse, error)
	OPDProcedure(ctx context.Context, in *VisitRequest, opts ...grpc.CallOption) (*OPDProcedureResponse, error)
	OPDDrug(ctx context.Context, in *VisitRequest, opts ...grpc.CallOption) (*OPDDrugResponse, error)
	Lab(ctx context.Context, in *VisitRequest, opts ...grpc.CallOption) (*LabResponse, error)
	Appointment(ctx context.Context, in *Request, opts ...grpc.CallOption) (*AppointmentResponse, error)
	ReferOut(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ReferOutResponse, error)
}

type emrServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmrServiceClient(cc grpc.ClientConnInterface) EmrServiceClient {
	return &emrServiceClient{cc}
}

func (c *emrServiceClient) PatientInfo(ctx context.Context, in *Request, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := c.cc.Invoke(ctx, "/proto.EmrService/PatientInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emrServiceClient) ServiceList(ctx context.Context, in *ServiceListRequest, opts ...grpc.CallOption) (*ServiceListResponse, error) {
	out := new(ServiceListResponse)
	err := c.cc.Invoke(ctx, "/proto.EmrService/ServiceList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emrServiceClient) OPDScreening(ctx context.Context, in *VisitRequest, opts ...grpc.CallOption) (*OPDScreenResponse, error) {
	out := new(OPDScreenResponse)
	err := c.cc.Invoke(ctx, "/proto.EmrService/OPDScreening", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emrServiceClient) OPDDiagnosis(ctx context.Context, in *VisitRequest, opts ...grpc.CallOption) (*OPDDiagResponse, error) {
	out := new(OPDDiagResponse)
	err := c.cc.Invoke(ctx, "/proto.EmrService/OPDDiagnosis", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emrServiceClient) OPDProcedure(ctx context.Context, in *VisitRequest, opts ...grpc.CallOption) (*OPDProcedureResponse, error) {
	out := new(OPDProcedureResponse)
	err := c.cc.Invoke(ctx, "/proto.EmrService/OPDProcedure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emrServiceClient) OPDDrug(ctx context.Context, in *VisitRequest, opts ...grpc.CallOption) (*OPDDrugResponse, error) {
	out := new(OPDDrugResponse)
	err := c.cc.Invoke(ctx, "/proto.EmrService/OPDDrug", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emrServiceClient) Lab(ctx context.Context, in *VisitRequest, opts ...grpc.CallOption) (*LabResponse, error) {
	out := new(LabResponse)
	err := c.cc.Invoke(ctx, "/proto.EmrService/Lab", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emrServiceClient) Appointment(ctx context.Context, in *Request, opts ...grpc.CallOption) (*AppointmentResponse, error) {
	out := new(AppointmentResponse)
	err := c.cc.Invoke(ctx, "/proto.EmrService/Appointment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emrServiceClient) ReferOut(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ReferOutResponse, error) {
	out := new(ReferOutResponse)
	err := c.cc.Invoke(ctx, "/proto.EmrService/ReferOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmrServiceServer is the server API for EmrService service.
// All implementations must embed UnimplementedEmrServiceServer
// for forward compatibility
type EmrServiceServer interface {
	PatientInfo(context.Context, *Request) (*InfoResponse, error)
	ServiceList(context.Context, *ServiceListRequest) (*ServiceListResponse, error)
	OPDScreening(context.Context, *VisitRequest) (*OPDScreenResponse, error)
	OPDDiagnosis(context.Context, *VisitRequest) (*OPDDiagResponse, error)
	OPDProcedure(context.Context, *VisitRequest) (*OPDProcedureResponse, error)
	OPDDrug(context.Context, *VisitRequest) (*OPDDrugResponse, error)
	Lab(context.Context, *VisitRequest) (*LabResponse, error)
	Appointment(context.Context, *Request) (*AppointmentResponse, error)
	ReferOut(context.Context, *Request) (*ReferOutResponse, error)
	mustEmbedUnimplementedEmrServiceServer()
}

// UnimplementedEmrServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEmrServiceServer struct {
}

func (UnimplementedEmrServiceServer) PatientInfo(context.Context, *Request) (*InfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatientInfo not implemented")
}
func (UnimplementedEmrServiceServer) ServiceList(context.Context, *ServiceListRequest) (*ServiceListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServiceList not implemented")
}
func (UnimplementedEmrServiceServer) OPDScreening(context.Context, *VisitRequest) (*OPDScreenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OPDScreening not implemented")
}
func (UnimplementedEmrServiceServer) OPDDiagnosis(context.Context, *VisitRequest) (*OPDDiagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OPDDiagnosis not implemented")
}
func (UnimplementedEmrServiceServer) OPDProcedure(context.Context, *VisitRequest) (*OPDProcedureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OPDProcedure not implemented")
}
func (UnimplementedEmrServiceServer) OPDDrug(context.Context, *VisitRequest) (*OPDDrugResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OPDDrug not implemented")
}
func (UnimplementedEmrServiceServer) Lab(context.Context, *VisitRequest) (*LabResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Lab not implemented")
}
func (UnimplementedEmrServiceServer) Appointment(context.Context, *Request) (*AppointmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Appointment not implemented")
}
func (UnimplementedEmrServiceServer) ReferOut(context.Context, *Request) (*ReferOutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReferOut not implemented")
}
func (UnimplementedEmrServiceServer) mustEmbedUnimplementedEmrServiceServer() {}

// UnsafeEmrServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmrServiceServer will
// result in compilation errors.
type UnsafeEmrServiceServer interface {
	mustEmbedUnimplementedEmrServiceServer()
}

func RegisterEmrServiceServer(s grpc.ServiceRegistrar, srv EmrServiceServer) {
	s.RegisterService(&EmrService_ServiceDesc, srv)
}

func _EmrService_PatientInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmrServiceServer).PatientInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EmrService/PatientInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmrServiceServer).PatientInfo(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmrService_ServiceList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmrServiceServer).ServiceList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EmrService/ServiceList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmrServiceServer).ServiceList(ctx, req.(*ServiceListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmrService_OPDScreening_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VisitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmrServiceServer).OPDScreening(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EmrService/OPDScreening",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmrServiceServer).OPDScreening(ctx, req.(*VisitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmrService_OPDDiagnosis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VisitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmrServiceServer).OPDDiagnosis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EmrService/OPDDiagnosis",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmrServiceServer).OPDDiagnosis(ctx, req.(*VisitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmrService_OPDProcedure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VisitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmrServiceServer).OPDProcedure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EmrService/OPDProcedure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmrServiceServer).OPDProcedure(ctx, req.(*VisitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmrService_OPDDrug_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VisitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmrServiceServer).OPDDrug(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EmrService/OPDDrug",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmrServiceServer).OPDDrug(ctx, req.(*VisitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmrService_Lab_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VisitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmrServiceServer).Lab(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EmrService/Lab",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmrServiceServer).Lab(ctx, req.(*VisitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmrService_Appointment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmrServiceServer).Appointment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EmrService/Appointment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmrServiceServer).Appointment(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmrService_ReferOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmrServiceServer).ReferOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EmrService/ReferOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmrServiceServer).ReferOut(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// EmrService_ServiceDesc is the grpc.ServiceDesc for EmrService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmrService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.EmrService",
	HandlerType: (*EmrServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PatientInfo",
			Handler:    _EmrService_PatientInfo_Handler,
		},
		{
			MethodName: "ServiceList",
			Handler:    _EmrService_ServiceList_Handler,
		},
		{
			MethodName: "OPDScreening",
			Handler:    _EmrService_OPDScreening_Handler,
		},
		{
			MethodName: "OPDDiagnosis",
			Handler:    _EmrService_OPDDiagnosis_Handler,
		},
		{
			MethodName: "OPDProcedure",
			Handler:    _EmrService_OPDProcedure_Handler,
		},
		{
			MethodName: "OPDDrug",
			Handler:    _EmrService_OPDDrug_Handler,
		},
		{
			MethodName: "Lab",
			Handler:    _EmrService_Lab_Handler,
		},
		{
			MethodName: "Appointment",
			Handler:    _EmrService_Appointment_Handler,
		},
		{
			MethodName: "ReferOut",
			Handler:    _EmrService_ReferOut_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "his.proto",
}

// MhealthServiceClient is the client API for MhealthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MhealthServiceClient interface {
	Appointment(ctx context.Context, in *AppointmentRequest, opts ...grpc.CallOption) (*AppointmentResponse, error)
}

type mhealthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMhealthServiceClient(cc grpc.ClientConnInterface) MhealthServiceClient {
	return &mhealthServiceClient{cc}
}

func (c *mhealthServiceClient) Appointment(ctx context.Context, in *AppointmentRequest, opts ...grpc.CallOption) (*AppointmentResponse, error) {
	out := new(AppointmentResponse)
	err := c.cc.Invoke(ctx, "/proto.MhealthService/Appointment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MhealthServiceServer is the server API for MhealthService service.
// All implementations must embed UnimplementedMhealthServiceServer
// for forward compatibility
type MhealthServiceServer interface {
	Appointment(context.Context, *AppointmentRequest) (*AppointmentResponse, error)
	mustEmbedUnimplementedMhealthServiceServer()
}

// UnimplementedMhealthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMhealthServiceServer struct {
}

func (UnimplementedMhealthServiceServer) Appointment(context.Context, *AppointmentRequest) (*AppointmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Appointment not implemented")
}
func (UnimplementedMhealthServiceServer) mustEmbedUnimplementedMhealthServiceServer() {}

// UnsafeMhealthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MhealthServiceServer will
// result in compilation errors.
type UnsafeMhealthServiceServer interface {
	mustEmbedUnimplementedMhealthServiceServer()
}

func RegisterMhealthServiceServer(s grpc.ServiceRegistrar, srv MhealthServiceServer) {
	s.RegisterService(&MhealthService_ServiceDesc, srv)
}

func _MhealthService_Appointment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppointmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MhealthServiceServer).Appointment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MhealthService/Appointment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MhealthServiceServer).Appointment(ctx, req.(*AppointmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MhealthService_ServiceDesc is the grpc.ServiceDesc for MhealthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MhealthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MhealthService",
	HandlerType: (*MhealthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Appointment",
			Handler:    _MhealthService_Appointment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "his.proto",
}

// MasterServiceClient is the client API for MasterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MasterServiceClient interface {
	DoctorList(ctx context.Context, in *RequestHospcode, opts ...grpc.CallOption) (*ListDoctorResponse, error)
	ClinicList(ctx context.Context, in *RequestHospcode, opts ...grpc.CallOption) (*ListClinicResponse, error)
	HisProviderList(ctx context.Context, in *RequestHospcode, opts ...grpc.CallOption) (*HisProviderResponse, error)
}

type masterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMasterServiceClient(cc grpc.ClientConnInterface) MasterServiceClient {
	return &masterServiceClient{cc}
}

func (c *masterServiceClient) DoctorList(ctx context.Context, in *RequestHospcode, opts ...grpc.CallOption) (*ListDoctorResponse, error) {
	out := new(ListDoctorResponse)
	err := c.cc.Invoke(ctx, "/proto.MasterService/DoctorList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) ClinicList(ctx context.Context, in *RequestHospcode, opts ...grpc.CallOption) (*ListClinicResponse, error) {
	out := new(ListClinicResponse)
	err := c.cc.Invoke(ctx, "/proto.MasterService/ClinicList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) HisProviderList(ctx context.Context, in *RequestHospcode, opts ...grpc.CallOption) (*HisProviderResponse, error) {
	out := new(HisProviderResponse)
	err := c.cc.Invoke(ctx, "/proto.MasterService/HisProviderList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterServiceServer is the server API for MasterService service.
// All implementations must embed UnimplementedMasterServiceServer
// for forward compatibility
type MasterServiceServer interface {
	DoctorList(context.Context, *RequestHospcode) (*ListDoctorResponse, error)
	ClinicList(context.Context, *RequestHospcode) (*ListClinicResponse, error)
	HisProviderList(context.Context, *RequestHospcode) (*HisProviderResponse, error)
	mustEmbedUnimplementedMasterServiceServer()
}

// UnimplementedMasterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMasterServiceServer struct {
}

func (UnimplementedMasterServiceServer) DoctorList(context.Context, *RequestHospcode) (*ListDoctorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoctorList not implemented")
}
func (UnimplementedMasterServiceServer) ClinicList(context.Context, *RequestHospcode) (*ListClinicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClinicList not implemented")
}
func (UnimplementedMasterServiceServer) HisProviderList(context.Context, *RequestHospcode) (*HisProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HisProviderList not implemented")
}
func (UnimplementedMasterServiceServer) mustEmbedUnimplementedMasterServiceServer() {}

// UnsafeMasterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MasterServiceServer will
// result in compilation errors.
type UnsafeMasterServiceServer interface {
	mustEmbedUnimplementedMasterServiceServer()
}

func RegisterMasterServiceServer(s grpc.ServiceRegistrar, srv MasterServiceServer) {
	s.RegisterService(&MasterService_ServiceDesc, srv)
}

func _MasterService_DoctorList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestHospcode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).DoctorList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MasterService/DoctorList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).DoctorList(ctx, req.(*RequestHospcode))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_ClinicList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestHospcode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).ClinicList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MasterService/ClinicList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).ClinicList(ctx, req.(*RequestHospcode))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_HisProviderList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestHospcode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).HisProviderList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MasterService/HisProviderList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).HisProviderList(ctx, req.(*RequestHospcode))
	}
	return interceptor(ctx, in, info, handler)
}

// MasterService_ServiceDesc is the grpc.ServiceDesc for MasterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MasterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MasterService",
	HandlerType: (*MasterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoctorList",
			Handler:    _MasterService_DoctorList_Handler,
		},
		{
			MethodName: "ClinicList",
			Handler:    _MasterService_ClinicList_Handler,
		},
		{
			MethodName: "HisProviderList",
			Handler:    _MasterService_HisProviderList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "his.proto",
}
