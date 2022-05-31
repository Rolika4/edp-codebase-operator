package chain

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	codebaseApi "github.com/epam/edp-codebase-operator/v2/pkg/apis/edp/v1"
	mockGit "github.com/epam/edp-codebase-operator/v2/pkg/controller/gitserver/mock"
	"github.com/epam/edp-codebase-operator/v2/pkg/util"
)

func TestGetRepositoryCredentialsIfExists_ShouldPass(t *testing.T) {
	c := &codebaseApi.Codebase{
		ObjectMeta: metaV1.ObjectMeta{
			Name:      "fake-name",
			Namespace: fakeNamespace,
		},
		Spec: codebaseApi.CodebaseSpec{
			Repository: &codebaseApi.Repository{
				Url: "repo",
			},
		},
	}
	s := &coreV1.Secret{
		ObjectMeta: metaV1.ObjectMeta{
			Name:      "repository-codebase-fake-name-temp",
			Namespace: fakeNamespace,
		},
		Data: map[string][]byte{
			"username": []byte("user"),
			"password": []byte("pass"),
		},
	}
	scheme := runtime.NewScheme()
	scheme.AddKnownTypes(coreV1.SchemeGroupVersion, s)
	scheme.AddKnownTypes(codebaseApi.SchemeGroupVersion, c)
	fakeCl := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(s, c).Build()
	u, p, err := GetRepositoryCredentialsIfExists(c, fakeCl)
	assert.Equal(t, u, util.GetStringP("user"))
	assert.Equal(t, p, util.GetStringP("pass"))
	assert.NoError(t, err)
}

func TestGetRepositoryCredentialsIfExists_ShouldFail(t *testing.T) {
	c := &codebaseApi.Codebase{
		ObjectMeta: metaV1.ObjectMeta{
			Name:      "fake-name",
			Namespace: fakeNamespace,
		},
		Spec: codebaseApi.CodebaseSpec{
			Repository: &codebaseApi.Repository{
				Url: "repo",
			},
		},
	}

	scheme := runtime.NewScheme()
	scheme.AddKnownTypes(codebaseApi.SchemeGroupVersion, c)
	fakeCl := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(c).Build()
	_, _, err := GetRepositoryCredentialsIfExists(c, fakeCl)
	assert.Error(t, err)
	if !strings.Contains(err.Error(), "Unable to get secret repository-codebase-fake-name-temp") {
		t.Fatalf("wrong error returned: %s", err.Error())
	}
}

func TestCheckoutBranch_ShouldFailOnGetSecret(t *testing.T) {
	c := &codebaseApi.Codebase{
		ObjectMeta: metaV1.ObjectMeta{
			Name:      "fake-name",
			Namespace: fakeNamespace,
		},
		Spec: codebaseApi.CodebaseSpec{
			Repository: &codebaseApi.Repository{
				Url: "repo",
			},
		},
	}

	scheme := runtime.NewScheme()
	scheme.AddKnownTypes(codebaseApi.SchemeGroupVersion, c)
	fakeCl := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(c).Build()

	mGit := new(mockGit.MockGit)

	err := CheckoutBranch(util.GetStringP("repo"), "project-path", "branch", mGit, c, fakeCl)
	assert.Error(t, err)
	if !strings.Contains(err.Error(), "Unable to get secret repository-codebase-fake-name-temp") {
		t.Fatalf("wrong error returned: %s", err.Error())
	}
}

func TestCheckoutBranch_ShouldFailOnCheckPermission(t *testing.T) {
	c := &codebaseApi.Codebase{
		ObjectMeta: metaV1.ObjectMeta{
			Name:      "fake-name",
			Namespace: fakeNamespace,
		},
		Spec: codebaseApi.CodebaseSpec{
			Repository: &codebaseApi.Repository{
				Url: "repo",
			},
		},
	}
	s := &coreV1.Secret{
		ObjectMeta: metaV1.ObjectMeta{
			Name:      "repository-codebase-fake-name-temp",
			Namespace: fakeNamespace,
		},
		Data: map[string][]byte{
			"username": []byte("user"),
			"password": []byte("pass"),
		},
	}
	scheme := runtime.NewScheme()
	scheme.AddKnownTypes(coreV1.SchemeGroupVersion, s)
	scheme.AddKnownTypes(codebaseApi.SchemeGroupVersion, c)
	fakeCl := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(s, c).Build()

	mGit := new(mockGit.MockGit)
	mGit.On("CheckPermissions", "repo", util.GetStringP("user"), util.GetStringP("pass")).Return(false)

	err := CheckoutBranch(util.GetStringP("repo"), "project-path", "branch", mGit, c, fakeCl)
	assert.Error(t, err)
	if !strings.Contains(err.Error(), "user user cannot get access to the repository repo") {
		t.Fatalf("wrong error returned: %s", err.Error())
	}
}

func TestCheckoutBranch_ShouldFailOnGetCurrentBranchName(t *testing.T) {
	c := &codebaseApi.Codebase{
		ObjectMeta: metaV1.ObjectMeta{
			Name:      "fake-name",
			Namespace: fakeNamespace,
		},
		Spec: codebaseApi.CodebaseSpec{
			Repository: &codebaseApi.Repository{
				Url: "repo",
			},
		},
	}
	s := &coreV1.Secret{
		ObjectMeta: metaV1.ObjectMeta{
			Name:      "repository-codebase-fake-name-temp",
			Namespace: fakeNamespace,
		},
		Data: map[string][]byte{
			"username": []byte("user"),
			"password": []byte("pass"),
		},
	}
	scheme := runtime.NewScheme()
	scheme.AddKnownTypes(coreV1.SchemeGroupVersion, s)
	scheme.AddKnownTypes(codebaseApi.SchemeGroupVersion, c)
	fakeCl := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(s, c).Build()

	mGit := new(mockGit.MockGit)
	mGit.On("CheckPermissions", "repo", util.GetStringP("user"), util.GetStringP("pass")).Return(true)
	mGit.On("GetCurrentBranchName", "project-path").Return("", errors.New("FATAL:FAILED"))

	err := CheckoutBranch(util.GetStringP("repo"), "project-path", "branch", mGit, c, fakeCl)
	assert.Error(t, err)
	if !strings.Contains(err.Error(), "FATAL:FAILED") {
		t.Fatalf("wrong error returned: %s", err.Error())
	}
}

func TestCheckoutBranch_ShouldFailOnCheckout(t *testing.T) {
	var (
		repo = "repo"
		u    = "user"
		p    = "pass"
	)
	c := &codebaseApi.Codebase{
		ObjectMeta: metaV1.ObjectMeta{
			Name:      "fake-name",
			Namespace: fakeNamespace,
		},
		Spec: codebaseApi.CodebaseSpec{
			Repository: &codebaseApi.Repository{
				Url: "repo",
			},
			Strategy: codebaseApi.Clone,
		},
	}
	s := &coreV1.Secret{
		ObjectMeta: metaV1.ObjectMeta{
			Name:      "repository-codebase-fake-name-temp",
			Namespace: fakeNamespace,
		},
		Data: map[string][]byte{
			"username": []byte("user"),
			"password": []byte("pass"),
		},
	}
	scheme := runtime.NewScheme()
	scheme.AddKnownTypes(coreV1.SchemeGroupVersion, s)
	scheme.AddKnownTypes(codebaseApi.SchemeGroupVersion, c)
	fakeCl := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(s, c).Build()

	mGit := new(mockGit.MockGit)
	mGit.On("CheckPermissions", "repo", &u, &p).Return(true)
	mGit.On("GetCurrentBranchName", "project-path").Return("some-other-branch", nil)
	mGit.On("Checkout", &u, &p, "project-path", "branch", true).Return(errors.New("FATAL:FAILED"))

	err := CheckoutBranch(&repo, "project-path", "branch", mGit, c, fakeCl)
	assert.Error(t, err)
	if !strings.Contains(err.Error(), "FATAL:FAILED") {
		t.Fatalf("wrong error returned: %s", err.Error())
	}
}

func TestCheckoutBranch_ShouldPassForCloneStrategy(t *testing.T) {
	var (
		repo = "repo"
		u    = "user"
		p    = "pass"
	)
	c := &codebaseApi.Codebase{
		ObjectMeta: metaV1.ObjectMeta{
			Name:      "fake-name",
			Namespace: fakeNamespace,
		},
		Spec: codebaseApi.CodebaseSpec{
			GitServer: "git",
			Repository: &codebaseApi.Repository{
				Url: "repo",
			},
			Strategy: codebaseApi.Import,
		},
	}
	gs := &codebaseApi.GitServer{
		ObjectMeta: metaV1.ObjectMeta{
			Name:      "git",
			Namespace: fakeNamespace,
		},
		Spec: codebaseApi.GitServerSpec{
			NameSshKeySecret: fakeName,
			GitHost:          fakeName,
			SshPort:          22,
			GitUser:          fakeName,
		},
	}
	s := &coreV1.Secret{
		ObjectMeta: metaV1.ObjectMeta{
			Name:      "repository-codebase-fake-name-temp",
			Namespace: fakeNamespace,
		},
		Data: map[string][]byte{
			"username": []byte("user"),
			"password": []byte("pass"),
		},
	}
	ssh := &coreV1.Secret{
		ObjectMeta: metaV1.ObjectMeta{
			Name:      fakeName,
			Namespace: fakeNamespace,
		},
		Data: map[string][]byte{
			util.PrivateSShKeyName: []byte("fake"),
		},
	}
	scheme := runtime.NewScheme()
	scheme.AddKnownTypes(coreV1.SchemeGroupVersion, s, ssh)
	scheme.AddKnownTypes(codebaseApi.SchemeGroupVersion, c, gs)
	fakeCl := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(s, c, gs, ssh).Build()

	mGit := new(mockGit.MockGit)
	mGit.On("CheckPermissions", "repo", &u, &p).Return(true)
	mGit.On("GetCurrentBranchName", "project-path").Return("some-other-branch", nil)
	mGit.On("CheckoutRemoteBranchBySSH", "fake", fakeName, "project-path", "branch").Return(nil)

	err := CheckoutBranch(&repo, "project-path", "branch", mGit, c, fakeCl)
	assert.NoError(t, err)
}
