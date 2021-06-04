/*
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

export default () => {
  return {
    generateRunItemsDropDown(resp) {
      let runItems = [];
      if (resp != null) {
        resp.forEach(run => {
          runItems.push({
            label: run.id + ' - ' + run.Alias,
            value: run
          });
        });
        runItems.sort();
        runItems.reverse();
      }
      return runItems;
    },
    setSelectedRun(runItems) {
      return runItems != null && runItems.length > 0 ? runItems[0].value : '';
    },
    generateApplicationItemsDropDown(resp) {
      //console.log('Building Dropdown with App List => ' + JSON.stringify(resp));
      let applicationItems = [];
      if (resp != null) {
        resp.forEach(app => {
          applicationItems.push({
            label: app.name,
            value: app.name
          });
        });
      }
      return applicationItems;
    },
    setSelectedApplication(applicationItems, app) {
      //console.log('Setting Selected App for DropDown in App View. App In [%s] AppItems [%s]', app, JSON.stringify(applicationItems))
      let result = '';

      if (applicationItems !== null && applicationItems.length > 0) {
        result = applicationItems[0].value;
        if (app !== '') {
          for (var i = 0; i < applicationItems.length; i++) {
            //console.log('found app [%s] for comparison', applicationItems[i].value)
            if (applicationItems[i].value.localeCompare(app) === 0) {
              result = app;
              break;
            }
          }
        }
      }
      console.log('Set app drop down => %s', result);
      return result;
    },
    generateLanguageItemsDropDown(resp) {
      let languageItems = [];
      if (resp != null) {
        resp.forEach(lang => {
          languageItems.push({ label: lang.language, value: lang.language });
        });
      }
      return languageItems;
    },
    setSelectedLanguage(languageItems) {
      return languageItems != null && languageItems.length > 0
        ? languageItems[0].value
        : '';
    },
    generateApiItemsDropDown(resp) {
      let apiItems = [];
      if (resp != null) {
        resp.forEach(api => {
          apiItems.push({ label: api.api, value: api.api });
        });
      }
      return apiItems;
    },
    setSelectedApi(apiItems) {
      return apiItems != null && apiItems.length > 0 ? apiItems[0].value : '';
    }
  };
};
