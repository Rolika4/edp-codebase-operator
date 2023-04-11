package chain

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	jenkinsApi "github.com/epam/edp-jenkins-operator/v2/pkg/apis/v2/v1"
	"github.com/epam/edp-jenkins-operator/v2/pkg/util/platform"

	codebaseApi "github.com/epam/edp-codebase-operator/v2/api/v1"
	"github.com/epam/edp-codebase-operator/v2/pkg/util"
)

type PutCDStageJenkinsDeployment struct {
	client client.Client
	log    logr.Logger
}

const (
	cdPipelinePostfix = "-cd-pipeline"
	jenkinsKey        = "jenkinsName"
	cdStageDeployKey  = "cdStageDeployName"
)

func (h *PutCDStageJenkinsDeployment) ServeRequest(ctx context.Context, stageDeploy *codebaseApi.CDStageDeploy) error {
	log := ctrl.LoggerFrom(ctx)
	log.Info("Creating CDStageJenkinsDeployment")

	jd, err := h.getCDStageJenkinsDeployment(stageDeploy.Name, stageDeploy.Namespace)
	if err != nil {
		return fmt.Errorf("failed to get %v cd stage jenkins deployment: %w", stageDeploy.Name, err)
	}

	if jd != nil {
		h.log.Info(fmt.Sprintf("CDStageJenkinsDeployment %s already exists with status %s and status message %s",
			stageDeploy.Name, jd.Status.Status, jd.Status.Message))

		return fmt.Errorf("failed to create CDStageJenkinsDeployment: %w", ErrCDStageJenkinsDeploymentHasNotBeenProcessed)
	}

	if err := h.create(stageDeploy); err != nil {
		return fmt.Errorf("failed to create %v cd stage jenkins deployment: %w", stageDeploy.Name, err)
	}

	log.Info("creating CDStageJenkinsDeployment has been finished.")

	return nil
}

func (h *PutCDStageJenkinsDeployment) getCDStageJenkinsDeployment(name, namespace string) (*jenkinsApi.CDStageJenkinsDeployment, error) {
	h.log.Info("getting cd stage jenkins deployment", "stageDeployment", name)

	ctx := context.Background()
	cdStageJenkinsDeployment := &jenkinsApi.CDStageJenkinsDeployment{}

	if err := h.client.Get(ctx,
		types.NamespacedName{
			Namespace: namespace,
			Name:      name,
		}, cdStageJenkinsDeployment,
	); err != nil {
		if k8sErrors.IsNotFound(err) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to fetch 'CDStageJenkinsDeployment' resource %q: %w", name, err)
	}

	return cdStageJenkinsDeployment, nil
}

func (h *PutCDStageJenkinsDeployment) create(stageDeploy *codebaseApi.CDStageDeploy) error {
	log := h.log.WithValues("name", stageDeploy.Name)
	log.Info("cd stage jenkins deployment is not present in cluster. start creating...")

	ctx := context.Background()

	labels, err := h.generateLabels(stageDeploy.Name, stageDeploy.Namespace)
	if err != nil {
		return fmt.Errorf("failed to generate labels: %w", err)
	}

	tagsList := make([]jenkinsApi.Tag, 0)
	for _, codebaseTag := range stageDeploy.Spec.Tags {
		tagsList = append(tagsList, jenkinsApi.Tag{
			Codebase: codebaseTag.Codebase,
			Tag:      codebaseTag.Tag,
		})
	}

	jdCommand := &jenkinsApi.CDStageJenkinsDeployment{
		TypeMeta: metaV1.TypeMeta{
			APIVersion: util.V2APIVersion,
			Kind:       util.CDStageJenkinsDeploymentKind,
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name:      stageDeploy.Name,
			Namespace: stageDeploy.Namespace,
			Labels:    labels,
		},
		Spec: jenkinsApi.CDStageJenkinsDeploymentSpec{
			Job: fmt.Sprintf("%v%v/job/%v", stageDeploy.Spec.Pipeline, cdPipelinePostfix, stageDeploy.Spec.Stage),
			Tag: jenkinsApi.Tag{
				Codebase: stageDeploy.Spec.Tag.Codebase,
				Tag:      stageDeploy.Spec.Tag.Tag,
			},
			Tags: tagsList,
		},
	}

	if err = h.client.Create(ctx, jdCommand); err != nil {
		return fmt.Errorf("failed to create CDStageJenkinsDeployment resource: %w", err)
	}

	log.Info("cd stage jenkins deployment has been created.")

	return nil
}

func (h *PutCDStageJenkinsDeployment) generateLabels(cdStageDeployName, ns string) (map[string]string, error) {
	ji, err := platform.GetFirstJenkinsInstance(h.client, ns)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch jenkins instance from cluster: %w", err)
	}

	return map[string]string{
		jenkinsKey:       ji.Name,
		cdStageDeployKey: cdStageDeployName,
	}, nil
}
