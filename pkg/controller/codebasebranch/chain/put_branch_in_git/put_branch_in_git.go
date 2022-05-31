package put_branch_in_git

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	codebaseApi "github.com/epam/edp-codebase-operator/v2/pkg/apis/edp/v1"
	"github.com/epam/edp-codebase-operator/v2/pkg/controller/codebasebranch/chain/handler"
	"github.com/epam/edp-codebase-operator/v2/pkg/controller/codebasebranch/service"
	"github.com/epam/edp-codebase-operator/v2/pkg/controller/gitserver"
	"github.com/epam/edp-codebase-operator/v2/pkg/model"
	"github.com/epam/edp-codebase-operator/v2/pkg/util"
)

type PutBranchInGit struct {
	Next    handler.CodebaseBranchHandler
	Client  client.Client
	Git     gitserver.Git
	Service service.CodebaseBranchService
}

var log = ctrl.Log.WithName("put-branch-in-git-chain")

func (h PutBranchInGit) ServeRequest(cb *codebaseApi.CodebaseBranch) error {
	rl := log.WithValues("namespace", cb.Namespace, "codebase branch", cb.Name)
	rl.Info("start PutBranchInGit method...")

	if err := h.setIntermediateSuccessFields(cb, codebaseApi.AcceptCodebaseBranchRegistration); err != nil {
		return err
	}

	c, err := util.GetCodebase(h.Client, cb.Spec.CodebaseName, cb.Namespace)
	if err != nil {
		setFailedFields(cb, codebaseApi.PutBranchForGitlabCiCodebase, err.Error())
		return err
	}

	if !c.Status.Available {
		log.Info("couldn't start reconciling for branch. codebase is unavailable", "codebase", c.Name)
		return util.NewCodebaseBranchReconcileError(fmt.Sprintf("%v codebase is unavailable", c.Name))
	}

	if c.Spec.Versioning.Type == util.VersioningTypeEDP && hasNewVersion(cb) {
		if err := h.processNewVersion(cb); err != nil {
			err = errors.Wrapf(err, "couldn't process new version for %v branch", cb.Name)
			setFailedFields(cb, codebaseApi.PutBranchForGitlabCiCodebase, err.Error())
			return err
		}
	}

	gs, err := util.GetGitServer(h.Client, c.Spec.GitServer, c.Namespace)
	if err != nil {
		setFailedFields(cb, codebaseApi.PutBranchForGitlabCiCodebase, err.Error())
		return err
	}

	secret, err := util.GetSecret(h.Client, gs.NameSshKeySecret, c.Namespace)
	if err != nil {
		err = errors.Wrapf(err, "an error has occurred while getting %v secret", gs.NameSshKeySecret)
		setFailedFields(cb, codebaseApi.PutBranchForGitlabCiCodebase, err.Error())
		return err
	}

	wd := util.GetWorkDir(cb.Spec.CodebaseName, fmt.Sprintf("%v-%v", cb.Namespace, cb.Spec.BranchName))
	if !checkDirectory(wd) {
		ru := fmt.Sprintf("%v:%v", gs.GitHost, *c.Spec.GitUrlPath)
		if err := h.Git.CloneRepositoryBySsh(string(secret.Data[util.PrivateSShKeyName]), gs.GitUser, ru, wd, gs.SshPort); err != nil {
			setFailedFields(cb, codebaseApi.PutBranchForGitlabCiCodebase, err.Error())
			return err
		}
	}

	if err := h.Git.CreateRemoteBranch(string(secret.Data[util.PrivateSShKeyName]), gs.GitUser, wd, cb.Spec.BranchName); err != nil {
		setFailedFields(cb, codebaseApi.PutBranchForGitlabCiCodebase, err.Error())
		return err
	}
	rl.Info("end PutBranchInGit method...")
	return handler.NextServeOrNil(h.Next, cb)
}

func (h PutBranchInGit) setIntermediateSuccessFields(cb *codebaseApi.CodebaseBranch, action codebaseApi.ActionType) error {
	cb.Status = codebaseApi.CodebaseBranchStatus{
		Status:              model.StatusInit,
		LastTimeUpdated:     metaV1.Now(),
		Action:              action,
		Result:              codebaseApi.Success,
		Username:            "system",
		Value:               "inactive",
		VersionHistory:      cb.Status.VersionHistory,
		LastSuccessfulBuild: cb.Status.LastSuccessfulBuild,
		Build:               cb.Status.Build,
	}

	if err := h.Client.Status().Update(context.TODO(), cb); err != nil {
		if err := h.Client.Update(context.TODO(), cb); err != nil {
			return err
		}
	}
	return nil
}

func setFailedFields(cb *codebaseApi.CodebaseBranch, a codebaseApi.ActionType, message string) {
	cb.Status = codebaseApi.CodebaseBranchStatus{
		Status:              util.StatusFailed,
		LastTimeUpdated:     metaV1.Now(),
		Username:            "system",
		Action:              a,
		Result:              codebaseApi.Error,
		DetailedMessage:     message,
		Value:               "failed",
		VersionHistory:      cb.Status.VersionHistory,
		LastSuccessfulBuild: cb.Status.LastSuccessfulBuild,
		Build:               cb.Status.Build,
	}
}

func checkDirectory(path string) bool {
	return util.DoesDirectoryExist(path) && !util.IsDirectoryEmpty(path)
}

func (h PutBranchInGit) processNewVersion(b *codebaseApi.CodebaseBranch) error {
	if err := h.Service.ResetBranchBuildCounter(b); err != nil {
		return err
	}

	if err := h.Service.ResetBranchSuccessBuildCounter(b); err != nil {
		return err
	}

	return h.Service.AppendVersionToTheHistorySlice(b)
}

func hasNewVersion(b *codebaseApi.CodebaseBranch) bool {
	return !util.SearchVersion(b.Status.VersionHistory, *b.Spec.Version)
}
