import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs/index';
import {UriConstants} from 'src/app/constants/uri-constants';
import {AnalyzerRuns} from '../model/analyzerrun';

@Injectable({
  providedIn: 'root'
})
export class AnalyzerRunService {

  constructor(private http: HttpClient) { }

  getDistinctRuns(): Observable<AnalyzerRuns> {
    return this.http.get<AnalyzerRuns>(UriConstants.ANALYZER_RUNS_URI);
  }

  getForgeVersion(): Observable<string> {
    return this.http.get<string>(UriConstants.VERSION_URI);
  }
}
