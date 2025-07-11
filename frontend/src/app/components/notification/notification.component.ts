import { Component, OnInit } from '@angular/core';
import { ApiService } from '../../services/api.service';

@Component({
  selector: 'app-notification',
  templateUrl: './notification.component.html',
  styleUrls: ['./notification.component.css']
})
export class NotificationComponent implements OnInit {
  notifications: any[] = [];

  constructor(private apiService: ApiService) {}

  ngOnInit(): void {
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