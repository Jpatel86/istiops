package operator

import (
	v1alpha32 "github.com/aspenmesh/istio-client-go/pkg/apis/networking/v1alpha3"
	"github.com/pismo/istiops/pkg/router"
)

type Operator interface {
	Get(s router.Shift) ([]v1alpha32.VirtualService, error)
	Update(s router.Shift) error
	Clear(s router.Shift) error
}
