import { Component, OnInit } from '@angular/core';
import { ApiService } from '../../../services/api.service';

@Component({
  selector: 'app-admin-notifications',
  templateUrl: './admin-notifications.component.html',
  styleUrls: ['./admin-notifications.component.css']
})
export class AdminNotificationsComponent implements OnInit {
  notification = { message: '' };
  notifications: any[] = [];

  constructor(private apiService: ApiService) {}

  ngOnInit(): void {
    this.loadNotifications();
  }

  createNotification(): void {
    this.apiService.createNotification(this.notification).subscribe(
      () => {
        this.notification.message = '';
        this.loadNotifications();
      },
      (error) => {
        console.error('Error creating notification', error);
      }
    );
  }

  loadNotifications(): void {
    this.apiService.getNotifications().subscribe(
      (response: any) => {
        this.notifications = response.notifications;
      },
      (error) => {
        console.error('Error fetching notifications', error);
      }
    );
  }
}