---
name: Bug Report
about: Help us diagnose and fix bugs in Official AWS Provider
labels: bug,needs:triage
title: "[Bug]: "
body:
  - type: textarea
    id: problem
    attributes:
      label: What happened?
      description: |
        Please provide as much info as possible. Not doing so may result in your bug not being addressed in a timely manner.
        If this matter is security related, please disclose it privately via https://kubernetes.io/security
    validations:
      required: true

  - type: textarea
    id: expected
    attributes:
      label: What did you expect to happen?
    validations:
      required: true

  - type: textarea
    id: repro
    attributes:
      label: How can we reproduce it (as minimally and precisely as possible)?
    validations:
      required: true

  - type: textarea
    id: additional
    attributes:
      label: Anything else we need to know?

  - type: textarea
    id: kubeVersion
    attributes:
      label: Kubernetes version
      value: |
        <details>

        ```console
        $ kubectl version
        # paste output here
        ```

        </details>
    validations:
      required: true

  - type: textarea
    id: cloudProvider
    attributes:
      label: Cloud provider
      value: |
        <details>

        </details>
    validations:
      required: true

  - type: textarea
    id: osVersion
    attributes:
      label: OS version
      value: |
        <details>

        ```console
        # On Linux:
        $ cat /etc/os-release
        # paste output here
        $ uname -a
        # paste output here

        # On Windows:
        C:\> wmic os get Caption, Version, BuildNumber, OSArchitecture
        # paste output here
        ```

        </details>

  - type: textarea
    id: installer
    attributes:
      label: Install tools
      value: |
        <details>

        </details>

  - type: textarea
    id: runtime
    attributes:
      label: Container runtime (CRI) and version (if applicable)
      value: |
        <details>

        </details>

  - type: textarea
    id: plugins
    attributes:
      label: Related plugins (CNI, CSI, ...) and versions (if applicable)
      value: |
        <details>

        </details>

---
<!--
Thank you for helping to improve Official AWS Provider!

Please be sure to search for open issues before raising a new one. We use issues
for bug reports and feature requests.
-->

### What happened?
<!--
Please let us know what behaviour you expected and how Official AWS Provider diverged from
that behaviour.
-->

### How can we reproduce it?
<!--
Help us to reproduce your bug as succinctly and precisely as possible. Artifacts
such as example manifests or a script that triggers the issue are highly
appreciated!
-->

### What environment did it happen in?

* Crossplane Version:
* Provider Version:
* Kubernetes Version: <!-- use `kubectl version` --> 
* Kubernetes Distribution: <!-- EKS, AKS, GKE, OpenShift, etc. -->
