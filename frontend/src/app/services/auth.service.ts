import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { tap } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private apiUrl = '/api';

  constructor(private http: HttpClient) {}

  login(loginId: string, password: string): Observable<any> {
    return this.http.post(`${this.apiUrl}/login`, { loginId, password }).pipe(
      tap((response: any) => {
        if (response.token) {
          localStorage.setItem('token', response.token);
          localStorage.setItem('role', response.role);
        }
      })
    );
  }

  register(user: any): Observable<any> {
    return this.http.post(`${this.apiUrl}/register`, user);
  }

  isAdmin(): boolean {
    return localStorage.getItem('role') === 'admin';
  }

  logout(): void {
    localStorage.removeItem('token');
    localStorage.removeItem('role');
  }
}