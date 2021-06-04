/*
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import axios from 'axios';

export default () => {
  return {
    getDistinctRuns() {
      const url = '/api/analyze-runs';
      return axios.get(url).then(response => response.data.runs);
    },
    getForgeVersion() {
      const url = '/api/version';
      return axios.get(url).then(response => response.data.version);
    }
  };
};
