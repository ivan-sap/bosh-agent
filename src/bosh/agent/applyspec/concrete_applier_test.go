package applyspec

import (
	fakebc "bosh/agent/applyspec/bundlecollection/fakes"
	models "bosh/agent/applyspec/models"
	fakepa "bosh/agent/applyspec/packageapplier/fakes"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApplyInstallsAndEnablesJobs(t *testing.T) {
	jobsBc, _, applier := buildApplier()
	job := buildJob()

	err := applier.Apply([]models.Job{job}, []models.Package{})
	assert.NoError(t, err)
	assert.True(t, jobsBc.IsInstalled(job))
	assert.True(t, jobsBc.IsEnabled(job))
}

func TestApplyErrsWhenJobInstallFails(t *testing.T) {
	jobsBc, _, applier := buildApplier()
	job := buildJob()

	jobsBc.InstallError = errors.New("fake-install-error")

	err := applier.Apply([]models.Job{job}, []models.Package{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "fake-install-error")
}

func TestApplyErrsWhenJobEnableFails(t *testing.T) {
	jobsBc, _, applier := buildApplier()
	job := buildJob()

	jobsBc.EnableError = errors.New("fake-enable-error")

	err := applier.Apply([]models.Job{job}, []models.Package{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "fake-enable-error")
}

func TestApplyAppliesPackages(t *testing.T) {
	_, packageApplier, applier := buildApplier()

	pkg1 := buildPackage()
	pkg2 := buildPackage()

	err := applier.Apply([]models.Job{}, []models.Package{pkg1, pkg2})
	assert.NoError(t, err)
	assert.Equal(t, packageApplier.AppliedPackages, []models.Package{pkg1, pkg2})
}

func TestApplyErrsWhenApplyingPackagesErrs(t *testing.T) {
	_, packageApplier, applier := buildApplier()
	pkg := buildPackage()

	packageApplier.ApplyError = errors.New("fake-apply-error")

	err := applier.Apply([]models.Job{}, []models.Package{pkg})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "fake-apply-error")
}

func buildApplier() (
	*fakebc.FakeBundleCollection,
	*fakepa.FakePackageApplier,
	Applier,
) {
	jobsBc := fakebc.NewFakeBundleCollection()
	packageApplier := fakepa.NewFakePackageApplier()
	applier := NewConcreteApplier(jobsBc, packageApplier)
	return jobsBc, packageApplier, applier
}

func buildJob() models.Job {
	return models.Job{Name: "fake-job-name", Version: "fake-version-name"}
}

func buildPackage() models.Package {
	return models.Package{Name: "fake-package-name", Version: "fake-package-name"}
}