/*
 * Copyright (c) 2018 - Present VMWare, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import React from 'react';
import ReactDOM from 'react-dom';
import axios from 'axios/index';
import MockAxios from 'axios-mock-adapter';

import ApiUsageSummaryComponent from './ApiUsageSummaryComponent';
import { shallow } from 'enzyme';

const mockAxios = new MockAxios(axios, { delayResponse: Math.random() * 100 });

it('renders without crashing', () => {
  const div = document.createElement('div');
  ReactDOM.render(<ApiUsageSummaryComponent />, div);
  ReactDOM.unmountComponentAtNode(div);
});
