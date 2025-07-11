import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './components/login/login.component';
import { BusSearchComponent } from './components/bus-search/bus-search.component';
import { BookingComponent } from './components/booking/booking.component';
import { AdminDashboardComponent } from './components/admin/admin-dashboard/admin-dashboard.component';
import { AdminBusesComponent } from './components/admin/admin-buses/admin-buses.component';
import { AdminRoutesComponent } from './components/admin/admin-routes/admin-routes.component';
import { AdminNotificationsComponent } from './components/admin/admin-notifications/admin-notifications.component';
import { AdminGuard } from './guards/admin.guard';

const routes: Routes = [
  { path: 'login', component: LoginComponent },
  { path: 'search', component: BusSearchComponent },
  { path: 'booking/:id', component: BookingComponent },
  { 
    path: 'admin', 
    component: AdminDashboardComponent,
    canActivate: [AdminGuard],
    children: [
      { path: 'buses', component: AdminBusesComponent },
      { path: 'routes', component: AdminRoutesComponent },
      { path: 'notifications', component: AdminNotificationsComponent }
    ]
  },
  { path: '', redirectTo: '/login', pathMatch: 'full' }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }