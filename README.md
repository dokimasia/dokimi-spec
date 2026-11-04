<!--
  ~ Copyright Dokimasia B.V. 2026
  ~ SPDX-License-Identifier: Apache-2.0
-->

# dokimi-spec

dokimi-spec contains the conformance spec of dokimi: the laws, claims, gates, probes, reason codes,
defect forms, machineries and stances that dokimi and its runtimes share.

You edit the YAML files of [`spec`](spec), and `go generate` renders each of them into its JSON
file, its JSON Schema and the constants of `names.gen.go`. The Go module `go.dokimi.dev/spec` in
`spec` embeds the JSON files, and dokimi and the Go runtime require that module. A runtime in
another language vendors the JSON files with [spec](https://github.com/stealth-scale/spec).

```sh
go get go.dokimi.dev/spec
```

## Development

`make check` runs the gate of the repository: the formatters, the linters and the tests of every
language. [CONTRIBUTING.md](CONTRIBUTING.md) describes how to propose a change.

## License

dokimi-spec is licensed under Apache-2.0. See [LICENSE](LICENSE).
