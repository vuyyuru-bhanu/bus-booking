import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { ApiService } from '../../services/api.service';

@Component({
  selector: 'app-booking',
  templateUrl: './booking.component.html',
  styleUrls: ['./booking.component.css']
})
export class BookingComponent implements OnInit {
  booking = { busId: 0, seatNumber: 0 };

  constructor(private apiService: ApiService, private route: ActivatedRoute, private router: Router) {}

  ngOnInit(): void {
    this.booking.busId = +this.route.snapshot.paramMap.get('id')!;
  }

  book(): void {
    this.apiService.createBooking(this.booking).subscribe(
      () => {
        this.router.navigate(['/search']);
      },
      (error) => {
        console.error('Booking failed', error);
      }
    );
  }
}