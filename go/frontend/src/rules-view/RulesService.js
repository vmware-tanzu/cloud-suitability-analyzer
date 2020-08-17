/*
 * Copyright (c) 2018 - Present VMWare, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import axios from 'axios';
export class RulesService {
  getRules() {
    return axios.get('/api/rules').then(response => response.data.rules);
  }
}
