import { Component, OnInit } from '@angular/core';
import { UserService } from '../services/user-service.service';

@Component({
  selector: 'app-register-business-user',
  templateUrl: './register-business-user.component.html',
  styleUrls: ['./register-business-user.component.css']
})
export class RegisterBusinessUserComponent implements OnInit {

  constructor(private userService: UserService) { }

  ngOnInit(): void {
  }

  user : any = new Object;

  register() {
    let usernameField = document.getElementById("username") as HTMLInputElement;
    let passwordField = document.getElementById("password") as HTMLInputElement;
    this.userService.regUserAuth(usernameField.value, passwordField.value, "business").subscribe(data => {
      console.log(data);
    })
  }
  
}
