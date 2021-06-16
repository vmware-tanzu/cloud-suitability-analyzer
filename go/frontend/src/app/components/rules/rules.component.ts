import {Component, OnInit} from '@angular/core';
import {ClarityIcons, downloadIcon} from '@cds/core/icon';
import '@cds/core/icon/register.js';
import '@cds/core/search/register.js';
import {Router} from "@angular/router";
import {RulesService} from "../../services/rules.service";
import {Rule} from "../../model/rules";
import {ToastrService} from 'ngx-toastr';
import {pushErrorNotification} from '../../utils/notificationutil';
import {ClrDatagridStringFilterInterface} from "@clr/angular";

@Component({
  selector: 'app-rules',
  templateUrl: './rules.component.html',
  styleUrls: ['./rules.component.css']
})
export class RulesComponent implements OnInit {

  rules: Rule[] = [];
  public searchCrit: any = '';
  public tagFilter = new TagFilter();

  constructor(private router: Router, private rulesService: RulesService, public toastr: ToastrService) {
    ClarityIcons.addIcons(downloadIcon);
  }

  ngOnInit(): void {
    this.fetchRules();
  }

  fetchRules(): void {
    this.rulesService.getRules().subscribe(rulesReturned => {
      /* istanbul ignore else */
      if (rulesReturned.rules) {
        this.rules = rulesReturned.rules;
      }
    }, error => {
      pushErrorNotification(error, this.toastr);
    });
  }
}

class TagFilter implements ClrDatagridStringFilterInterface<Rule> {
  accepts(ruleRow: Rule, search: string): boolean {
    // @ts-ignore
    /* istanbul ignore else */
    if (ruleRow.Tags) {
      return ruleRow.Tags.filter(tag => tag.Value.toLowerCase().includes(search)).length > 0;
    }
    return false;
  }
}
