import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';

import { Class, Post } from '../types'

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

  getClassPosts(className: String): Observable<Post[]> {
    const httpOptions = {
      headers: new HttpHeaders({
        'Accept': 'application/json'
      })
    };
    return this.http.get<Post[]>(`${this.apiUrl}/getPostsByClassName/${className}`, httpOptions);
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

  createClassPost(className:string, url: String): Observable<any> {
    const httpOptions = {
      headers: new HttpHeaders({
        'Accept': 'application/json'
      })
    };
    const body = {
      "postClassName" : className,
      "postContent": url,
      "postName" : className
    };
    return this.http.post(`${this.apiUrl}/createClassPost`, body, httpOptions);
  }

  increasePostVotes(className: String, postID: number): Observable<any> {
    const httpOptions = {
      headers: new HttpHeaders({
        'Accept': 'application/json'
      })
    };
    console.log(postID)
    const body = {
      "postID" : postID,
      "postClassName": className,
    };
    return this.http.post(`${this.apiUrl}/increasePostVotes`, body, httpOptions);
  }

  decreasePostVotes(className: String, postID: number): Observable<any> {
    const httpOptions = {
      headers: new HttpHeaders({
        'Accept': 'application/json'
      })
    };
    const body = {
      "postID" : postID,
      "postClassName": className,
    };
    return this.http.post(`${this.apiUrl}/decreasePostVotes`, body, httpOptions);
  }
}