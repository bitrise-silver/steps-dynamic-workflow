The Step generates dynamic yml file based on template and commit to a repo which later triggers a Bitrise run.

<details>
<summary>Description</summary>

The template uses syntax defined in [Go template](https://pkg.go.dev/text/template)

TBA

### Configuring the Step

1. The **Git repository URL** and the ** Git branch** fields are essential and require a SSH key installed with write permission so that the step can commit generated Bitrise.yml file to the repo. The branch should be ephemeral and unique per run.

TBA.

### Related Steps
 
- [Activate SSH key (RSA private key)](https://www.bitrise.io/integrations/steps/activate-ssh-key)

</details>

## 🧩 Get started

Add this step directly to your workflow in the [Bitrise Workflow Editor](https://devcenter.bitrise.io/steps-and-workflows/steps-and-workflows-index/).

You can also run this step directly with [Bitrise CLI](https://github.com/bitrise-io/bitrise).

## ⚙️ Configuration

<details>
<summary>Inputs</summary>

| Key | Description | Flags | Default |
| --- | --- | --- | --- |
| `template_file` | Path to template file | required |  |
| `value_file` | File where values defined before being passed to template | required |  |
| `repository_url` | Repository which generated Bitrise yml is commited into | required | $GIT_REPOSITORY_URL |
| `repository_branch` | Branch which generated Bitrise yml is commited into. Ideally the branch should be ephemeral and unique per run. | required | bitrise-dynamic-workflow-$BITRISE_BUILD_NUMBER |
| `clean_up_branch_after_used` | Wait for the build to complete then remove the `repository_branch` |  | `yes` |
| `access_token` | Bitrise Access Token to trigger new builds | required | $BITRISE_TOKEN |
| `pipeline_id` | Pipeline ID to be trigger | required | |
</details>

<details>
<summary>Outputs</summary>

| Environment Variable | Description |
| --- | --- |
| `DYNAMIC_TRIGGERED_BUILD_ID` | Triggered Bitrise build id |
</details>

## 🙋 Contributing

We welcome [pull requests](https://github.com/bitrise-steplib/steps-git-clone/pulls) and [issues](https://github.com/bitrise-steplib/steps-git-clone/issues) against this repository.

For pull requests, work on your changes in a forked repository and use the Bitrise CLI to [run step tests locally](https://devcenter.bitrise.io/bitrise-cli/run-your-first-build/).

Learn more about developing steps:

- [Create your own step](https://devcenter.bitrise.io/contributors/create-your-own-step/)
- [Testing your Step](https://devcenter.bitrise.io/contributors/testing-and-versioning-your-steps/)