/*
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import { configure } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';

// Globally Configure the Enzyme adapter for testing react code
configure({ adapter: new Adapter() });
