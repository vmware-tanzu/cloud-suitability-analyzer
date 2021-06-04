/*
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import React from 'react';
import ReactDOM from 'react-dom';
import { shallow } from 'enzyme';
import axios from 'axios/index';
import MockAxios from 'axios-mock-adapter';

import SourceCodeComponent from './SourceCodeComponent';

const mockAxios = new MockAxios(axios, { delayResponse: Math.random() * 100 });

it('renders without crashing', () => {
  const div = document.createElement('div');
  ReactDOM.render(<SourceCodeComponent />, div);
  ReactDOM.unmountComponentAtNode(div);
});
