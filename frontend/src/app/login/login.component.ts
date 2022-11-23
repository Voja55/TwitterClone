import { Component, OnInit } from '@angular/core';
import { UserService } from '../services/user-service.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  constructor(private userService: UserService) { }

  ngOnInit(): void {
  }

  user : any = new Object;
  submitted : boolean = false;

  login() {
    let usernameField = document.getElementById("username") as HTMLInputElement;
    let passwordField = document.getElementById("password") as HTMLInputElement;
    this.userService.loginAuth(usernameField.value, passwordField.value, "regular").subscribe(data => {
      console.log(data);
    })
  }

}
