import { Component, OnChanges } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { HttpHeaders , HttpErrorResponse} from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/index';
import { UriConstants } from 'src/app/constants/uri-constants';
import { Scores } from "../model/scores";
import { Language } from "../model/language"
import { Api } from "../model/api"
import { Tags } from "../model/tags"
import { ApplicationFindings } from '../model/applicationfindings';
import { ApplicationScoreCard } from '../model/applicationscorecard';
import { Tag } from '../model/applicationscore';

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

export class ApplicationSummaryService {

  constructor(private http: HttpClient){
  }

  getApplicationByRun(runid: number): Observable<Scores>  {
    const url =
    UriConstants.RUNS_URI +
    runid +
    '/summary/application_scores?rangeMin=1&rangeMax=10&reverse=false';
    return this.http.get<Scores>(url);
  }

  getLanguagesByApplicationRun(runid: number, appName: string): Observable<Language[]>  {
    const url = UriConstants.RUNS_URI + runid + '/apps/' + appName + '/languages';
    return this.http.get<Language[]>(url);
  }

  getApisByApplicationRun(runid: number, appName: string): Observable<Api[]>  {
    const url = UriConstants.RUNS_URI + runid + '/apps/' + appName + '/apis';
    return this.http.get<Api[]>(url);
  }

  getTagsForApplication(runid: number, appName: string): Observable<Tags> {
    const url = UriConstants.RUNS_URI + runid + '/apps/' + appName + '/tags';
    return this.http.get<Tags>(url);
  }

  getApplicationFindings(runid: number, appName: string, tags: Tag[]): Observable<ApplicationFindings>  {
    const url = UriConstants.RUNS_URI + runid + '/apps/' + appName + '/findings/scorecard/low,medium,high?includeFF=false';
    return this.http.post<ApplicationFindings>(url, "");
  }
  
  getApplicationScorecard(runid: number, appName: string, tags: Tag[]): Observable<ApplicationScoreCard>  {
    const url = UriConstants.RUNS_URI + runid + '/apps/' + appName + '/scorecard?includeFF=false';
    return this.http.post<ApplicationScoreCard>(url, "");
  }
}
