import { Component} from '@angular/core';
import { Router } from '@angular/router';
import { StoreService } from '../services/store-service.service';
import { UserService } from '../services/user-service.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {

  captcha: string;
  email: string;


  constructor(private userService: UserService, public storeService:StoreService, private router : Router) {
    this.captcha = '';
    this.email = 'Test';
   }

  resolved(captchaResponse: string) {
    this.captcha = captchaResponse;
    console.log('resolved captcha with response: ' + this.captcha);
    if (captchaResponse != '') {
      this.captchaValid = true
    }
  }


  user : any = new Object;
  submitted : boolean = false;
  captchaValid : boolean = false;

  login() {
    let usernameField = document.getElementById("username") as HTMLInputElement;
    let passwordField = document.getElementById("password") as HTMLInputElement;
    if (this.captcha != '') {
      this.userService.loginAuth(usernameField.value, passwordField.value, "regular").subscribe(data => {
        console.log(data);
        this.storeService.login(data.jwt)
        this.router.navigateByUrl("/");
      })
    } else {
      console.log("captcha not passed")
    }

  }

}
