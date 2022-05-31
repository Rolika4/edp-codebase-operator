package chain

import (
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	codebaseApi "github.com/epam/edp-codebase-operator/v2/pkg/apis/edp/v1"
	"github.com/epam/edp-codebase-operator/v2/pkg/client/jira"
	"github.com/epam/edp-codebase-operator/v2/pkg/controller/jiraserver/chain/handler"
)

var log = ctrl.Log.WithName("jira_server_handler")

func CreateDefChain(jc jira.Client, client client.Client) handler.JiraServerHandler {
	return CheckConnection{
		next: PutJiraEDPComponent{
			next:   nil,
			client: client,
		},
		client: jc,
	}
}

func nextServeOrNil(next handler.JiraServerHandler, jira *codebaseApi.JiraServer) error {
	if next != nil {
		return next.ServeRequest(jira)
	}
	log.Info("handling of JiraServer has been finished", "jira server name", jira.Name)
	return nil
}
