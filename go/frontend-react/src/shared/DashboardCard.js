/*
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import React from 'react';

export const DashboardCard = ({
  title,
  detail,
  summaryType,
  backgroundclass
}) => (
  <div className="card summary">
    <span className="csa-title">{title}</span>
    <span className="detail">{detail}</span>
    <span className={backgroundclass}>{summaryType}</span>
  </div>
);
