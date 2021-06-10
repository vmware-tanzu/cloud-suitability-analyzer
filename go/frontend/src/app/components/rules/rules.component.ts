import { Component, OnInit } from '@angular/core';
import {ClarityIcons, downloadIcon} from '@cds/core/icon';
import '@cds/core/icon/register.js';
import '@cds/core/search/register.js';
import {Router} from "@angular/router";
import {RulesService} from "../../services/rules.service";
import {Rule} from "../../model/rules";
import {ToastrService} from 'ngx-toastr';
import { pushErrorNotification } from '../../utils/notificationutil';

@Component({
  selector: 'app-rules',
  templateUrl: './rules.component.html',
  styleUrls: ['./rules.component.css']
})
export class RulesComponent implements OnInit {

  rules: Rule[] = [];
  public searchCrit: any = '';

  constructor(private router: Router, private rulesService: RulesService, public toastr: ToastrService) {
    ClarityIcons.addIcons(downloadIcon);
  }

  ngOnInit(): void {
    this.fetchRules();
  }

  fetchRules(): void {
    this.rulesService.getRules().subscribe(rulesReturned => {
      if (rulesReturned.rules) {
        this.rules = rulesReturned.rules;
      }
    }, error => {
      pushErrorNotification(error, this.toastr);
    });
  }
}
