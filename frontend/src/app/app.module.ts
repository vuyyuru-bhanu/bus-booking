import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LoginComponent } from './components/login/login.component';
import { BusSearchComponent } from './components/bus-search/bus-search.component';
import { BookingComponent } from './components/booking/booking.component';
import { AdminDashboardComponent } from './components/admin/admin-dashboard/admin-dashboard.component';
import { AdminBusesComponent } from './components/admin/admin-buses/admin-buses.component';
import { AdminRoutesComponent } from './components/admin/admin-routes/admin-routes.component';
import { AdminNotificationsComponent } from './components/admin/admin-notifications/admin-notifications.component';
import { NotificationComponent } from './components/notification/notification.component';
import { AuthInterceptor } from './services/auth.interceptor';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatSelectModule } from '@angular/material/select';
import { MatCardModule } from '@angular/material/card';
import { MatFormFieldModule } from '@angular/material/form-field';

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    BusSearchComponent,
    BookingComponent,
    AdminDashboardComponent,
    AdminBusesComponent,
    AdminRoutesComponent,
    AdminNotificationsComponent,
    NotificationComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule,
    BrowserAnimationsModule,
    MatInputModule,
    MatButtonModule,
    MatSelectModule,
    MatCardModule,
    MatFormFieldModule
  ],
  providers: [
    { provide: HTTP_INTERCEPTORS, useClass: AuthInterceptor, multi: true }
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }