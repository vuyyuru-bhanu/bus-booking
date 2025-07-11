import { Component } from '@angular/core';
import { AuthService } from '../../services/auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  user = { email: '', password: '' };

  constructor(private authService: AuthService, private router: Router) {}

  login(): void {
    this.authService.login(this.user.email, this.user.password).subscribe(
      () => {
        if (this.authService.isAdmin()) {
          this.router.navigate(['/admin']);
        } else {
          this.router.navigate(['/search']);
        }
      },
      (error) => {
        console.error('Login failed', error);
      }
    );
  }

  register(): void {
    this.authService.register(this.user).subscribe(
      () => {
        this.login();
      },
      (error) => {
        console.error('Registration failed', error);
      }
    );
  }
}