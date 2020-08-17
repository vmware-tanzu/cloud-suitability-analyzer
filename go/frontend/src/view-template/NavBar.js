/*
 * Copyright (c) 2018 - Present VMWare, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import React, { Component } from 'react';
import { Menubar } from 'primereact/components/menubar/Menubar';

const items = [
  {
    label: 'Summary',
    command: () => {
      window.location.hash = '/';
    }
  },
  {
    label: 'Portfolio',
    command: () => {
      window.location.hash = '/portfolio';
    }
  },
  {
    label: 'Application',
    command: () => {
      window.location.hash = '/applications';
    }
  },
  {
    label: 'Data',
    command: () => {
      window.location.hash = '/data';
    }
  },
  {
    label: 'Rule',
    command: () => {
      window.location.hash = '/rules';
    }
  }
];

export const NavBar = () => (
  <div>
    <Menubar model={items} style={{ border: 'none', background: '#edf0f5' }}>
      <img
        id="logo"
        src="/VMWare-logo.svg"
        width="60"
        height="60"
        alt="vmware-logo"
      />
    </Menubar>
  </div>
);
