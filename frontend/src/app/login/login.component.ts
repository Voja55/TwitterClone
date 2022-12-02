import { Component, OnInit } from '@angular/core';
import { throws } from 'assert';
import { StoreService } from '../services/store-service.service';
import { UserService } from '../services/user-service.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  captcha: string;
  email: string;


  constructor(private userService: UserService, private storeService:StoreService) {
    this.captcha = '';
    this.email = 'Test';
   }

  ngOnInit(): void {
  }

  resolved(captchaResponse: string) {
    this.captcha = captchaResponse;
    console.log('resolved captcha with response: ' + this.captcha);
  }


  user : any = new Object;
  submitted : boolean = false;

  login() {
    let usernameField = document.getElementById("username") as HTMLInputElement;
    let passwordField = document.getElementById("password") as HTMLInputElement;
    if (this.captcha != '') {
      this.userService.loginAuth(usernameField.value, passwordField.value, "regular").subscribe(data => {
        console.log(data);
        this.storeService.login(data.jwt)
      })
    } else {
      console.log("captcha not passed")
    }

  }

}
