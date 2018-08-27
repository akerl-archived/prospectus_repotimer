prospectus_repotimer
=========

[![Gem Version](https://img.shields.io/gem/v/prospectus_repotimer.svg)](https://rubygems.org/gems/prospectus_repotimer)
[![Build Status](https://img.shields.io/circleci/project/akerl/prospectus_repotimer.svg)](https://circleci.com/gh/akerl/prospectus_repotimer)
[![Coverage Status](https://img.shields.io/codecov/c/github/akerl/prospectus_repotimer.svg)](https://codecov.io/github/akerl/prospectus_repotimer)
[![Code Quality](https://img.shields.io/codacy/c5623564a4034ece993510d28edb19de.svg)](https://www.codacy.com/app/akerl/prospectus_repotimer)
[![MIT Licensed](https://img.shields.io/badge/license-MIT-green.svg)](https://tldrlegal.com/license/mit-license)

[Prospectus](https://github.com/akerl/prospectus) helpers for checking for repo staleness

## Usage

Add the following 2 lines to the .prospectus:

```
## Add this at the top
Prospectus.extra_dep('file', 'prospectus_repotimer')

## Add this inside your item that needs a timer
extend ProspectusRepotimer::Timer.new(days_ago)
```

## Installation

    gem install prospectus_repotimer

## License

prospectus_repotimer is released under the MIT License. See the bundled LICENSE file for details.

