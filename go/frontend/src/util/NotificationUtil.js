/*
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

export function pushErrorNotification(err, toast) {
  if (toast != null && typeof toast.show !== 'undefined') {
    toast.show({
      severity: 'error',
      summary: 'Error Message',
      detail: err.toString()
    });
  } else {
    console.error(
      'An Error occurred and could not be displayed via toast! Details => ' +
        JSON.stringify(err)
    );
  }
}

export function pushNotification(msg, toast) {
  if (toast != null && typeof toast.show !== 'undefined') {
    toast.show({
      severity: 'info',
      summary: 'Information',
      detail: msg
    });
  } else {
    console.log(
      'An informational event/msg was received and could not be displayed via toast! Details => ' +
        JSON.stringify(msg)
    );
  }
}

export function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms));
}
