import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UserService } from '../services/user-service.service';

@Component({
  selector: 'app-register-business-user',
  templateUrl: './register-business-user.component.html',
  styleUrls: ['./register-business-user.component.css']
})
export class RegisterBusinessUserComponent implements OnInit {

  constructor(private userService: UserService, private router : Router) { }

  ngOnInit(): void {
  }

  user : any = new Object;

  register() {
    let usernameField = document.getElementById("username") as HTMLInputElement;
    let passwordField = document.getElementById("password") as HTMLInputElement;
    let emailField = document.getElementById("email") as HTMLInputElement;
    this.userService.regUserAuth(usernameField.value, passwordField.value, emailField.value, "business").subscribe(data => {
      console.log(data);
      this.router.navigateByUrl("/login")
    })
  }
  
}
