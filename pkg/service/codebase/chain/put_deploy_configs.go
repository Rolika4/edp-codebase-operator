package chain

import (
	"fmt"
	"github.com/epmd-edp/codebase-operator/v2/pkg/apis/edp/v1alpha1"
	edpv1alpha1 "github.com/epmd-edp/codebase-operator/v2/pkg/apis/edp/v1alpha1"
	git "github.com/epmd-edp/codebase-operator/v2/pkg/controller/gitserver"
	"github.com/epmd-edp/codebase-operator/v2/pkg/gerrit"
	"github.com/epmd-edp/codebase-operator/v2/pkg/openshift"
	"github.com/epmd-edp/codebase-operator/v2/pkg/service/codebase/chain/handler"
	"github.com/epmd-edp/codebase-operator/v2/pkg/service/codebase/template"
	"github.com/epmd-edp/codebase-operator/v2/pkg/util"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
)

type PutDeployConfigs struct {
	next      handler.CodebaseHandler
	clientSet openshift.ClientSet
}

func (h PutDeployConfigs) ServeRequest(c *v1alpha1.Codebase) error {
	rLog := log.WithValues("codebase name", c.Name)
	rLog.Info("Start pushing configs...")

	gs, _, err := util.GetConfigSettings(h.clientSet.CoreClient, c.Namespace)
	if err != nil {
		return errors.Wrap(err, "unable get config settings")
	}

	if err := h.tryToPushConfigs(*c, gs.SshPort); err != nil {
		setFailedFields(*c, edpv1alpha1.SetupDeploymentTemplates, err.Error())
		return errors.Wrapf(err, "couldn't push deploy configs", "codebase name", c.Name)
	}

	wd := fmt.Sprintf("/home/codebase-operator/edp/%v/%v", c.Namespace, c.Name)
	if err := util.RemoveDirectory(wd); err != nil {
		return err
	}
	log.Info("directory was cleaned", "path", wd)
	rLog.Info("end pushing configs")
	return nextServeOrNil(h.next, c)
}

func (h PutDeployConfigs) tryToPushConfigs(c edpv1alpha1.Codebase, sshPort int32) error {
	s, err := util.GetSecret(*h.clientSet.CoreClient, "gerrit-project-creator", c.Namespace)
	if err != nil {
		return errors.Wrap(err, "unable to get gerrit-project-creator secret")
	}
	idrsa := string(s.Data[util.PrivateSShKeyName])

	u := "project-creator"
	url := fmt.Sprintf("ssh://%v@gerrit.%v:%v/%v", u, c.Namespace, sshPort, c.Name)
	wd := fmt.Sprintf("/home/codebase-operator/edp/%v/%v", c.Namespace, c.Name)
	td := util.GetTemplateDirectory(wd, c.Spec.DeploymentScript)
	d := fmt.Sprintf("%v/%v", td, c.Name)

	if !util.DoesDirectoryExist(d) || util.IsDirectoryEmpty(d) {
		if err := cloneProjectRepoFromGerrit(sshPort, idrsa, c.Name, c.Namespace, url, td); err != nil {
			return err
		}
	}

	cf := fmt.Sprintf("%v/%v/config-files", td, c.Name)
	if err := util.CreateDirectory(cf); err != nil {
		return err
	}

	fn := "Readme.md"
	src := fmt.Sprintf("%v/%v", util.GerritTemplates, fn)
	dest := fmt.Sprintf("%v/%v", cf, fn)
	if err := util.CopyFile(src, dest); err != nil {
		return err
	}

	if err := template.PrepareTemplates(h.clientSet.CoreClient, c); err != nil {
		return err
	}

	if err := git.CommitChanges(d); err != nil {
		return err
	}

	if err := git.PushChanges(idrsa, u, d); err != nil {
		return err
	}
	return nil
}

func cloneProjectRepoFromGerrit(sshPort int32, idrsa, name, namespace, cloneSshUrl, td string) error {
	log.Info("start cloning repository from Gerrit", "ssh url", cloneSshUrl)

	var (
		s *ssh.Session
		c *ssh.Client

		h = fmt.Sprintf("gerrit.%v", namespace)
	)

	sshcl, err := gerrit.SshInit(sshPort, idrsa, h)
	if err != nil {
		return errors.Wrap(err, "couldn't initialize SSH client")
	}

	if s, c, err = sshcl.NewSession(); err != nil {
		return errors.Wrap(err, "couldn't initialize SSH session")
	}

	defer func() {
		if deferErr := s.Close(); deferErr != nil {
			err = deferErr
		}
		if deferErr := c.Close(); deferErr != nil {
			err = deferErr
		}
	}()

	d := fmt.Sprintf("%v/%v", td, name)
	if err := git.CloneRepositoryBySsh(idrsa, "project-creator", cloneSshUrl, d); err != nil {
		return err
	}

	destinationPath := fmt.Sprintf("%v/%v/.git/hooks", td, name)
	if err := util.CreateDirectory(destinationPath); err != nil {
		return errors.Wrapf(err, "couldn't create folder %v", destinationPath)
	}

	sourcePath := "/usr/local/bin/configs"
	fileName := "commit-msg"
	src := fmt.Sprintf("%v/%v", sourcePath, fileName)
	dest := fmt.Sprintf("%v/%v", destinationPath, fileName)
	if err := util.CopyFile(src, dest); err != nil {
		return errors.Wrapf(err, "couldn't copy file %v", fileName)
	}
	return nil
}
