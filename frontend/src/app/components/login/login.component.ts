import { Component } from '@angular/core';
import { AuthService } from '../../services/auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  user = { name: '', email: '', password: '' };

  constructor(private authService: AuthService, private router: Router) {}

  login(): void {
    let loginId = '';
    if (this.user.name === 'admin') {
      // Admin login with username
      loginId = this.user.name;
    } else {
      // Regular user login with email
      loginId = this.user.email;
    }
    this.authService.login(loginId, this.user.password).subscribe(
      () => {
        if (this.authService.isAdmin()) {
          this.router.navigate(['/admin']);
        } else {
          this.router.navigate(['/search']);
        }
      },
      (error) => {
        console.error('Login failed', error);
        alert('Login failed: ' + (error.error?.error || 'Please check your credentials.'));
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
        if (error && error.error && error.error.error) {
          alert('Registration failed: ' + error.error.error);
        } else {
          alert('Registration failed. Please check your input.');
        }
      }
    );
  }
}