import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class LoaningService {

  private baseUrl = '/api/loanings';

  constructor(private http: HttpClient) { }

  getLoaning(ID: number): Observable<any> {
    return this.http.get(`${this.baseUrl}/${ID}`);
  }

  createLoaning(loaning: Object): Observable<Object> {
    return this.http.post(`${this.baseUrl}`, loaning);
  }

  updateLoaning(ID: number, value: any): Observable<Object> {
    return this.http.put(`${this.baseUrl}/${ID}`, value);
  }

  deleteLoaning(ID: number): Observable<any> {
    return this.http.delete(`${this.baseUrl}/${ID}`, { responseType: 'text' });
  }

  getLoaningsList(): Observable<any> {
    return this.http.get(`${this.baseUrl}`);
  }
}
