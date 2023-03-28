import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';

import { Class } from '../types'

@Injectable({
  providedIn: 'root'
})
export class ClassAPIService {
  private apiUrl = 'http://localhost:8000';

  constructor(private http: HttpClient) {}

  getClasses(): Observable<Class[]> {
    const httpOptions = {
      headers: new HttpHeaders({
        'Accept': 'application/json'
      })
    };
    return this.http.get<Class[]>(`${this.apiUrl}/getClasses`, httpOptions);
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