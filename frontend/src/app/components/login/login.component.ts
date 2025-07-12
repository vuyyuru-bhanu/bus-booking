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
    const loginId = this.user.name ? this.user.name : this.user.email;
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