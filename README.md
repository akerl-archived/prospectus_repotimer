prospectus_repotimer
=========

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/akerl/prospectus_repotimer/Build)](https://github.com/akerl/prospectus_repotimer/actions)
[![GitHub release](https://img.shields.io/github/release/akerl/prospectus_repotimer.svg)](https://github.com/akerl/prospectus_repotimer/releases)
[![MIT Licensed](https://img.shields.io/badge/license-MIT-green.svg)](https://tldrlegal.com/license/mit-license)


[Prospectus](https://github.com/akerl/prospectus) plugin for checking for repo staleness

## Usage

Add a check to your `.prospectus.d` directory with the following contents:

```
#!/usr/bin/env prospectus_repotimer
days: 7
```

Replace the `days` value as desired to set how frequently the repo should be updated.

## Installation

```
go get https://github.com/akerl/prospectus_repotimer
```

## License

prospectus_repotimer is released under the MIT License. See the bundled LICENSE file for details.

