#!/usr/bin/env python
# Copyright 2014 Marc-Antoine Ruel. All rights reserved.
# Use of this source code is governed by the Apache v2.0 license that can be
# found in the LICENSE file.

import sys

import find_gae_sdk

# TODO(maruel): Figure out a way so pass --skip_sdk_update_check
sys.exit(find_gae_sdk.run(['goapp', 'deploy', '-oauth'] + sys.argv[1:]))
