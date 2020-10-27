/*
 * Copyright (c) 2018 - Present VMWare, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import React from 'react';
import ReactDOM from 'react-dom';
import { DashboardCard } from './DashboardCard';

it('renders without crashing', () => {
  const div = document.createElement('div');
  ReactDOM.render(<DashboardCard />, div);
  ReactDOM.unmountComponentAtNode(div);
});
