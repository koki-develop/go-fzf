#!/bin/bash

set -euo pipefail

vhs < ./docs/tapes/cli-demo.tape
vhs < ./docs/tapes/library-demo.tape
