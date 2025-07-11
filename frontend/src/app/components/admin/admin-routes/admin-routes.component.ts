import { Component } from '@angular/core';
import { ApiService } from '../../../services/api.service';

@Component({
  selector: 'app-admin-routes',
  templateUrl: './admin-routes.component.html',
  styleUrls: ['./admin-routes.component.css']
})
export class AdminRoutesComponent {
  route = {
    origin: '',
    destination: '',
    departureTime: '',
    duration: 0
  };

  constructor(private apiService: ApiService) {}

  createRoute(): void {
    this.apiService.createRoute(this.route).subscribe(
      () => {
        this.route = { origin: '', destination: '', departureTime: '', duration: 0 };
      },
      (error) => {
        console.error('Error creating route', error);
      }
    );
  }
}