package webhook

import (
	"context"
	"testing"

	"github.com/go-logr/logr"
	"github.com/stretchr/testify/require"
	"k8s.io/api/admission/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	codebaseApi "github.com/epam/edp-codebase-operator/v2/api/v1"
	"github.com/epam/edp-codebase-operator/v2/pkg/util"
)

func TestCodebaseWebhook_ValidateCreate(t *testing.T) {
	t.Parallel()

	scheme := runtime.NewScheme()
	require.NoError(t, codebaseApi.AddToScheme(scheme))

	tests := []struct {
		name    string
		client  client.Client
		ctx     context.Context
		obj     runtime.Object
		wantErr require.ErrorAssertionFunc
	}{
		{
			name: "should return error if GitUrlPath already exists",
			client: fake.NewClientBuilder().WithScheme(scheme).WithObjects(&codebaseApi.Codebase{
				ObjectMeta: metaV1.ObjectMeta{
					Name:      "codebase2",
					Namespace: "default",
				},
				Spec: codebaseApi.CodebaseSpec{
					GitUrlPath: util.GetStringP("user/repo"),
				},
			}).Build(),
			ctx: admission.NewContextWithRequest(context.Background(), admission.Request{
				AdmissionRequest: v1.AdmissionRequest{
					Name:      "codebase",
					Namespace: "default",
				},
			}),
			obj: &codebaseApi.Codebase{
				ObjectMeta: metaV1.ObjectMeta{
					Name:      "codebase",
					Namespace: "default",
				},
				Spec: codebaseApi.CodebaseSpec{
					GitUrlPath: util.GetStringP("user/repo"),
				},
			},
			wantErr: func(t require.TestingT, err error, _ ...any) {
				require.Error(t, err)

				require.Contains(t, err.Error(), "codebase codebase2 with GitUrlPath user/repo already exists")
			},
		},
		{
			name: "should return error if GitUrlPath already exists with, check .git suffix",
			client: fake.NewClientBuilder().WithScheme(scheme).WithObjects(&codebaseApi.Codebase{
				ObjectMeta: metaV1.ObjectMeta{
					Name:      "codebase2",
					Namespace: "default",
				},
				Spec: codebaseApi.CodebaseSpec{
					GitUrlPath: util.GetStringP("user/repo"),
				},
			}).Build(),
			ctx: admission.NewContextWithRequest(context.Background(), admission.Request{
				AdmissionRequest: v1.AdmissionRequest{
					Name:      "codebase",
					Namespace: "default",
				},
			}),
			obj: &codebaseApi.Codebase{
				ObjectMeta: metaV1.ObjectMeta{
					Name:      "codebase",
					Namespace: "default",
				},
				Spec: codebaseApi.CodebaseSpec{
					GitUrlPath: util.GetStringP("user/repo.git"),
				},
			},
			wantErr: func(t require.TestingT, err error, _ ...any) {
				require.Error(t, err)

				require.Contains(t, err.Error(), "codebase codebase2 with GitUrlPath user/repo already exists")
			},
		},
		{
			name: "should skip codebase with same GitUrlPath in the another namespace",
			client: fake.NewClientBuilder().WithScheme(scheme).WithObjects(&codebaseApi.Codebase{
				ObjectMeta: metaV1.ObjectMeta{
					Name:      "codebase2",
					Namespace: "prod",
				},
				Spec: codebaseApi.CodebaseSpec{
					GitUrlPath: util.GetStringP("user/repo"),
				},
			}).Build(),
			ctx: admission.NewContextWithRequest(context.Background(), admission.Request{
				AdmissionRequest: v1.AdmissionRequest{
					Name:      "codebase",
					Namespace: "default",
				},
			}),
			obj: &codebaseApi.Codebase{
				ObjectMeta: metaV1.ObjectMeta{
					Name:      "codebase",
					Namespace: "default",
				},
				Spec: codebaseApi.CodebaseSpec{
					GitUrlPath: util.GetStringP("user/repo"),
				},
			},
			wantErr: require.NoError,
		},
		{
			name:   "invalid admission.Request",
			client: fake.NewClientBuilder().WithScheme(scheme).Build(),
			ctx:    context.Background(),
			wantErr: func(t require.TestingT, err error, _ ...any) {
				require.Error(t, err)

				require.Contains(t, err.Error(), "expected admission.Request in ctx")
			},
		},
		{
			name:   "skip validation - empty GitUrlPath",
			client: fake.NewClientBuilder().WithScheme(scheme).Build(),
			ctx: admission.NewContextWithRequest(context.Background(), admission.Request{
				AdmissionRequest: v1.AdmissionRequest{
					Name:      "codebase",
					Namespace: "default",
				},
			}),
			obj: &codebaseApi.Codebase{
				ObjectMeta: metaV1.ObjectMeta{
					Name:      "codebase",
					Namespace: "default",
				},
			},
			wantErr: require.NoError,
		},
		{
			name:   "should return error if GitUrlPath is not valid",
			client: fake.NewClientBuilder().WithScheme(scheme).Build(),
			ctx: admission.NewContextWithRequest(context.Background(), admission.Request{
				AdmissionRequest: v1.AdmissionRequest{
					Name:      "codebase",
					Namespace: "default",
				},
			}),
			obj: &codebaseApi.Codebase{
				ObjectMeta: metaV1.ObjectMeta{
					Name:      "codebase",
					Namespace: "default",
				},
				Spec: codebaseApi.CodebaseSpec{
					GitUrlPath: util.GetStringP(".git"),
				},
			},
			wantErr: func(t require.TestingT, err error, _ ...any) {
				require.Error(t, err)

				require.Contains(t, err.Error(), "gitUrlPath .git is invalid")
			},
		},
		{
			name:   "skip validation - wrong object",
			client: fake.NewClientBuilder().WithScheme(scheme).Build(),
			ctx: admission.NewContextWithRequest(context.Background(), admission.Request{
				AdmissionRequest: v1.AdmissionRequest{
					Name:      "codebase",
					Namespace: "default",
				},
			}),
			obj: &codebaseApi.CodebaseBranch{
				ObjectMeta: metaV1.ObjectMeta{
					Name:      "branch",
					Namespace: "default",
				},
			},
			wantErr: require.NoError,
		},
		{
			name: "nil GitUrlPath in existing codebase",
			client: fake.NewClientBuilder().WithScheme(scheme).WithObjects(&codebaseApi.Codebase{
				ObjectMeta: metaV1.ObjectMeta{
					Name:      "codebase2",
					Namespace: "default",
				},
			}).Build(),
			ctx: admission.NewContextWithRequest(context.Background(), admission.Request{
				AdmissionRequest: v1.AdmissionRequest{
					Name:      "codebase",
					Namespace: "default",
				},
			}),
			obj: &codebaseApi.Codebase{
				ObjectMeta: metaV1.ObjectMeta{
					Name:      "codebase",
					Namespace: "default",
				},
				Spec: codebaseApi.CodebaseSpec{
					GitUrlPath: util.GetStringP("user/repo"),
				},
			},
			wantErr: require.NoError,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r := NewCodebaseValidationWebhook(tt.client, logr.Discard())

			err := r.ValidateCreate(tt.ctx, tt.obj)
			tt.wantErr(t, err)
		})
	}
}

func TestCodebaseValidationWebhook_ValidateUpdate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		ctx     context.Context
		newObj  runtime.Object
		oldObj  runtime.Object
		wantErr require.ErrorAssertionFunc
	}{
		{
			name: "invalid admission.Request",
			ctx:  context.Background(),
			newObj: &codebaseApi.Codebase{
				ObjectMeta: metaV1.ObjectMeta{
					Name:      "codebase",
					Namespace: "default",
				},
			},
			oldObj: &codebaseApi.Codebase{
				ObjectMeta: metaV1.ObjectMeta{
					Name:      "codebase",
					Namespace: "default",
				},
			},
			wantErr: func(t require.TestingT, err error, _ ...any) {
				require.Error(t, err)

				require.Contains(t, err.Error(), "expected admission.Request in ctx")
			},
		},
		{
			name: "skip validation",
			ctx: admission.NewContextWithRequest(context.Background(), admission.Request{
				AdmissionRequest: v1.AdmissionRequest{
					Name:      "codebase",
					Namespace: "default",
				},
			}),
			newObj: &codebaseApi.Codebase{
				ObjectMeta: metaV1.ObjectMeta{
					Name:      "codebase",
					Namespace: "default",
				},
			},
			oldObj: &codebaseApi.Codebase{
				ObjectMeta: metaV1.ObjectMeta{
					Name:      "codebase",
					Namespace: "default",
				},
			},
			wantErr: require.NoError,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r := NewCodebaseValidationWebhook(fake.NewClientBuilder().Build(), logr.Discard())

			err := r.ValidateUpdate(tt.ctx, tt.oldObj, tt.newObj)
			tt.wantErr(t, err)
		})
	}
}

func TestCodebaseValidationWebhook_ValidateDelete(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		ctx     context.Context
		obj     runtime.Object
		wantErr require.ErrorAssertionFunc
	}{
		{
			name: "invalid admission.Request",
			ctx:  context.Background(),
			obj: &codebaseApi.Codebase{
				ObjectMeta: metaV1.ObjectMeta{
					Name:      "codebase",
					Namespace: "default",
				},
			},
			wantErr: func(t require.TestingT, err error, _ ...any) {
				require.Error(t, err)

				require.Contains(t, err.Error(), "expected admission.Request in ctx")
			},
		},
		{
			name: "skip validation",
			ctx: admission.NewContextWithRequest(context.Background(), admission.Request{
				AdmissionRequest: v1.AdmissionRequest{
					Name:      "codebase",
					Namespace: "default",
				},
			}),
			obj: &codebaseApi.Codebase{
				ObjectMeta: metaV1.ObjectMeta{
					Name:      "codebase",
					Namespace: "default",
				},
			},
			wantErr: require.NoError,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r := NewCodebaseValidationWebhook(fake.NewClientBuilder().Build(), logr.Discard())

			err := r.ValidateDelete(tt.ctx, tt.obj)
			tt.wantErr(t, err)
		})
	}
}
