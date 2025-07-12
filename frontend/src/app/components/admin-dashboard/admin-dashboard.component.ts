import { Component, OnInit } from '@angular/core';
import { ApiService } from '../../services/api.service';

interface User {
  id: number;
  email: string;
  name: string;
  role: string;
}

@Component({
  selector: 'app-admin-dashboard',
  templateUrl: './admin-dashboard.component.html',
  styleUrls: ['./admin-dashboard.component.css']
})
export class AdminDashboardComponent implements OnInit {
  users: User[] = [];
  newUser = { email: '', name: '', role: 'user', password: '' };
  editUserId: number | null = null;
  editUserData: Partial<User> = {};
  errorMessage: string = '';
  successMessage: string = '';

  constructor(private apiService: ApiService) {}

  ngOnInit(): void {
    this.loadUsers();
  }

  loadUsers(): void {
    this.apiService.getUsers().subscribe({
      next: (data) => {
        this.users = data.users;
      },
      error: (err) => {
        this.errorMessage = 'Failed to load users';
      }
    });
  }

  addUser(): void {
    if (!this.newUser.email || !this.newUser.name || !this.newUser.password) {
      this.errorMessage = 'Please fill all required fields';
      return;
    }
    this.errorMessage = '';
    this.apiService.createUser(this.newUser).subscribe({
      next: () => {
        this.successMessage = 'User added successfully';
        this.newUser = { email: '', name: '', role: 'user', password: '' };
        this.loadUsers();
      },
      error: () => {
        this.errorMessage = 'Failed to add user';
      }
    });
  }

  editUser(user: User): void {
    this.editUserId = user.id;
    this.editUserData = { ...user };
  }

  updateUser(): void {
    if (!this.editUserId) return;
    this.apiService.updateUser(this.editUserId, this.editUserData).subscribe({
      next: () => {
        this.successMessage = 'User updated successfully';
        this.editUserId = null;
        this.loadUsers();
      },
      error: () => {
        this.errorMessage = 'Failed to update user';
      }
    });
  }

  cancelEdit(): void {
    this.editUserId = null;
    this.editUserData = {};
  }

  deleteUser(id: number): void {
    this.apiService.deleteUser(id).subscribe({
      next: () => {
        this.successMessage = 'User deleted successfully';
        this.loadUsers();
      },
      error: () => {
        this.errorMessage = 'Failed to delete user';
      }
    });
  }
}
