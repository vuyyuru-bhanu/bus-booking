import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class BusService {
  private apiUrl = 'http://localhost:8080/api';

  constructor(private http: HttpClient) { }

  searchBuses(query: any): Observable<any> {
    return this.http.get(`${this.apiUrl}/buses`, { params: query });
  }

  bookBus(bookingDetails: any): Observable<any> {
    return this.http.post(`${this.apiUrl}/book`, bookingDetails);
  }
}

