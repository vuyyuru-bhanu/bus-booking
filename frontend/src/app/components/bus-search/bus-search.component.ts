import { Component, OnInit } from '@angular/core';
import { ApiService } from '../../services/api.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-bus-search',
  templateUrl: './bus-search.component.html',
  styleUrls: ['./bus-search.component.css']
})
export class BusSearchComponent implements OnInit {
  search = {
    origin: '',
    destination: '',
    ac: '',
    type: '',
    company: ''
  };
  buses: any[] = [];

  constructor(private apiService: ApiService, private router: Router) {}

  ngOnInit(): void {}

  searchBuses(): void {
    this.apiService.getBuses(this.search).subscribe(
      (response: any) => {
        this.buses = response.buses;
      },
      (error) => {
        console.error('Error fetching buses', error);
      }
    );
  }

  formatDuration(minutes: number): string {
    const hours = Math.floor(minutes / 60);
    const mins = minutes % 60;
    return `${hours}h ${mins}m`;
  }
}