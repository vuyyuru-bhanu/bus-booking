import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ApiService {
  private apiUrl = '/api';

  constructor(private http: HttpClient) {}

  getBuses(origin: string, destination: string, ac: boolean, type: string, company: string): Observable<any> {
    const params: any = {
      origin,
      destination,
      ac: ac.toString(),
      type,
      company
    };
    return this.http.get(this.apiUrl + '/buses', { params });
  }

  createBooking(booking: any): Observable<any> {
    return this.http.post(this.apiUrl + '/bookings', booking);
  }

  createRoute(route: any): Observable<any> {
    return this.http.post(this.apiUrl + '/admin/routes', route);
  }

  createBus(bus: any): Observable<any> {
    return this.http.post(this.apiUrl + '/admin/buses', bus);
  }

  uploadBusImage(busId: number, image: File): Observable<any> {
    const formData = new FormData();
    formData.append('image', image);
    return this.http.post(this.apiUrl + '/admin/buses/' + busId + '/image', formData);
  }

  createNotification(notification: any): Observable<any> {
    return this.http.post(this.apiUrl + '/admin/notifications', notification);
  }

  getNotifications(): Observable<any> {
    return this.http.get(this.apiUrl + '/notifications');
  }

  getRoutes(): Observable<any> {
    return this.http.get(this.apiUrl + '/routes');
  }

  getUsers(): Observable<any> {
    return this.http.get(this.apiUrl + '/admin/users');
  }

  createUser(user: any): Observable<any> {
    return this.http.post(this.apiUrl + '/admin/users', user);
  }

  updateUser(id: number, user: any): Observable<any> {
    return this.http.put(this.apiUrl + '/admin/users/' + id, user);
  }

  deleteUser(id: number): Observable<any> {
    return this.http.delete(this.apiUrl + '/admin/users/' + id);
  }
}
