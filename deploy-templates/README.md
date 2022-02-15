# codebase-operator

![Version: 2.11.0-SNAPSHOT](https://img.shields.io/badge/Version-2.11.0--SNAPSHOT-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 2.11.0-SNAPSHOT](https://img.shields.io/badge/AppVersion-2.11.0--SNAPSHOT-informational?style=flat-square)

A Helm chart for EDP Codebase Operator

**Homepage:** <https://epam.github.io/edp-install/>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| epmd-edp | SupportEPMD-EDP@epam.com | https://solutionshub.epam.com/solution/epam-delivery-platform |
| sergk |  | https://github.com/SergK |

## Source Code

* <https://github.com/epam/edp-codebase-operator>

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` |  |
| annotations | object | `{}` |  |
| envs[0].name | string | `"RECONCILATION_PERIOD"` |  |
| envs[0].value | string | `"360"` |  |
| envs[1].name | string | `"CODEBASE_BRANCH_MAX_CONCURRENT_RECONCILES"` |  |
| envs[1].value | string | `"3"` |  |
| global.database.enabled | bool | `true` |  |
| global.database.host | string | `nil` |  |
| global.database.name | string | `"edp-db"` |  |
| global.database.port | int | `5432` |  |
| global.edpName | string | `""` |  |
| global.platform | string | `"openshift"` |  |
| image.name | string | `"epamedp/codebase-operator"` |  |
| image.version | string | `nil` |  |
| imagePullPolicy | string | `"IfNotPresent"` |  |
| jira.apiUrl | string | `"https://jiraeu-api.example.com"` |  |
| jira.credentialName | string | `"jira-user"` |  |
| jira.integration | bool | `false` |  |
| jira.name | string | `"jira"` |  |
| jira.rootUrl | string | `"https://jiraeu.example.com"` |  |
| name | string | `"codebase-operator"` |  |
| nodeSelector | object | `{}` |  |
| resources.limits.memory | string | `"192Mi"` |  |
| resources.requests.cpu | string | `"50m"` |  |
| resources.requests.memory | string | `"64Mi"` |  |
| tolerations | list | `[]` |  |
