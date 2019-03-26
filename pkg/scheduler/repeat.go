package scheduler

import (
	"github.com/kyma-incubator/octopus/pkg/apis/testing/v1alpha1"
)

type repeatStrategy struct{}

func (s *repeatStrategy) GetTestToRunConcurrently(suite v1alpha1.ClusterTestSuite) *v1alpha1.TestResult {
	return s.getTest(suite, func(tr v1alpha1.TestResult) bool {
		return tr.DisableConcurrency == false
	})
}

func (s *repeatStrategy) GetTestToRunSequentially(suite v1alpha1.ClusterTestSuite) *v1alpha1.TestResult {
	return s.getTest(suite, func(tr v1alpha1.TestResult) bool {
		return tr.DisableConcurrency == true
	})
}

func (s *repeatStrategy) getTest(suite v1alpha1.ClusterTestSuite, match func(tr v1alpha1.TestResult) bool) *v1alpha1.TestResult {
	for _, tr := range suite.Status.Results {
		if !match(tr) {
			continue
		}
		if len(tr.Executions) < int(suite.Spec.Count) {
			return &tr
		}
	}
	return nil
}
