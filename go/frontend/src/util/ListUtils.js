/*
 * Copyright (c) 2018 - Present VMWare, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import React from 'react';

export default () => {
  return {
    renderList(items) {
      return (
        <div>
          {items.map(item => {
            return <li key={item}>{item}</li>;
          })}
        </div>
      );
    },
    renderLinks(items) {
      return (
        <div>
          {items.map(item => {
            return (
              <li key={item}>
                <a className="csa-link" href={item}>
                  {item}
                </a>
              </li>
            );
          })}
        </div>
      );
    }
  };
};
