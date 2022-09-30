package template

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/pkg/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	codebaseApi "github.com/epam/edp-codebase-operator/v2/pkg/apis/edp/v1"
	"github.com/epam/edp-codebase-operator/v2/pkg/controller/platform"
	"github.com/epam/edp-codebase-operator/v2/pkg/model"
	"github.com/epam/edp-codebase-operator/v2/pkg/util"
)

var log = ctrl.Log.WithName("template")

func PrepareTemplates(c client.Client, cb *codebaseApi.Codebase, workDir, assetsDir string) error {
	log.Info("start preparing deploy templates", "codebase", cb.Name)

	cf, err := buildTemplateConfig(c, cb)
	if err != nil {
		return err
	}

	if cb.Spec.Type == util.Application {
		if err := util.CopyTemplate(cb.Spec.DeploymentScript, workDir, assetsDir, cf); err != nil {
			return errors.Wrapf(err, "an error has occurred while copying template for %v codebase", cb.Name)
		}
	}

	if cb.Spec.Strategy != util.ImportStrategy {
		if err := copySonarConfigs(workDir, assetsDir, cf); err != nil {
			return err
		}
	}
	log.Info("end preparing deploy templates", "codebase", cb.Name)
	return nil
}

func PrepareGitlabCITemplates(c client.Client, cb *codebaseApi.Codebase, workDir, assetsDir string) error {
	log.Info("start preparing deploy templates", "codebase", cb.Name)

	if cb.Spec.Type != util.Application {
		log.Info("codebase is not application. skip copying templates", "name", cb.Name)
		return nil
	}

	cf, err := buildTemplateConfig(c, cb)
	if err != nil {
		return err
	}

	if err := util.CopyTemplate(cb.Spec.DeploymentScript, workDir, assetsDir, cf); err != nil {
		return errors.Wrapf(err, "an error has occurred while copying template for %v codebase", cb.Name)
	}

	log.Info("end preparing deploy templates", "codebase", cb.Name)
	return nil
}

func buildTemplateConfig(c client.Client, cb *codebaseApi.Codebase) (*model.ConfigGoTemplating, error) {
	log.Info("start creating template config", "codebase_name", cb.Name)
	us, err := util.GetUserSettings(c, cb.Namespace)
	if err != nil {
		return nil, errors.Wrap(err, "unable get user settings settings")
	}

	cf := model.ConfigGoTemplating{
		Name:         cb.Name,
		PlatformType: platform.GetPlatformType(),
		Lang:         cb.Spec.Lang,
		DnsWildcard:  us.DnsWildcard,
	}
	if cb.Spec.Framework != nil {
		cf.Framework = *cb.Spec.Framework
	}
	cf.GitURL, err = getProjectUrl(c, &cb.Spec, cb.Namespace)
	if err != nil {
		return nil, errors.Wrap(err, "unable get project url")
	}

	log.Info("end creating template config", "codebase_name", cb.Name)
	return &cf, nil
}

func getProjectUrl(c client.Client, s *codebaseApi.CodebaseSpec, n string) (string, error) {
	switch s.Strategy {
	case "create":
		p := util.BuildRepoUrl(s)
		return p, nil

	case "clone":
		p := s.Repository.Url
		return p, nil

	case "import":
		gs, err := util.GetGitServer(c, s.GitServer, n)
		if err != nil {
			return "", errors.Wrap(err, "unable get git server")
		}
		return fmt.Sprintf("https://%v%v", gs.GitHost, *s.GitUrlPath), nil

	default:
		return "", errors.New("unable get project url, caused by the unsupported strategy")
	}
}

// Copy sonar configurations for JavaScript, Python, Go
// It expects workDir - work dir, which contains codebase; td - template dir, which contains sonar.property file
// It returns error in case of issue.
func copySonarConfigs(workDir, td string, config *model.ConfigGoTemplating) (err error) {
	languagesForSonarTemplates := []string{util.LanguageJavascript, util.LanguagePython, util.LanguageGo}
	if !util.CheckElementInArray(languagesForSonarTemplates, strings.ToLower(config.Lang)) {
		return
	}

	sonarConfigPath := fmt.Sprintf("%v/sonar-project.properties", workDir)
	log.Info("start copying sonar configs", "path", sonarConfigPath)

	_, statErr := os.Stat(sonarConfigPath)
	if statErr == nil {
		return nil
	}

	f, err := os.Create(sonarConfigPath)
	if err != nil {
		return err
	}

	defer util.CloseWithErrorCapture(&err, f, "failed to close sonar config file")

	sonarTemplateName := fmt.Sprintf("%v-sonar-project.properties.tmpl", strings.ToLower(config.Lang))
	sonarTemplateFile := fmt.Sprintf("%v/templates/sonar/%v", td, sonarTemplateName)

	tmpl, err := template.New(sonarTemplateName).ParseFiles(sonarTemplateFile)
	if err != nil {
		return err
	}

	err = tmpl.Execute(f, config)
	if err != nil {
		return errors.Wrapf(err, "couldn't render Sonar configs fo %v app: %v", config.Lang, config.Name)
	}

	log.Info("Sonar configs has been copied", "codebase_name", config.Name)

	return
}
