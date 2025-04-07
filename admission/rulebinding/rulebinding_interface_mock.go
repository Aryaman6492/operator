package rulebinding

import (
	"context"

	"github.com/Aryaman6492/operator/admission/rules"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var _ RuleBindingCache = (*RuleBindingCacheMock)(nil)

type RuleBindingCacheMock struct {
}

func (r *RuleBindingCacheMock) ListRulesForObject(_ context.Context, _ *unstructured.Unstructured) []rules.RuleEvaluator {
	return []rules.RuleEvaluator{}
}
