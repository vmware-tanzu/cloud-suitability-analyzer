/*
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

export default () => {
  return {
    setNavBarLinkActive(index) {
      const links = document.getElementsByTagName('a');
      var i,
        link,
        len = links.length;
      for (i = 0; i < len; ++i) {
        if (i in links) {
          link = links[i];
          if (i === index) {
            link.classList.add('active');
          } else {
            link.classList.remove('active');
          }
        }
      }
      return;
    },
    unsetNavBarLinkActive(index) {
      const link = document.getElementsByTagName('a')[index];
      return link != null ? link.classList.remove('active') : '';
    }
  };
};
