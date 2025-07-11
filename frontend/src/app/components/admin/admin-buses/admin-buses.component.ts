import { Component } from '@angular/core';
import { ApiService } from '../../../services/api.service';

@Component({
  selector: 'app-admin-buses',
  templateUrl: './admin-buses.component.html',
  styleUrls: ['./admin-buses.component.css']
})
export class AdminBusesComponent {
  bus = {
    routeId: 0,
    company: '',
    ac: false,
    type: '',
    capacity: 0,
    availableSeats: 0,
    price: 0,
    amenities: ''
  };

  constructor(private apiService: ApiService) {}

  createBus(): void {
    const busData = { ...this.bus, amenities: this.bus.amenities.split(',').map(item => item.trim()) };
    this.apiService.createBus(busData).subscribe(
      () => {
        this.bus = { routeId: 0, company: '', ac: false, type: '', capacity: 0, availableSeats: 0, price: 0, amenities: '' };
      },
      (error) => {
        console.error('Error creating bus', error);
      }
    );
  }
}