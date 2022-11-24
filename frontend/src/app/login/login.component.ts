import { Component, OnInit } from '@angular/core';
import { StoreService } from '../services/store-service.service';
import { UserService } from '../services/user-service.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  constructor(private userService: UserService, private storeService:StoreService) { }

  ngOnInit(): void {
  }

  user : any = new Object;
  submitted : boolean = false;

  login() {
    let usernameField = document.getElementById("username") as HTMLInputElement;
    let passwordField = document.getElementById("password") as HTMLInputElement;
    this.userService.loginAuth(usernameField.value, passwordField.value, "regular").subscribe(data => {
      console.log(data);
      this.storeService.login(data.jwt)
    })
  }

}
