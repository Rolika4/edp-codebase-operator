package handler

import (
	ctrl "sigs.k8s.io/controller-runtime"

	codebaseApi "github.com/epam/edp-codebase-operator/v2/pkg/apis/edp/v1"
)

type CodebaseBranchHandler interface {
	ServeRequest(c *codebaseApi.CodebaseBranch) error
}

var log = ctrl.Log.WithName("codebase_branch_handler")

func NextServeOrNil(next CodebaseBranchHandler, cb *codebaseApi.CodebaseBranch) error {
	if next != nil {
		return next.ServeRequest(cb)
	}
	log.Info("handling of codebase branch has been finished", "name", cb.Name)
	return nil
}
