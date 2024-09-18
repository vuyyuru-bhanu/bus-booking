import { Component } from '@angular/core';
import { BusService } from '../services/bus.service';

@Component({
  selector: 'app-booking',
  templateUrl: './booking.component.html',
  styleUrls: ['./booking.component.css']
})
export class BookingComponent {
  bookingDetails: any = {};

  constructor(private busService: BusService) {}

  book() {
    this.busService.bookBus(this.bookingDetails).subscribe(response => {
      alert('Booking successful!');
    });
  }
}

