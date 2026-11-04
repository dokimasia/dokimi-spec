<!--
  ~ Copyright Dokimasia B.V. 2026
  ~ SPDX-License-Identifier: Apache-2.0
-->

# dokimi-spec

dokimi-spec contains the runtime contract of the conformance spec of dokimi. The contract consists
of the laws, claims, gates, probes, reason codes, defect forms, machineries and stances, the key
form of a check identity and the schema of the suite record.

Each language repository vendors the contract with [spec](https://github.com/stealth-scale/spec)
and checks its runtime against it. [dokimi-go](https://github.com/dokimasia/dokimi-go) vendors the
contract into its module `go.dokimi.dev/spec`. The generator
[dokimi](https://github.com/dokimasia/dokimi) reads the contract through that module, and it keeps
the rules that derive checks from a declaration in its own repository.

## Change the contract

1. Edit the YAML files of [`spec`](spec).
2. Set `spec/VERSION` to the version of the change.
3. Run `go generate ./...` in `spec`. The renderer writes the JSON file and the JSON Schema of each
   YAML file, `naming.json`, `keys.json` and `names.gen.go`.
4. Run `make generate-spec` to write `spec/manifest.json`.

`ergon tool run spec.spec -- try` writes the working tree into the checkout of each implementation
beside this repository, such as `../dokimi-go`, and runs its tests there.

The Go module in `spec` renders the files and checks that they agree with each other. No
repository imports the module, and it has no release.

## Development

`make check` runs the gate of the repository: the formatters, the linters and the tests of every
language. [CONTRIBUTING.md](CONTRIBUTING.md) describes how to propose a change.

## License

dokimi-spec is licensed under Apache-2.0. See [LICENSE](LICENSE).
