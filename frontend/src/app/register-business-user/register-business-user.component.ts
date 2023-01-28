import { Component} from '@angular/core';
import { Router } from '@angular/router';
import { Profile } from '../model/profile';
import { ProfileService } from '../services/profile.service';
import { UserService } from '../services/user-service.service';

@Component({
  selector: 'app-register-business-user',
  templateUrl: './register-business-user.component.html',
  styleUrls: ['./register-business-user.component.css']
})
export class RegisterBusinessUserComponent {

  constructor(private userService: UserService, private router : Router, private profileService : ProfileService) { }

  user : any = new Object;
  profile : Profile = new Profile();

  register() {
    let usernameField = document.getElementById("username") as HTMLInputElement;
    let passwordField = document.getElementById("password") as HTMLInputElement;
    let emailField = document.getElementById("email") as HTMLInputElement;

    let companyNameField = document.getElementById("companyName") as HTMLInputElement;
    let webSiteField = document.getElementById("webSite") as HTMLInputElement;

    

    this.userService.regUserAuth(usernameField.value, passwordField.value, emailField.value, "business").subscribe(data => {
      console.log(data);
      this.router.navigateByUrl("/login")
    })
    
    this.profileService.regUserBusiness(usernameField.value, emailField.value, companyNameField.value, webSiteField.value).subscribe(data => {
      console.log(data);
    })
  }
  
}
