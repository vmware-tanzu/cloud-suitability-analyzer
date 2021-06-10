import { Injectable } from '@angular/core';
import {Observable} from "rxjs";
import {CsaRules} from "../model/rules";
import {UriConstants} from "../constants/uri-constants";
import {HttpClient} from "@angular/common/http";

@Injectable({
  providedIn: 'root'
})
export class RulesService {

  constructor(private http: HttpClient) { }

  getRules(): Observable<CsaRules>  {
    const url = UriConstants.RULES_URI;
    return this.http.get<CsaRules>(url);
  }
}
