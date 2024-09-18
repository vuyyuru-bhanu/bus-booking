import { Component } from '@angular/core';
import { BusService } from '../services/bus.service';

@Component({
  selector: 'app-bus-search',
  templateUrl: './bus-search.component.html',
  styleUrls: ['./bus-search.component.css']
})
export class BusSearchComponent {
  buses: any[] = [];
  searchQuery: string = '';

  constructor(private busService: BusService) {}

  search() {
    this.busService.searchBuses({ query: this.searchQuery }).subscribe(data => {
      this.buses = data;
    });
  }
}

