import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Profile } from '../model/profile';
import { ProfileService } from '../services/profile.service';
import { UserService } from '../services/user-service.service';

@Component({
  selector: 'app-register-user',
  templateUrl: './register-user.component.html',
  styleUrls: ['./register-user.component.css']
})
export class RegisterUserComponent implements OnInit {

  constructor(private userService: UserService, private router : Router, private profileService : ProfileService) { }

  ngOnInit(): void {
  }

  user : any = new Object;
  profile : Profile = new Profile();

  register() {
    let usernameField = document.getElementById("username") as HTMLInputElement;
    let passwordField = document.getElementById("password") as HTMLInputElement;
    let emailField = document.getElementById("email") as HTMLInputElement;

    let firstNameField = document.getElementById("firstName") as HTMLInputElement;
    let lastNameField = document.getElementById("lastName") as HTMLInputElement;
    let addressField = document.getElementById("address") as HTMLInputElement;
    let ageField = document.getElementById("age") as HTMLInputElement;
    //TODO: izboriti se sa ovime
    //let gender = document.getElementById("gender") as HTMLInputElement;

    this.userService.regUserAuth(usernameField.value, passwordField.value, emailField.value, "regular").subscribe(data => {
      console.log(data);
      this.router.navigateByUrl("/login")
    })

    this.profileService.regUserRegular(usernameField.value, firstNameField.value, lastNameField.value,
      //Dodati za gender
      addressField.value,  Number.parseInt(ageField.value), true).subscribe(data => {
      console.log(data);
    })
  }

}
