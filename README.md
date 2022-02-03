# Run Expo Application Services (EAS) build

[![Step changelog](https://shields.io/github/v/release/bitrise-steplib/bitrise-step-run-eas-build?include_prereleases&label=changelog&color=blueviolet)](https://github.com/bitrise-steplib/bitrise-step-run-eas-build/releases)

Runs a build on Expo Application Services (EAS).

<details>
<summary>Description</summary>

Runs a build on Expo Application Services (EAS).

The step runs `EXPO_TOKEN=[access_token] npx eas-cli build --platform [platform] --non-interactive [eas_options]`

in the provided `[work_dir]`.
</details>

## 🧩 Get started

Add this step directly to your workflow in the [Bitrise Workflow Editor](https://devcenter.bitrise.io/steps-and-workflows/steps-and-workflows-index/).

You can also run this step directly with [Bitrise CLI](https://github.com/bitrise-io/bitrise).

## ⚙️ Configuration

<details>
<summary>Inputs</summary>

| Key | Description | Flags | Default |
| --- | --- | --- | --- |
| `access_token` | Expo Access Token.  Visit [Expo Documentation](https://docs.expo.dev/accounts/programmatic-access) to generate one. | required, sensitive |  |
| `platform` | Platform to build. | required | `all` |
| `work_dir` | Directory containing the Expo project (`app.json`). | required | `$BITRISE_SOURCE_DIR` |
| `eas_options` | Additional options for the eas command.  The step runs `EXPO_TOKEN=[access_token] npx eas-cli build --platform [platform] --non-interactive`, use this input to pass additional option to the command.  Example: `--profile=development`. |  |  |
</details>

<details>
<summary>Outputs</summary>
There are no outputs defined in this step
</details>

## 🙋 Contributing

We welcome [pull requests](https://github.com/bitrise-steplib/bitrise-step-run-eas-build/pulls) and [issues](https://github.com/bitrise-steplib/bitrise-step-run-eas-build/issues) against this repository.

For pull requests, work on your changes in a forked repository and use the Bitrise CLI to [run step tests locally](https://devcenter.bitrise.io/bitrise-cli/run-your-first-build/).

Learn more about developing steps:

- [Create your own step](https://devcenter.bitrise.io/contributors/create-your-own-step/)
- [Testing your Step](https://devcenter.bitrise.io/contributors/testing-and-versioning-your-steps/)
