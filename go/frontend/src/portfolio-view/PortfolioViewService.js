/*
 * Copyright (c) 2018 - Present VMWare, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import axios from 'axios';

const BASE_URL = '/api/runs';
export default () => {
  return {
    /* api/runs/1/summary/top_languages_by_codelines */
    getLanguagesByLoc(runid) {
      return axios
        .get(BASE_URL + '/' + runid + '/summary/top_languages_by_codelines')
        .then(response => response.data);
    },

    /* /api/runs/1/summary/top_apis_by_count */
    getApisByScore(runid) {
      return axios
        .get(BASE_URL + '/' + runid + '/summary/top_apis_by_score')
        .then(response => response.data);
    },

    /* /api/runs/1/summary/apps_for_language?lang=java */
    getAppsByLanguage(runid, language) {
      return axios
        .get(
          BASE_URL + '/' + runid + '/summary/apps_for_language?lang=' + language
        )
        .then(response => response.data);
    },

    /* /api/runs/1/summary/top_apps_for_api?api=jsfapi */
    getAppsByApi(runid, api) {
      return axios
        .get(BASE_URL + '/' + runid + '/summary/top_apps_for_api?api=' + api)
        .then(response => response.data);
    }
  };
};
