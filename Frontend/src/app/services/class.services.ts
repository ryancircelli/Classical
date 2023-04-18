import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';

import { Class } from '../types'

@Injectable({
  providedIn: 'root'
})
export class ClassAPIService {
  public apiUrl = 'http://localhost:8000';

  constructor(public http: HttpClient) {}

  getTrendingClasses(): Observable<Class[]> {
    const httpOptions = {
      headers: new HttpHeaders({
        'Accept': 'application/json'
      })
    };
    return this.http.get<Class[]>(`${this.apiUrl}/getTrendingClasses`, httpOptions);
  }

  getSearchResults(className: String): Observable<Class[]> {
    const httpOptions = {
      headers: new HttpHeaders({
        'Accept': 'application/json'
      })
    };
    return this.http.get<Class[]>(`${this.apiUrl}/getClassesByName/${className}`, httpOptions);
  }

  addClass(className: String): Observable<any> {
    const httpOptions = {
      headers: new HttpHeaders({
        'Accept': 'application/json'
      })
    };
    const body = {
      "className": className
    };
    return this.http.post(`${this.apiUrl}/createClass`, body, httpOptions);
  }
}