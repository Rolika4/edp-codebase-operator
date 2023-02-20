package gittag

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	predicateLib "github.com/operator-framework/operator-lib/predicate"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	codebaseApi "github.com/epam/edp-codebase-operator/v2/api/v1"
	"github.com/epam/edp-codebase-operator/v2/controllers/gittag/chain"
	"github.com/epam/edp-codebase-operator/v2/pkg/util"
)

func NewReconcileGitTag(c client.Client, log logr.Logger) *ReconcileGitTag {
	return &ReconcileGitTag{
		client: c,
		log:    log.WithName("git-tag"),
	}
}

type ReconcileGitTag struct {
	client client.Client
	log    logr.Logger
}

func (r *ReconcileGitTag) SetupWithManager(mgr ctrl.Manager) error {
	p := predicate.Funcs{
		UpdateFunc: func(e event.UpdateEvent) bool {
			oldObject, ok := e.ObjectOld.(*codebaseApi.GitTag)
			if !ok {
				return false
			}

			newObject, ok := e.ObjectNew.(*codebaseApi.GitTag)
			if !ok {
				return false
			}

			return oldObject.Status == newObject.Status
		},
	}

	pause, err := predicateLib.NewPause(util.PauseAnnotation)
	if err != nil {
		return fmt.Errorf("failed to create pause predicate: %w", err)
	}

	err = ctrl.NewControllerManagedBy(mgr).
		For(&codebaseApi.GitTag{}, builder.WithPredicates(p, pause)).
		Complete(r)
	if err != nil {
		return fmt.Errorf("failed to build GitTag controller: %w", err)
	}

	return nil
}

//+kubebuilder:rbac:groups=v2.edp.epam.com,namespace=placeholder,resources=gittags,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=v2.edp.epam.com,namespace=placeholder,resources=gittags/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=v2.edp.epam.com,namespace=placeholder,resources=gittags/finalizers,verbs=update

// Reconcile reads that state of the cluster for a GitTag object and makes changes based on the state.
func (r *ReconcileGitTag) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	log := r.log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	log.Info("Reconciling GitTag")

	gt := &codebaseApi.GitTag{}
	if err := r.client.Get(ctx, request.NamespacedName, gt); err != nil {
		if k8sErrors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}

		return reconcile.Result{}, fmt.Errorf("failed to fetch GitTag resource %q: %w", request.NamespacedName, err)
	}

	gtChain := chain.CreateDefChain(r.client)
	if err := gtChain.ServeRequest(gt); err != nil {
		log.Error(err, err.Error())

		return reconcile.Result{}, fmt.Errorf("failed to process `Default chain`: %w", err)
	}

	log.Info("Reconciling GitTag has been finished")

	return reconcile.Result{}, nil
}