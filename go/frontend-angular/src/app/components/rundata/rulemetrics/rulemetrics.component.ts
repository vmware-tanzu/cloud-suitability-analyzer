import { Component, OnInit } from '@angular/core';
import {Router} from "@angular/router";
import {RundataService} from "../../../services/rundata.service";

@Component({
  selector: 'rulemetrics',
  templateUrl: './rulemetrics.component.html',
  styleUrls: ['./rulemetrics.component.css']
})
export class RuleMetricsComponent implements OnInit {

  public searchCrit: any = '';
  public fileName: string = 'rule-metrics.xlsx';

  constructor(private router: Router, private rundataService: RundataService) {

  }

  ngOnInit(): void {
  }

}
