package rulebinding

import (
	"context"

	"github.com/Aryaman6492/operator/admission/rules"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type RuleBindingCache interface {
	ListRulesForObject(ctx context.Context, object *unstructured.Unstructured) []rules.RuleEvaluator
}
