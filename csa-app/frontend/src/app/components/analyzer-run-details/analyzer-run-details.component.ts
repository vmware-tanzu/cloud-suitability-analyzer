import {Component, Input, OnInit} from '@angular/core';
import {Router} from '@angular/router';

@Component({
  selector: 'app-analyzer-run-details',
  templateUrl: './analyzer-run-details.component.html',
  styleUrls: ['./analyzer-run-details.component.css']
})
export class AnalyzerRunDetailsComponent implements OnInit {

  @Input()
  selectedRunId: number | undefined;

  @Input()
  analyzerRun: any;

  constructor(private router: Router) {
  }

  ngOnInit(): void {
  }

}
