# konflux-github-sample

Minimal Docker-based HTTP service for learning [Konflux](https://konflux.pages.redhat.com/) component onboarding from GitHub.

The app exposes:

- `GET /` — plain-text hello message
- `GET /health` — JSON health check

Konflux expects a repository with source code and a **`Dockerfile` at the repo root** (default path). This repo intentionally does **not** include `.tekton/` pipelines yet — Konflux creates those when you onboard the component.

## Quick local test

```bash
# Run directly
go run .

# Or build and run the container
make docker-run
curl http://localhost:8080/
curl http://localhost:8080/health
```

## Push to GitHub

Create a new GitHub repository and push this code:

```bash
cd konflux-github-sample
git init
git add .
git commit -S -m "Initial commit: Konflux learning sample"
git branch -M main
git remote add origin git@github.com:<your-user>/konflux-github-sample.git
git push -u origin main
```

Replace `<your-user>` with your GitHub username or organization.

## Onboard in Konflux (GitHub UI flow)

These steps follow the Red Hat Konflux guide for [creating a component from GitHub](https://konflux.pages.redhat.com/docs/users/building/creating-github.html#creating-a-component-with-the-ui).

### 1. Install the Konflux GitHub App

1. Open the Konflux GitHub App install page (from Konflux docs or your tenant onboarding email).
2. Click **Install App** and choose your GitHub user or organization.
3. Select **Only select repositories** and pick `konflux-github-sample`.

Konflux needs read access to the repo and permission to open PRs that add Tekton pipeline definitions.

### 2. Create an Application

1. Log in to the Konflux UI for your tenant (for example `experimental-bmandal-tenant`).
2. Go to **Applications** → **Create application**.
3. Enter a name such as `konflux-learning`.
4. Click **Create application**.

### 3. Add a Component from GitHub

Inside the new application, click **Add component** and fill in:

| Field | Value |
| --- | --- |
| Git repository URL | `https://github.com/<your-user>/konflux-github-sample.git` |
| Revision | `main` |
| Context directory | `./` (default) |
| Dockerfile | `Dockerfile` (default — leave blank in UI if it auto-fills) |
| Pipeline | `docker-build-oci-ta` |
| Private image | **Unchecked** for first tests (PR pipelines work more smoothly with public images) |

Then click **Create application** / **Create component**.

### 4. Merge the onboarding PR

Konflux opens a GitHub pull request in your repo that adds:

- `.tekton/` pipeline definitions
- `pipelinesascode.tekton.dev` configuration
- Other build metadata

Review and merge that PR. After merge, Konflux triggers a build pipeline run.

### 5. Watch the build

In the Konflux UI:

1. Open your application → component.
2. Check **Pipeline runs** for clone → build → push stages.
3. When successful, the built image appears in your tenant registry.

You can also watch checks on GitHub after the onboarding PR merges.

## What Konflux is doing (learning checklist)

Use this repo to observe each Konflux concept:

- [ ] **Application** groups one or more components
- [ ] **Component** maps to a Git repo + Dockerfile path
- [ ] **Pipelines as Code** adds `.tekton/` via PR
- [ ] **Pipeline run** clones GitHub, builds with your Dockerfile, pushes an image
- [ ] **Enterprise Contract / tasks** run policy checks on the build
- [ ] **Image repository** is created for the component output

## Repository layout

```text
.
├── Dockerfile          # Default build file Konflux looks for
├── main.go             # Small HTTP server
├── go.mod
├── Makefile            # Local dev helpers
└── README.md           # This file
```

## CLI alternative (optional)

If you prefer GitOps over the UI, you can create resources with `kubectl` in your tenant namespace. See the Konflux docs for the `Component` manifest fields:

- `spec.source.git.url` — HTTPS clone URL ending in `.git`
- `spec.source.git.dockerfileUrl` — `Dockerfile`
- Annotations such as `build.appstudio.openshift.io/request: configure-pac`

Example doc: https://konflux-ci.dev/docs/building/creating-github/

## Troubleshooting

| Symptom | Things to check |
| --- | --- |
| Onboarding PR never appears | GitHub App installed on the correct repo; repo URL uses `https://` form |
| Pipeline fails cloning | Repository is accessible to the Konflux GitHub integration |
| Build fails on base image | VPN/network access to `registry.access.redhat.com` from the build cluster |
| PR pipelines do not run | Try with **public** output image on first onboarding |

## References

- [Creating a component from GitHub (Red Hat Konflux docs)](https://konflux.pages.redhat.com/docs/users/building/creating-github.html)
- [Official Konflux sample (Golang)](https://github.com/konflux-ci/sample-component-golang)
- [Creating applications and components](https://konflux-ci.dev/docs/building/creating/)
