import { Component, OnInit } from '@angular/core';
import {Router} from "@angular/router";
import {RundataService} from "../../../services/rundata.service";

@Component({
  selector: 'rulebyapp',
  templateUrl: './rulebyapp.component.html',
  styleUrls: ['./rulebyapp.component.css']
})
export class RuleByAppComponent implements OnInit {

  public searchCrit: any = '';
  public fileName: string = 'rule-by-app.xlsx';

  constructor(private router: Router, private rundataService: RundataService) {

  }

  ngOnInit(): void {
  }

}
