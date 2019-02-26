package lightstep_test

import (
	. "github.com/lightstep/lightstep-tracer-go"
	"github.com/lightstep/lightstep-tracer-go/collectorpb"
	"github.com/lightstep/lightstep-tracer-go/collectorpb/collectorpbfakes"
)

type cpbfakesFakeClient struct {
	collectorpbfakes.FakeCollectorServiceClient
}

func newGrpcFakeClient() fakeCollectorClient {
	fakeClient := new(collectorpbfakes.FakeCollectorServiceClient)
	fakeClient.ReportReturns(&collectorpb.ReportResponse{}, nil)
	return &cpbfakesFakeClient{FakeCollectorServiceClient: *fakeClient}
}

func (fakeClient *cpbfakesFakeClient) ConnectorFactory() ConnectorFactory {
	return fakeGrpcConnection(&fakeClient.FakeCollectorServiceClient)
}

func (fakeClient *cpbfakesFakeClient) getSpans() []*collectorpb.Span {
	return getReportedGRPCSpans(&fakeClient.FakeCollectorServiceClient)
}

func (fakeClient *cpbfakesFakeClient) GetSpansLen() int {
	return len(fakeClient.getSpans())
}
