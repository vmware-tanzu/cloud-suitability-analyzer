import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/index';
import { UriConstants } from 'src/app/constants/uri-constants';
import { Scores } from "../model/scores";
import { Language } from "../model/language"
import { Api } from "../model/api"
import { ApplicationFinding } from '../model/applicationfinding';
import { ApplicationScoreCard } from '../model/applicationscorecard';
import { Bin, Tag } from '../model/applicationscore';
import { TagRequest, TagRequests } from '../model/tagrequest';
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

  getApplicationAllFindings(runid: number, appName: string): Observable<ApplicationFinding[]>  {
    const url = UriConstants.RUNS_URI + runid + '/apps/' + appName + '/findings/scorecard/low,medium,high,info?includeFF=false';
    return this.http.post<ApplicationFinding[]>(url, "");
  }

  getApplicationFindingsByTag(runid: number, appName: string, tag: Tag): Observable<ApplicationFinding[]>  {
    const url = UriConstants.RUNS_URI + runid + '/apps/' + appName + '/findings/scorecard/low,medium,high,info?includeFF=false';

    let tagReq: TagRequest = new TagRequest(tag.Value, true);

    let tagsReq: TagRequest [] = [];
    tagsReq.push(tagReq);

    let tagsRequest: TagRequests = new TagRequests(tagsReq);

    return this.http.post<ApplicationFinding[]>(url, tagsRequest);
  }

  getApplicationFindingsByTags(runid: number, appName: string, bin: Bin): Observable<ApplicationFinding[]>  {
    const url = UriConstants.RUNS_URI + runid + '/apps/' + appName + '/findings/scorecard/low,medium,high,info?includeFF=false';

    let tagsReq: TagRequest [] = [];
    bin.tags.forEach(tag => {
      let tagReq: TagRequest = new TagRequest(tag.name, true);
      tagsReq.push(tagReq);
    });

    let tagsRequest: TagRequests = new TagRequests(tagsReq);

    return this.http.post<ApplicationFinding[]>(url, tagsRequest);
  }

  getApplicationScorecard(runid: number, appName: string): Observable<ApplicationScoreCard>  {
    const url = UriConstants.RUNS_URI + runid + '/apps/' + appName + '/scorecard?includeFF=false';
    return this.http.post<ApplicationScoreCard>(url, "");
  }
}
