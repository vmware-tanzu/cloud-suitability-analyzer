import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs/index';
import {UriConstants} from 'src/app/constants/uri-constants';
import {AnalyzerRun} from "../model/analyzerrun";

@Injectable({
  providedIn: 'root'
})
export class AnalyzerRunService {

  constructor(private http: HttpClient) { }

  getDistinctRuns(): Observable<AnalyzerRun[]> {
    return this.http.get<AnalyzerRun[]>(UriConstants.ANALYZER_RUNS_URI);
  }

  getForgeVersion(): Observable<string> {
    return this.http.get<string>(UriConstants.VERSION_URI);
  }
}
