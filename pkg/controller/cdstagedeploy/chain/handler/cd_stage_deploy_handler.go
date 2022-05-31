package handler

import codebaseApi "github.com/epam/edp-codebase-operator/v2/pkg/apis/edp/v1"

type CDStageDeployHandler interface {
	ServeRequest(stageDeploy *codebaseApi.CDStageDeploy) error
}
