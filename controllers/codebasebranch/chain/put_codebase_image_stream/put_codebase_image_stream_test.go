package put_codebase_image_stream

import (
	"context"
	"testing"

	"github.com/go-logr/logr"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	edpComponentApi "github.com/epam/edp-component-operator/api/v1"

	codebaseApi "github.com/epam/edp-codebase-operator/v2/api/v1"
	"github.com/epam/edp-codebase-operator/v2/pkg/util"
)

func TestPutCodebaseImageStream_ServeRequest(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		codebaseBranch *codebaseApi.CodebaseBranch
		objects        []client.Object
		wantErr        require.ErrorAssertionFunc
	}{
		{
			name: "successfully put codebase image stream - get docker registry url from config map",
			codebaseBranch: &codebaseApi.CodebaseBranch{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-branch",
					Namespace: "default",
				},
				Spec: codebaseApi.CodebaseBranchSpec{
					CodebaseName: "test-codebase",
					BranchName:   "test-branch-master",
				},
			},
			objects: []client.Object{
				&codebaseApi.Codebase{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-codebase",
						Namespace: "default",
					},
				},
				&corev1.ConfigMap{
					ObjectMeta: metav1.ObjectMeta{
						Name:      util.EdpConfigMap,
						Namespace: "default",
					},
					Data: map[string]string{
						edpConfigContainerRegistryHost:  "test-registry",
						edpConfigContainerRegistrySpace: "test-space",
					},
				},
			},
			wantErr: require.NoError,
		},
		{
			name: "successfully put codebase image stream - get docker registry url from EDPComponent",
			codebaseBranch: &codebaseApi.CodebaseBranch{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-branch",
					Namespace: "default",
				},
				Spec: codebaseApi.CodebaseBranchSpec{
					CodebaseName: "test-codebase",
					BranchName:   "test-branch-master",
				},
			},
			objects: []client.Object{
				&codebaseApi.Codebase{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-codebase",
						Namespace: "default",
					},
				},
				&edpComponentApi.EDPComponent{
					ObjectMeta: metav1.ObjectMeta{
						Name:      dockerRegistryName,
						Namespace: "default",
					},
					Spec: edpComponentApi.EDPComponentSpec{
						Url: "test-registry/test-space",
					},
				},
			},
			wantErr: require.NoError,
		},
		{
			name: "codebase image stream already exists",
			codebaseBranch: &codebaseApi.CodebaseBranch{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-branch",
					Namespace: "default",
				},
				Spec: codebaseApi.CodebaseBranchSpec{
					CodebaseName: "test-codebase",
					BranchName:   "test-branch-master",
				},
			},
			objects: []client.Object{
				&codebaseApi.Codebase{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-codebase",
						Namespace: "default",
					},
				},
				&codebaseApi.CodebaseImageStream{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-codebase-test-branch-master",
						Namespace: "default",
					},
				},
				&corev1.ConfigMap{
					ObjectMeta: metav1.ObjectMeta{
						Name:      util.EdpConfigMap,
						Namespace: "default",
					},
					Data: map[string]string{
						edpConfigContainerRegistryHost:  "test-registry",
						edpConfigContainerRegistrySpace: "test-space",
					},
				},
			},
			wantErr: require.NoError,
		},
		{
			name: "failed to get registry url",
			codebaseBranch: &codebaseApi.CodebaseBranch{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-branch",
					Namespace: "default",
				},
				Spec: codebaseApi.CodebaseBranchSpec{
					CodebaseName: "test-codebase",
					BranchName:   "test-branch-master",
				},
			},
			objects: []client.Object{
				&codebaseApi.Codebase{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-codebase",
						Namespace: "default",
					},
				},
			},
			wantErr: func(t require.TestingT, err error, i ...interface{}) {
				require.Error(t, err)
				require.Contains(t, err.Error(), "failed to get container registry url")
			},
		},
		{
			name: "failed to get codebase",
			codebaseBranch: &codebaseApi.CodebaseBranch{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-branch",
					Namespace: "default",
				},
				Spec: codebaseApi.CodebaseBranchSpec{
					CodebaseName: "test-codebase",
					BranchName:   "test-branch-master",
				},
			},
			objects: []client.Object{},
			wantErr: func(t require.TestingT, err error, i ...interface{}) {
				require.Error(t, err)
				require.Contains(t, err.Error(), "failed to fetch Codebase resource")
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			scheme := runtime.NewScheme()
			require.NoError(t, codebaseApi.AddToScheme(scheme))
			require.NoError(t, edpComponentApi.AddToScheme(scheme))
			require.NoError(t, corev1.AddToScheme(scheme))

			h := PutCodebaseImageStream{
				Client: fake.NewClientBuilder().WithScheme(scheme).WithObjects(append(tt.objects, tt.codebaseBranch)...).Build(),
			}

			err := h.ServeRequest(ctrl.LoggerInto(context.Background(), logr.Discard()), tt.codebaseBranch)
			tt.wantErr(t, err)
		})
	}
}
