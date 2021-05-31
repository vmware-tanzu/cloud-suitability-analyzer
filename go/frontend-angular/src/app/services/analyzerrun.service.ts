import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs/index';
import { catchError } from 'rxjs/operators';
import { throwError } from 'rxjs';
import {HttpHeaders, HttpErrorResponse} from '@angular/common/http';
import {UriConstants} from 'src/app/constants/uri-constants';
import {AnalyzerRun} from "../model/analyzerrun";

const httpOptions = {
  headers: new HttpHeaders({
    'Content-Type': 'application/json',
    'Access-Control-Request-Headers': '*',
    'Access-Control-Allow-Origin': '*'
  })
};

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
