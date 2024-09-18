import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms'; // Import this if you're using ngModel
import { AppComponent } from './app.component';
import { BookingComponent } from './booking/booking.component';
import { BusSearchComponent } from './bus-search/bus-search.component';

@NgModule({
  declarations: [
    AppComponent,
    BookingComponent, // Add your components here
    BusSearchComponent
  ],
  imports: [
    BrowserModule,
    FormsModule, // Ensure FormsModule is imported for ngModel
    // other modules
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
