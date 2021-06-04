/*
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import React from 'react';
import App from './App';
import ShallowRenderer from 'react-test-renderer/shallow';

it('renders without crashing', () => {
  const renderer = new ShallowRenderer();
  renderer.render(<App />);
  expect(renderer.getRenderOutput().type).toBe('div');
});
