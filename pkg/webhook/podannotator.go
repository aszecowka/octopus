package webhook

import (
	"context"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/runtime/inject"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission/types"
)

type PodAnnotator struct {
	cli     client.Client
	decoder types.Decoder
}

var _ admission.Handler = &PodAnnotator{}

func (a *PodAnnotator) Handle(ctx context.Context, req types.Request) types.Response {
	return types.Response{}
}

var _ inject.Client = &PodAnnotator{}

func (a *PodAnnotator) InjectClient(c client.Client) error {
	a.cli = c
	return nil
}

var _ inject.Decoder = &PodAnnotator{}

func (a *PodAnnotator) InjectDecoder(d types.Decoder) error {
	a.decoder = d
	return nil
}
