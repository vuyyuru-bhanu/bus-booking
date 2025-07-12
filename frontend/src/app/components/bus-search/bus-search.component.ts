import { Component, OnInit } from '@angular/core';
import { ApiService } from '../../services/api.service';

interface Route {
  id: number;
  origin: string;
  destination: string;
}

interface Bus {
  id: number;
  company: string;
  type: string;
  ac: boolean;
  route: Route;
  price: number;
  departure_time: string;
  arrival_time: string;
}

@Component({
  selector: 'app-bus-search',
  templateUrl: './bus-search.component.html',
  styleUrls: ['./bus-search.component.css']
})
export class BusSearchComponent implements OnInit {
  routes: Route[] = [];
  fromRouteId: number | null = null;
  toRouteId: number | null = null;
  ac: boolean | null = null;
  type: string = '';
  company: string = '';
  buses: Bus[] = [];
  errorMessage: string = '';

  constructor(private apiService: ApiService) {}

  ngOnInit(): void {
    this.loadRoutes();
  }

  loadRoutes(): void {
    this.apiService.getRoutes().subscribe({
      next: (data) => {
        this.routes = data.routes;
      },
      error: (err) => {
        this.errorMessage = 'Failed to load routes';
      }
    });
  }

  searchBuses(): void {
    if (!this.fromRouteId || !this.toRouteId) {
      this.errorMessage = 'Please select both origin and destination';
      return;
    }
    this.errorMessage = '';
    // Convert fromRouteId and toRouteId to numbers in case they are strings
    const fromId = typeof this.fromRouteId === 'string' ? parseInt(this.fromRouteId, 10) : this.fromRouteId;
    const toId = typeof this.toRouteId === 'string' ? parseInt(this.toRouteId, 10) : this.toRouteId;
    const origin = this.routes.find(r => r.id === fromId)?.origin || '';
    const destination = this.routes.find(r => r.id === toId)?.destination || '';
    this.apiService.getBuses(origin, destination, this.ac, this.type, this.company).subscribe({
      next: (data) => {
        this.buses = data.buses;
      },
      error: (err) => {
        this.errorMessage = 'Failed to load buses';
      }
    });
  }
}
