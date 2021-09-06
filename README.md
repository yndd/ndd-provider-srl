<p align=center><a href="https://nddocs.yndd.io"><img src=https://gitlab.com/rdodin/pics/-/wikis/uploads/9f2e581a8d207a21ff024a312679a239/containerlab_export_white_ink_3?sanitize=true/></a></p>

[![github release](https://img.shields.io/github/release/yndd/ndd-provider-srl.svg?style=flat-square&color=00c9ff&labelColor=bec8d2))](https://github.com/yndd/ndd-provider-srl/releases/)
[![Go Report](https://img.shields.io/badge/go%20report-A%2B-blue?style=flat-square&color=00c9ff&labelColor=bec8d2)](https://goreportcard.com/report/github.com/yndd/ndd-provider-srl)
[![Godoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square&color=00c9ff&labelColor=bec8d2)](https://godoc.org/github.com/yndd/ndd-provider-srl)
[![Doc](https://img.shields.io/badge/Docs-ndddocs.yndd.io-blue?style=flat-square&color=00c9ff&labelColor=bec8d2)](https://ndddocs.yndd.io)
[![Docker Pulls](https://img.shields.io/docker/pulls/yndd/ndd-provider-srl-controller.svg?style=flat-square&color=00c9ff&labelColor=bec8d2))](https://img.shields.io/docker/pulls/yndd/ndd-provider-srl-controller.svg) 
[![Twitter Follow](https://img.shields.io/twitter/follow/yndd.svg?style=social&label=Follow)](https://twitter.com/intent/follow?screen_name=yndd&user_id=1434394355385651201)


## Overview
 
ndd-provider-srl implements an [srlinux] provider, which exposes its configuration in kubernetes through [CRs]. 

Features:

* Device discovery and Provider registration
* Declaritive CRUD configuration of network devices through [CRs]
* Configuration Input Validation:
    - Declarative validation using an OpenAPI v3 schema derived from [YANG]
    - Runtime Dependency Management amongst the various resources comsumed within a device (parent dependency management and leaf reference dependency management amont resources)
* Automatic or Operator interacted configuration drift management
* Delete Policy, and Active etc  

## Releases

ndd-provider-srl is in alpha phase so dont use it in production

## Getting Started

Take a look at the [documentation] to get started.

## Get involved

ndd is a community driven project and we welcome contribution.

- Discord: [discord]
- Twitter: [@yndd]
- Email: [info@yndd.io]

For filling bugs, suggesting improvments, or requesting new feature, please open an [issue].

## Code of conduct

## Licensing

ndd-provider-srl is under the Apache 2.0 license.

[documentation]: https://ndddocs.yndd.io
[issue]: https://github.com/yndd/ndd-core/issues
[roadmap]: https//github.com/yndd/tbd
[discord]: https://discord.gg/prHcBMSq
[@yndd]: https://twitter.com/yndd
[info@yndd.io]: mailto:info@yndd.io

[Kubernetes]: https://kubernetes.io
[YANG]: https://en.wikipedia.org/wiki/YANG
[CRs]: https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/
[kubebuilder]: https://kubebuilder.io
[operator-pattern]: https://kubernetes.io/docs/concepts/extend-kubernetes/operator/
[srlinux]: https://www.nokia.com/networks/products/service-router-linux-NOS/