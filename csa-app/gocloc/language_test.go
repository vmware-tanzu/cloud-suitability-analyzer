/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 ******************************************************************************/

package gocloc

import (
	"testing"

	"csa-app/util"
)

func TestGetShebang(t *testing.T) {
	lang := "py"
	shebang := "#!/usr/bin/env python"

	s, ok := util.GetShebang(shebang)
	if !ok {
		t.Errorf("invalid logic. shebang=[%v]", shebang)
	}

	if lang != s {
		t.Errorf("invalid logic. lang=[%v] shebang=[%v]", lang, s)
	}
}

func TestGetShebangWithSpace(t *testing.T) {
	lang := "py"
	shebang := "#! /usr/bin/env python"

	s, ok := util.GetShebang(shebang)
	if !ok {
		t.Errorf("invalid logic. shebang=[%v]", shebang)
	}

	if lang != s {
		t.Errorf("invalid logic. lang=[%v] shebang=[%v]", lang, s)
	}
}

func TestGetShebangBashWithEnv(t *testing.T) {
	lang := "bash"
	shebang := "#!/usr/bin/env bash"

	s, ok := util.GetShebang(shebang)
	if !ok {
		t.Errorf("invalid logic. shebang=[%v]", shebang)
	}

	if lang != s {
		t.Errorf("invalid logic. lang=[%v] shebang=[%v]", lang, s)
	}
}

func TestGetShebangBash(t *testing.T) {
	lang := "bash"
	shebang := "#!/usr/bin/bash"

	s, ok := util.GetShebang(shebang)
	if !ok {
		t.Errorf("invalid logic. shebang=[%v]", shebang)
	}

	if lang != s {
		t.Errorf("invalid logic. lang=[%v] shebang=[%v]", lang, s)
	}
}

func TestGetShebangBashWithSpace(t *testing.T) {
	lang := "bash"
	shebang := "#! /usr/bin/bash"

	s, ok := util.GetShebang(shebang)
	if !ok {
		t.Errorf("invalid logic. shebang=[%v]", shebang)
	}

	if lang != s {
		t.Errorf("invalid logic. lang=[%v] shebang=[%v]", lang, s)
	}
}

func TestGetShebangPlan9Shell(t *testing.T) {
	lang := "plan9sh"
	shebang := "#!/usr/rc"

	s, ok := util.GetShebang(shebang)
	if !ok {
		t.Errorf("invalid logic. shebang=[%v]", shebang)
	}

	if lang != s {
		t.Errorf("invalid logic. lang=[%v] shebang=[%v]", lang, s)
	}
}

func TestGetShebangStartDot(t *testing.T) {
	lang := "pl"
	shebang := "#!./perl -o"

	s, ok := util.GetShebang(shebang)
	if !ok {
		t.Errorf("invalid logic. shebang=[%v]", shebang)
	}

	if lang != s {
		t.Errorf("invalid logic. lang=[%v] shebang=[%v]", lang, s)
	}
}
