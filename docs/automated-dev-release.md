# Automated dev release

On every push to `main`, CI tests and publishes an immutable ARM64 image to
Amazon ECR with the tag `sha-<full-commit-sha>`.

When the GitHub App secrets are configured, the same workflow creates a branch
in `courier-common`, changes only the user-service image tag, opens a release
pull request, and immediately squash-merges it for the development environment.
Argo CD reconciles that merge and performs the rollout.

The release pull request is an audit record, not an approval gate. Its body
always records the service, source pull request, update summary, image digest,
and source CI run.

## Required GitHub Actions secrets

Configure these in `courier-user-svc` after installing the release GitHub App
on `courier-common` with `Contents: Read and write` and `Pull requests: Read
and write`:

- `COURIER_RELEASE_APP_ID`
- `COURIER_RELEASE_APP_PRIVATE_KEY`

Until both secrets exist, CI still tests and publishes ECR images but explicitly
skips GitOps promotion.
